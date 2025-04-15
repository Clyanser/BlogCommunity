package utils

func IsInList(key string, list []string) bool {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}
