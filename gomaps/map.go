package gomaps

import (
	"errors"
)

var ErrNotFound = errors.New("NOT_PRESENT")
var ErrKeyAlreadyExist = errors.New("KEY_ALREADY_EXIST")

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dict map[string]string, test string) string {
	val, ok := dict[test]
	if !ok {
		return ""
	}
	return val
}

func (d Dictionary) SearchInMap(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.SearchInMap(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.SearchInMap(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
