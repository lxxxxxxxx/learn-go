package maps

import (
	"testing"
)

const (
	ErrorNotFound     = DictionaryErr("can not found word")
	ErrorWordExist    = DictionaryErr("word exist")
	ErrorWordNotExist = DictionaryErr("word not exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (res string, err error) {
	str, ok := d[word]
	if !ok {
		res = ""
		err = ErrorNotFound
	} else {
		res = str
		err = nil
	}
	return
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		return ErrorWordNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test", "aa": "bb"}

	t.Run("search konw word", func(t *testing.T) {
		want := "this is just a test"
		got, _ := dictionary.Search("test")

		AssertStrings(t, want, got)
	})

	t.Run("search unkonw word", func(t *testing.T) {
		_, err := dictionary.Search("unkonw")
		AssertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("test new", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "aa"
		value := "bb"
		dictionary.Add(key, value)

		AssertDefinition(t, dictionary, key, value)
	})

	t.Run("test existing", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "aa"
		value := "bb"

		err := dictionary.Add(key, value)
		AssertError(t, nil, err)

		err = dictionary.Add(key, value)
		AssertError(t, err, ErrorWordExist)
		AssertDefinition(t, dictionary, key, value)
	})

}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{}

	key := "aa"
	value := "bb"
	err := dictionary.Add(key, value)

	AssertError(t, err, nil)

	newValue := "cc"
	err = dictionary.Update(key, newValue)
	AssertError(t, err, nil)
	AssertDefinition(t, dictionary, key, newValue)
}

func TestDelete(t *testing.T) {
	key := "aa"
	value := "bb"

	dictionary := Dictionary{key: value}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	AssertError(t, err, ErrorNotFound)

}

func AssertStrings(t *testing.T, want string, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()
	if want != got {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

func AssertDefinition(t *testing.T, dictionary Dictionary, key, word string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should found word but not :", err)
	}
	if word != got {
		t.Errorf("want '%s' but got '%s'", word, got)
	}
}
