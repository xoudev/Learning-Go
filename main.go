package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

type Dictionnaire struct {
	Mots []string `json:"mots"`
}

type Partie struct {
	MotSecret        string
	MotVisible       []string
	LettresUtilisees []string
	Vies             int
}

// Efface l'écran du terminal proprement
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func extractData() []string {
	donnees, err := os.ReadFile("data/mots_pendu.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}

	var dict Dictionnaire
	err = json.Unmarshal(donnees, &dict)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return nil
	}

	return dict.Mots
}

func initialiserPartie(motChoisi string) Partie {
	motCache := make([]string, len(motChoisi))
	for i := range motCache {
		motCache[i] = "_"
	}

	return Partie{
		MotSecret:        motChoisi,
		MotVisible:       motCache,
		LettresUtilisees: []string{},
		Vies:             10,
	}
}

func dejaUtilisee(lettre string, lettres []string) bool {
	for _, l := range lettres {
		if l == lettre {
			return true
		}
	}
	return false
}

func estGagne(motVisible []string) bool {
	for _, char := range motVisible {
		if char == "_" {
			return false
		}
	}
	return true
}

// Affiche l'interface de jeu propre à chaque tour
func afficherInterface(maPartie Partie, message string) {
	clearScreen()
	fmt.Println("==================================================")
	fmt.Println("          💀 LE JEU DU PENDU DE LA MORT 💀        ")
	fmt.Println("==================================================")
	fmt.Printf("❤️  Vies restantes : %d / 10\n", maPartie.Vies)
	fmt.Printf("🔤 Lettres essayées : %s\n", strings.Join(maPartie.LettresUtilisees, ", "))
	fmt.Println("--------------------------------------------------")
	fmt.Printf("\n   Mot à deviner :  %s\n\n", strings.Join(maPartie.MotVisible, "   "))
	fmt.Println("--------------------------------------------------")

	if message != "" {
		fmt.Println("👉 " + message)
		fmt.Println("--------------------------------------------------")
	}
}

func main() {
	mots := extractData()
	if len(mots) == 0 {
		return
	}

	motAleatoire := mots[rand.Intn(len(mots))]
	maPartie := initialiserPartie(motAleatoire)

	message := "Bienvenue ! Essayez de deviner le mot."

	// Boucle principale du jeu
	for maPartie.Vies > 0 {
		// 1. On affiche l'interface propre
		afficherInterface(maPartie, message)

		// 2. On demande la saisie
		var lettreUser string
		fmt.Print("Entrez une lettre : ")
		fmt.Scan(&lettreUser)

		lettreUser = strings.TrimSpace(strings.ToLower(lettreUser))
		caracteres := []rune(lettreUser)

		// Validations
		if len(caracteres) != 1 {
			message = "❌ Erreur : Vous devez entrer un et un seul caractère."
			continue
		}

		if !unicode.IsLetter(caracteres[0]) {
			message = "❌ Erreur : Ce n'est pas une lettre valide (chiffres/symboles interdits)."
			continue
		}

		if dejaUtilisee(lettreUser, maPartie.LettresUtilisees) {
			message = fmt.Sprintf("⚠️ Vous avez déjà proposé la lettre '%s' !", lettreUser)
			continue
		}

		// Ajout aux lettres utilisées
		maPartie.LettresUtilisees = append(maPartie.LettresUtilisees, lettreUser)

		// Vérification dans le mot secret
		lettreTrouvee := false
		for i, char := range []rune(maPartie.MotSecret) {
			if string(char) == lettreUser {
				maPartie.MotVisible[i] = lettreUser
				lettreTrouvee = true
			}
		}

		// Feedback et mise à jour des vies
		if !lettreTrouvee {
			maPartie.Vies--
			message = fmt.Sprintf("❌ Raté ! La lettre '%s' n'est pas dans le mot.", lettreUser)
		} else {
			message = fmt.Sprintf("✨ Bien joué ! La lettre '%s' est présente.", lettreUser)
		}

		// Vérification de la victoire
		if estGagne(maPartie.MotVisible) {
			clearScreen()
			fmt.Println("==================================================")
			fmt.Println("🎉 INCROYABLE ! VOUS AVEZ GAGNÉ ! 🎉")
			fmt.Printf("Le mot secret était bien : %s\n", maPartie.MotSecret)
			fmt.Println("==================================================")
			return
		}
	}

	// Défaite
	clearScreen()
	fmt.Println("==================================================")
	fmt.Println("💀 PERDU ! VOUS ÊTES PENDU ! 💀")
	fmt.Printf("Le mot secret était : %s\n", maPartie.MotSecret)
	fmt.Println("==================================================")
}
