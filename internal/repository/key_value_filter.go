package repository

// KeyValueFilter matches when all key-value pairs are present and equal.
type KeyValueFilter map[string]string

// Match returns true if metadata contains all filter key-value pairs.
func (f KeyValueFilter) Match(metadata map[string]string) bool {
	if metadata == nil {
		return len(f) == 0
	}
	for k, v := range f {
		if metadata[k] != v {
			return false
		}
	}
	return true
}
