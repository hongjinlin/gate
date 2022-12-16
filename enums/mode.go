package enums

import "strings"

type Mode int

var (
	modeMap = map[string]Mode{
		"main":  Main,
		"check": Check,
	}
)

const (
	Main  Mode = iota
	Check Mode = iota
)

func ParseMode(str string) (Mode, bool) {
	c, ok := modeMap[strings.ToLower(str)]
	return c, ok
}
