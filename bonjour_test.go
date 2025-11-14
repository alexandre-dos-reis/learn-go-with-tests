package main

import "testing"

// Il doit être dans un fichier avec un nom comme xxx_test.go
// La fonction de test doit commencer par le mot Test La fonction de test
// ne prend qu'un seul argument t *testing.T
// Pour utiliser le type *testing.T, vous devez import "testing",
// comme nous l'avons fait avec fmt dans l'autre fichier
func TestBonjour(t *testing.T) {
	verifierMessageCorrect := func(t testing.TB, resultat, attendu string) {
		t.Helper()
		// t.Helper() est nécessaire pour dire à la suite de tests que cette méthode
		// est une fonction helper.
		// En faisant cela, quand elle échoue, le numéro de ligne rapporté sera dans
		// notre appel de fonction plutôt que dans notre fonction helper de test
		if resultat != attendu {
			t.Errorf("reçu %q attendu %q", resultat, attendu)
		}
	}

	t.Run("dire bonjour aux personnes", func(t *testing.T) {
		resultat := Bonjour("Chris")
		attendu := "Bonjour, Chris"
		verifierMessageCorrect(t, resultat, attendu)
	})

	t.Run("dire 'Bonjour, Monde' quand une chaîne vide est fournie", func(t *testing.T) {
		resultat := Bonjour("")
		attendu := "Bonjour, Monde"
		verifierMessageCorrect(t, resultat, attendu)
	})
}
