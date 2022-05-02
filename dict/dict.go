package dict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("not found")
	errWordExists = errors.New("word already exists")
	errCantUpdate = errors.New("can't update non-existing word")
)

// Search word
func (d Dictionary) Search(word string) (string, error) {
	value, isExists := d[word]

	if isExists {
		return value, nil
	}

	return "", errNotFound
}

// Add word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

// Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}

	return nil
}

// Delete word
func (d Dictionary) Delete(word string) {
		delete(d, word)
}
