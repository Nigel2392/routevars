package routevars

import (
	"fmt"
	"strings"
)

type URLFormatter string

func (uf URLFormatter) Format(v ...any) string {
	var path = string(uf)
	// If the length of the path is less than the length of the pre/suffix and the delimiter
	// then there are no variables in the path
	if len(path) <= len(RT_PATH_VAR_DELIM)+len(RT_PATH_VAR_PREFIX)+len(RT_PATH_VAR_SUFFIX) {
		return path
	}
	// Remove the first and last slash if they exist
	var hasPrefixSlash = strings.HasPrefix(path, "/")
	var hasTrailingSlash = strings.HasSuffix(path, "/")
	if hasPrefixSlash {
		path = path[1:]
	}
	if hasTrailingSlash {
		path = path[:len(path)-1]
	}
	// Split the path into parts
	var parts = strings.Split(path, "/")
	// Replace the parts that are variables with the arguments
	for i, part := range parts {
		if strings.HasPrefix(part, RT_PATH_VAR_PREFIX) && strings.HasSuffix(part, RT_PATH_VAR_SUFFIX) {
			if len(v) == 0 {
				panic("not enough arguments for URL: " + path)
			}
			var arg = v[0]
			v = v[1:]
			parts[i] = fmt.Sprintf("%v", arg)
		}
	}
	// Join the parts back together
	path = strings.Join(parts, "/")
	// Add the slashes back if they were there
	if hasPrefixSlash {
		path = "/" + path
	}
	if hasTrailingSlash {
		path = path + "/"
	}
	return path
}
