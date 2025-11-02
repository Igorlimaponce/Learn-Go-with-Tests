package maps

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		// Aqui o .Error() pega a string e passa para assertStrings
		assertStrings(t, err, ErrNotFound)
	})

}

func assertString(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertStrings(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

// By creating a new helper we were able to simplify our test, and start using our ErrNotFound variable so our test doesn't fail if we change the error text in the future.
