package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
