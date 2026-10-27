// Ce fichier présente les bases de Go.
// Exécute-le depuis la racine du projet avec :
//
//	go run ./basics
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// 1. Variables et constantes
	var age int = 25
	nom := "Alice" // Déclaration courte, uniquement dans une fonction.
	const pays = "France"
	fmt.Println("Variables :", nom, age, pays)

	// 2. Types de base
	var (
		entier    int     = 42
		decimal   float64 = 3.14
		vrai      bool    = true
		caractere rune    = 'A'
		texte     string  = "Bonjour"
	)
	fmt.Printf("Types : %d, %.2f, %t, %c, %s\n", entier, decimal, vrai, caractere, texte)

	// 3. Conditions
	if age >= 18 {
		fmt.Println("Condition : majeur")
	} else {
		fmt.Println("Condition : mineur")
	}

	// 4. Switch
	jour := 2
	switch jour {
	case 1:
		fmt.Println("Switch : lundi")
	case 2:
		fmt.Println("Switch : mardi")
	default:
		fmt.Println("Switch : autre jour")
	}

	// 5. Boucles
	for i := 1; i <= 3; i++ {
		fmt.Println("Boucle classique :", i)
	}

	compteur := 3
	for compteur > 0 {
		fmt.Println("Boucle conditionnelle :", compteur)
		compteur--
	}

	// 6. Array : taille fixe
	jours := [3]string{"lundi", "mardi", "mercredi"}
	for index, jour := range jours {
		fmt.Println("Array :", index, jour)
	}

	// 7. Slice : taille flexible
	nombres := []int{1, 2, 3}
	nombres = append(nombres, 4, 5)
	fmt.Println("Slice :", nombres, "longueur :", len(nombres), "capacité :", cap(nombres))
	fmt.Println("Sous-slice :", nombres[1:4])

	// 8. Map : clé => valeur
	scores := map[string]int{"Alice": 15, "Bob": 8}
	scores["Charlie"] = 18
	delete(scores, "Bob")
	score, existe := scores["Alice"]
	fmt.Println("Map :", scores, "Alice :", score, "existe :", existe)

	// 9. Fonctions et valeurs de retour multiples
	somme, produit := calculer(4, 5)
	fmt.Println("Fonction :", somme, produit)

	// 10. Pointeurs
	valeur := 10
	augmenter(&valeur)
	fmt.Println("Pointeur :", valeur)

	// 11. Struct et méthode
	etudiant := Etudiant{Nom: "Alice", Note: 15}
	fmt.Println("Struct :", etudiant.Nom, etudiant.Statut())

	// 12. Interface
	var annonceur Annonceur = etudiant
	fmt.Println("Interface :", annonceur.Annoncer())

	// 13. Gestion d'erreurs
	if err := verifierNote(21); err != nil {
		fmt.Println("Erreur :", err)
	}

	// 14. Chaînes de caractères
	mot := "go"
	fmt.Println("Chaîne :", strings.ToUpper(mot), strings.Contains(mot, "o"))
}

func calculer(a, b int) (int, int) {
	return a + b, a * b
}

func augmenter(nombre *int) {
	*nombre++
}

type Etudiant struct {
	Nom  string
	Note int
}

func (e Etudiant) Statut() string {
	if e.Note >= 10 {
		return "admis"
	}
	return "ajourné"
}

type Annonceur interface {
	Annoncer() string
}

func (e Etudiant) Annoncer() string {
	return e.Nom + " a obtenu " + fmt.Sprint(e.Note) + "/20"
}

func verifierNote(note int) error {
	if note < 0 || note > 20 {
		return errors.New("la note doit être comprise entre 0 et 20")
	}
	return nil
}
