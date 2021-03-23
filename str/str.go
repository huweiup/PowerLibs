package str

import (
	. "github.com/ArtisanCloud/go-libs/objects"
	"regexp"
	"strings"
	"unicode"
)

var (
	/**
	 * The cache of snake-cased words.
	 *
	 * @var array
	 */
	SnakeCache StringMap

	/**
	 * The cache of camel-cased words.
	 *
	 * @var array
	 */
	CamelCache StringMap

	/**
	 * The cache of studly-cased words.
	 *
	 * @var array
	 */
	StudlyCache StringMap

	/**
	 * The callback that should be used to generate UUIDs.
	 *
	 * @var callable
	 */
	UidFactory StringMap
)

func init() {

	SnakeCache = StringMap{}
	CamelCache = StringMap{}
	StudlyCache = StringMap{}
	UidFactory = StringMap{}
}

/**
 * Convert a value to camel case.
 *
 * @param  string  value
 * @return string
 */
func Camel(value string) string {

	// if cache has converted
	if _, ok := CamelCache[value]; ok {
		return CamelCache[value]
	}

	// low case first char and store into cache
	CamelCache[value] = LCFirst(Studly(value))

	return CamelCache[value]
}

/**
 * Convert a string to snake case.
 *
 * @param  string  $value
 * @param  string  $delimiter
 * @return string
 */
func snake(value string, delimiter string) string {
	if delimiter == "" {
		delimiter = "_"
	}

	key := value + delimiter
	// if cache has converted
	if _, ok := CamelCache[key]; ok {
		return CamelCache[key]
	}

	if !IsUpper(value) {
		value = RegexpReplace("/\\s+/u", "", UCWords(value))

		value = Lower(RegexpReplace("/(.)(?=[A-Z])/u", "$1"+delimiter, value))
	}
	//
	CamelCache[key] = value
	return CamelCache[key]
}

/**
 * Convert a value to studly caps case.
 *
 * @param  string  value
 * @return string
 */
func Studly(value string) string {

	// if cache has converted
	if _, ok := StudlyCache[value]; ok {
		return StudlyCache[value]
	}

	// replace "-" or "_" with " "
	value = strings.ReplaceAll(value, "-", " ")
	value = strings.ReplaceAll(value, "_", " ")

	// Up Case words
	value = UCWords(value)

	// replace " " with "", and store into cache
	StudlyCache[value] = strings.ReplaceAll(value, " ", "")

	return StudlyCache[value]

}

/**
 * Make a string's first character lowercase
 * @param string str <p>
 * The input string.
 * </p>
 * @return string the resulting string.
 */
func LCFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

/**
 * Uppercase the first character of each word in a string
 * @param string str <p>
 * The input string.
 * </p>
 * @param string delimiters [optional] <p>
 * @return string the modified string.
 */
func UCWords(str string) string {
	return strings.Title(str)
}

/**
 * Check for uppercase character(s)
 * @param string str <p>
 * The input string.
 * </p>
 * @param string<p>
 * @return bool.
 */
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/**
 * Check for lowercase character(s)
 * @param string str <p>
 * The input string.
 * </p>
 * @param string<p>
 * @return bool.
 */
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/**
 * Convert the given string to lower-case.
 *
 * @param  string  $value
 * @return string
 */
func Lower(value string) string {
	return strings.ToLower(value)
}

/**
 * Convert the given string to upper-case.
 *
 * @param  string  $value
 * @return string
 */
func Upper(value string) string {
	return strings.ToUpper(value)
}

/**
 * Replace by regex
 *
 * @param  string  $value
 * @return string
 */
func RegexpReplace(pattern string, replacement string, subject string) string {
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllString(subject, replacement)

}