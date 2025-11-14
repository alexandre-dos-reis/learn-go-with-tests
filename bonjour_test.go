package main

import "testing"

// Il doit être dans un fichier avec un nom comme xxx_test.go
// La fonction de test doit commencer par le mot Test La fonction de test
// ne prend qu'un seul argument t *testing.T
// Pour utiliser le type *testing.T, vous devez import "testing",
// comme nous l'avons fait avec fmt dans l'autre fichier
func TestBonjour(t *testing.T) {
	resultat := Bonjour("Chris")
	attendu := "Bonjour, Chris"

	if resultat != attendu {
		t.Errorf("reçu %q attendu %q", resultat, attendu)
	}
}
