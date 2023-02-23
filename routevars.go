package routevars

import (
	"regexp"
	"strings"
)

/*
EXAMPLE:

	var path = "/users/<<id:int>>/<<name:string>>"

	// No match
	var other = "/users/1234/johna/a"
	var ok, vars = Match(path, other)
	if ok {
		fmt.Println("Matched!")
		for k, v := range vars {
			fmt.Println(k, v)
		}
	} else {
		fmt.Println("No match")
	}

	// Match
	other = "/users/1234/john"
	ok, vars = Match(path, other)
	if ok {
		fmt.Println("Matched!")
		for k, v := range vars {
			fmt.Println(k, v)
		}
	} else {
		fmt.Println("No match")
	}
*/

// Router regex delimiters
const (
	RT_PATH_VAR_PREFIX = "<<"
	RT_PATH_VAR_SUFFIX = ">>"
	RT_PATH_VAR_DELIM  = ":"
)

// Router regex types.
const (
	NameInt    = "int"
	NameString = "string"
	NameSlug   = "slug"
	NameUUID   = "uuid"
	NameAny    = "any"
	NameHex    = "hex"
)

// Router regex patterns
const (
	// Match any character
	RT_PATH_REGEX_ANY = ".+"
	// Match any number
	RT_PATH_REGEX_NUM = "[0-9]+"
	// Match any string
	RT_PATH_REGEX_STR = "[a-zA-Z]+"
	// Match any hex number
	RT_PATH_REGEX_HEX = "[0-9a-fA-F]+"
	// Match any UUID
	RT_PATH_REGEX_UUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	// Match any alphanumeric string
	RT_PATH_REGEX_ALPHANUMERIC = "[0-9a-zA-Z_-]+"
)

// var paths = make(map[string]string)

func Match(path, pathToMatch string) (bool, map[string]string) {

	// Check if the path equals, if so, return true.
	if path == pathToMatch && !isRegexRoute(path) {
		return true, nil
	}

	var regex = ""
	// Check wether the path contains trailing or leading slashes,
	// store the answer for later.
	var hasPrefixSlash = strings.HasPrefix(path, "/")
	var hasTrailingSlash = strings.HasSuffix(path, "/")

	// Remove leading slash if it exists.
	if hasPrefixSlash {
		path = path[1:]
	}

	// Remove trailing slash if it exists.
	if hasTrailingSlash && len(path) > 1 && !hasPrefixSlash {
		path = path[:len(path)-1]
	}

	// Split the path into parts.
	var parts = strings.Split(path, "/")

	// The path likely exists, since we are here.
	// We will now match the regex.
	for i, part := range parts {
		parts[i] = toRegex(part)
	}

	regex = strings.Join(parts, "/")
	if hasPrefixSlash {
		regex = "/" + regex
	}
	if hasTrailingSlash {
		regex = regex + "/"
	}

	//	paths[path] = regex

	return matchRegex(regex, pathToMatch)
}

// Check if a path is a regex route.
func isRegexRoute(path string) bool {
	return strings.Contains(path, RT_PATH_VAR_PREFIX) && strings.Contains(path, RT_PATH_VAR_SUFFIX)
}

func matchRegex(regex, pathToMatch string) (bool, map[string]string) {
	var rex = regexp.MustCompile(regex)
	var m = rex.FindStringSubmatch(pathToMatch)
	// Get named capture groups
	var vars = make(map[string]string, len(m))
	var subNames = rex.SubexpNames()
	if len(subNames) != len(m) {
		return false, nil
	}
	for i, name := range subNames {
		if i != 0 && name != "" {
			vars[name] = m[i]
		}
	}
	if len(m) > 0 && m[0] == pathToMatch {
		return true, vars
	}
	return false, nil
}

// Convert a string to a regex string with a capture group.
func toRegex(str string) string {
	if !strings.HasPrefix(str, RT_PATH_VAR_PREFIX) || !strings.HasSuffix(str, RT_PATH_VAR_SUFFIX) {
		return str
	}
	str = strings.TrimPrefix(str, RT_PATH_VAR_PREFIX)
	str = strings.TrimSuffix(str, RT_PATH_VAR_SUFFIX)
	var parts = strings.Split(str, RT_PATH_VAR_DELIM)
	if len(parts) == 1 {
		return "(?P<" + parts[0] + ">" + typToRegx(parts[0]) + ")"
	} else if len(parts) != 2 {
		return str
	}
	var groupName = parts[0]
	var typ = parts[1]
	return "(?P<" + groupName + ">" + typToRegx(typ) + ")"
}

// Convert a type (string) to a regex for use in capture groups.
func typToRegx(typ string) string {
	// regex for raw is: raw(REGEX)
	var hasRaw string = strings.ToLower(typ)
	if strings.HasPrefix(hasRaw, "raw(") && strings.HasSuffix(hasRaw, ")") {
		return hasRaw[4 : len(hasRaw)-1]
	}
	switch typ {
	case NameInt:
		return RT_PATH_REGEX_NUM
	case NameString, NameSlug:
		return RT_PATH_REGEX_ALPHANUMERIC
	case NameUUID:
		return RT_PATH_REGEX_UUID
	case NameAny:
		return RT_PATH_REGEX_ANY
	case NameHex:
		return RT_PATH_REGEX_HEX
	default:
		return RT_PATH_REGEX_STR
	}
}
