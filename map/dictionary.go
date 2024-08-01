package dictionary

/*
An interesting property of maps is that you
can modify them without passing as an address to it (e.g &myMap).

A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it,
but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value.
A nil map behaves like an empty map when reading, but attempts to write
to a nil map will cause a runtime panic.

Therefore, you should never initialize a nil map variable: `var m map[string]string`

Instead, you can initialize an empty map or use the make keyword to create a map for you:

var dictionary = map[string]string{}
var dictionary = make(map[string]string)

Both approaches create an empty hash map and point dictionary at it.
Which ensures that you will never get a runtime panic.
*/
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update/delete word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	// prevent new words from being added
	_, err := d.Search(word)

	if err == nil {
		d[word] = newDefinition
		return nil
	}

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	return err
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	if err == nil {
		delete(d, word)
	}

	return err
}
