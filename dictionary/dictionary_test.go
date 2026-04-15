package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"awake": "become aware of"}
	dictionary := Dictionary{"awake": "become aware of"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("awake")
		want := "become aware of"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
