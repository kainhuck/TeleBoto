package telegram

import "os"

func Exists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return !s.IsDir()
		}
		return false
	}

	return !s.IsDir()
}
