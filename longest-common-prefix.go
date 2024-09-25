package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Берем первую строку в качестве начального префикса
	prefix := strs[0]

	// Проходим по всем остальным строкам
	for _, s := range strs[1:] {
		// Пока текущая строка не содержит префикс, сокращаем префикс
		for len(prefix) > 0 && len(s) < len(prefix) || s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		// Если префикс стал пустым, выходим
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	// Пример 1
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1)) // "fl"

	// Пример 2
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs2)) // ""
}
