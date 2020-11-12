package discord

import "strings"

// Color Enum
const (
	C_DEFAULT     = 0
	C_AQUA        = 1752220
	C_GREEN       = 3066993
	C_BLUE        = 3447003
	C_PURPLE      = 10181046
	C_GOLD        = 15844367
	C_ORANGE      = 15105570
	C_RED         = 15158332
	C_GREY        = 9807270
	C_DARKER_GREY = 8359053
	C_NAVY        = 3426654
	C_DARK_AQUA   = 1146986
	C_DARK_GREEN  = 2067276
	C_DARK_BLUE   = 2123412
	C_DARK_PURPLE = 7419530
	C_DARK_GOLD   = 12745742
	C_DARK_ORANGE = 11027200
	C_DARK_RED    = 10038562
	C_DARK_GREY           = 9936031
	C_LIGHT_GREY          = 12370112
	C_DARK_NAVY           = 2899536
	C_LUMINOUS_VIVID_PINK = 16580705
	C_DARK_VIVID_PINK     = 12320855
)

// Generic message format for errors
func ErrorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}

// Generic message format for successful operations
func SuccessMessage(title string, message string) string {
	return "✅  **" + title + "**\n" + message
}

// stringHasPrefix checks whether or not the string contains one of the given prefixes and returns the string without the prefix
func stringHasPrefix(str string, prefixes []string, ignoreCase bool) (bool, string) {
	for _, prefix := range prefixes {
		stringToCheck := str
		if ignoreCase {
			stringToCheck = strings.ToLower(stringToCheck)
			prefix = strings.ToLower(prefix)
		}
		if strings.HasPrefix(stringToCheck, prefix) {
			return true, string(str[len(prefix):])
		}
	}
	return false, str
}

// stringArrayContains checks whether or not the given string array contains the given string
func stringArrayContains(array []string, str string, ignoreCase bool) bool {
	if ignoreCase {
		str = strings.ToLower(str)
	}
	for _, value := range array {
		if ignoreCase {
			value = strings.ToLower(value)
		}
		if value == str {
			return true
		}
	}
	return false
}
