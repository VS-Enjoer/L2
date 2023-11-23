package task

import (
	"sort"
	"strings"
)

func FindAnagram(array []string) map[string][]string {
	//Мапа для возвращения из функции (по заданию)
	anagramMap := make(map[string][]string)
	//Бежим по массиву полученному
	for _, val := range array {
		//Приводим слово к нижнему регистру и сортируем его буквы/символы
		sortWords := sortString(strings.ToLower(val))
		//Проверяем есть ли такой ключ
		if set, flag := anagramMap[sortWords]; flag {
			//Если ключ у нас такой есть то мы добавляем слово в него
			anagramMap[sortWords] = append(set, strings.ToLower(val))
		} else {
			anagramMap[sortWords] = []string{strings.ToLower(val)}
		}
	}
	//Если 1 элемент в множестве, тогда удаляем его
	for key, val := range anagramMap {
		if len(val) <= 1 {
			delete(anagramMap, key)
		} else { //Иначе сортируем множества по ключу
			sort.Slice(val, func(i, j int) bool {
				return val[i] < val[j]
			})
			anagramMap[val[0]] = val
			// Если ключ изменился, удаляем старый
			if key != val[0] {
				delete(anagramMap, key)
			}
		}
	}
	//Возвращаем мапу
	return anagramMap
}

// Сортируем буквы/символы рун(одного элемента массива)
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
