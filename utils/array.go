package utils

import (
	"fmt"
	"strings"
)

func FindIndex(array []string, match string) int {
	index := 0
	for _, a := range array {
		if a == match {
			return index
		}
		index = index + 1
	}
	return -1
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ArrayIntoSqlValues(array []string) string {
	quotedItems := make([]string, len(array))

	// Iterate over the items and surround each item with double quotes
	for i, item := range array {
		quotedItems[i] = fmt.Sprintf(`'%s'`, item)
	}

	// Join the quoted items with commas using strings.Join
	result := strings.Join(quotedItems, ", ")
	return result
}

func ArrayIntoSqlColumns(array []string) string {
	quotedItems := make([]string, len(array))

	// Iterate over the items and surround each item with double quotes
	for i, item := range array {
		quotedItems[i] = fmt.Sprintf(`"%s"`, item)
	}

	// Join the quoted items with commas using strings.Join
	result := strings.Join(quotedItems, ", ")
	return result
}
