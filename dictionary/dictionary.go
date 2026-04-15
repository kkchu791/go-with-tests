package dictionary

type Dictionary map[string]string // {"key": "value"}

// Use a pointer when you need to mutate the value or the struct is large.
// Use a value when you just need to read it.
func (d Dictionary) Search(word string) string, error {
	return d[word], nil
}
