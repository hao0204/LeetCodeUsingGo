package main

import "fmt"

func main() {
	var a = []string{"abdfs","cccd","a","qwwewm"}
	var d = findWords(a)
	for _, i := range d {
		fmt.Println(i)
	}
}


func findWords(words []string) []string {
	if words == nil || len(words) == 0 {
		return nil
	}
	var keyMap map[string]int
	keyMap = make(map[string]int)
	key1 := "QWERTYUIOPqwertyuiop"
	key2 := "ASDFGHJKLasdfghjkl"
	key3 := "ZXCVBNMzxcvbnm"

	k1 := []rune(key1)
	k2 := []rune(key2)
	k3 := []rune(key3)
	for i := 0; i < len(k1); i++ {
		keyMap[string(k1[i])] = 1
	}
	for i := 0; i < len(k2); i++ {
		keyMap[string(k2[i])] = 2
	}
	for i := 0; i < len(k3); i++ {
		keyMap[string(k3[i])] = 3
	}
	var result = words[:]
	count := 0
	for _, word := range words {
		flag := true
		k4 := []rune(word)
		key := keyMap[string(k4[0])]
		for i := 1; i < len(k4); i++ {
			if key != keyMap[string(k4[i])] {
				flag = false
				break;
			}
		}
		if flag {
			result[count] = word
			count++
		}
	}
	return result[:count]
}