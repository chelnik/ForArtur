package main

import (
	"fmt"
)

func main() {
	disabledKeys := make(map[string]bool)

	// Обработка входных данных
	input := []string{"DISABLE 1000 2", "DISABLE 2 1", "DISABLE 3 3", "GETMAX", "RESET 1", "RESET 2", "DISABLE 1 2", "DISABLE 1 3", "DISABLE 2 2", "GETMAX", "RESET 3"}
	for _, cmd := range input {
		var arg1, arg2 string
		fmt.Sscanf(cmd, "%s %s %s", &arg1, &arg2)
		switch arg1 {
		case "DISABLE":
			disabledKeys[arg2] = true
			disabledKeys[fmt.Sprintf("%s %s", arg2, arg2)] = false
			fmt.Println(arg2, arg2)
		case "RESET":
			for key := range disabledKeys {
				if key == arg2 || key == fmt.Sprintf("%s %s", arg2, arg2) {
					delete(disabledKeys, key)
				}
			}
		case "GETMAX":
			maxCount := 0
			for key, val := range disabledKeys {
				if val {
					continue
				}
				count := 1
				if key[0:1] == key[2:3] {
					count = 2
				}
				if count > maxCount {
					maxCount = count
				}
			}
			fmt.Println(maxCount)
		}
	}
}
