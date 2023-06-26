package gomaps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "hello world"}

	got := Search(dict, "test")
	want := "hello world"
	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	// fmt.Println("errors: ", got, want)
	if got != want {
		t.Errorf("got %q error but want %q error", got, want)
	}
}

//A map value is a pointer to a runtime.hmap structure.
// We did not need to use the pointers explicitly in map.
//Map types are reference types, like pointers or slices, and so the value of m(var m map[string]string) is nil;
//it doesn’t point to an initialized map.
//A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that.
//To initialize a map, use the built in make function or initialize it like m := map[string]string{}

func TestSearchInMap(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}
	got, _ := dict.SearchInMap("test")
	want := "this is a test"
	assertString(t, got, want)

	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{}
		key := "my"
		value := "num"
		dict.Add(key, value)
		got, _ := dict.SearchInMap("my")
		want := "num"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"one": "oneval"}
		_, err := dict.SearchInMap("two")
		want := "NOT_PRESENT"

		if err == nil {
			t.Errorf("expected error but didn't got")
		}
		assertString(t, err.Error(), want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dict.SearchInMap("unknown")
		assertError(t, err, ErrNotFound)
	})

	t.Run("add word and test", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is just a test")
		got, err := dict.SearchInMap("test")
		want := "this is just a test"
		if err != nil {
			t.Errorf("word should be added to dict")
		}
		assertString(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "hey"
		value := "added to the dictionary"
		err := dict.Add(word, value)
		if err != nil {
			t.Errorf("should added the new word")
		}
		assertError(t, err, nil)
		assertDefinition(t, dict, word, value)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "hey"
		value := "added to dict"
		dict := Dictionary{word: value}
		err := dict.Add(word, value)
		assertError(t, err, ErrKeyAlreadyExist)
		assertDefinition(t, dict, word, value)
	})
}

func assertDefinition(t *testing.T, dict Dictionary, word, value string) {
	t.Helper()
	got, err := dict.SearchInMap(word)
	if err != nil {
		t.Errorf("word should be added to the dictionary")
	}
	assertString(t, got, value)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "hey"
		definition := "updated definition"
		dict := Dictionary{word: definition}
		newDefinition := "new value for hey word"
		err := dict.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word update", func(t *testing.T) {
		word := "test"
		definition := "new definition for test"
		dict := Dictionary{}
		err := dict.Update(word, definition)
		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "new word to delete"}

	_, err := dict.SearchInMap(word)
	assertError(t, err, nil)
	dict.Delete(word)
	t.Log("Deleted from dictionary")
}
