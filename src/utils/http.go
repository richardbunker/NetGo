package utils

import (
	"strings"
)



func ExtractPathParam(urlPath string, resource string) string {
	var pathParts []string = strings.Split(urlPath, "/")
	if index := pathContainsResource(pathParts, resource); index >= 0 && index < len(pathParts) {
		return pathParts[index]
	}
	return ""
}


func pathContainsResource(parts []string, val string) int {
	for index, item := range parts {
		if item == val {
			return index + 1
		}
	}
	return -1
}