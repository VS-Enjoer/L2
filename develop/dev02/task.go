package dev02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func main(str string) (string, error) {

	//Полученную строку преобразовываем к слайсу рун
	MyArrayRune := []rune(str)
	//Сразу проверяем что в слайсе есть элементы и первая руна не равна числу
	if len(MyArrayRune) < 1 {
		return "", errors.New("Некорректная строка ")
	} else if unicode.IsDigit(MyArrayRune[0]) {
		return "", errors.New("Некорректная строка ")
	}
	//Колво итераций по слайсу которых нужно пройти
	length := len(MyArrayRune)
	// Устанавливаем флаг для проверки является ли символ цифрой
	flag := true
	//Переменная для возврата результата
	var result string
	//Переменная для проверки escape последовательности
	escapeFlag := false
	//Переменная для понимания какое число лежит в руне(если оно есть)
	var count int
	//Проверяем есть ли escape последовательности
	for _, elem := range MyArrayRune {
		if elem == '\\' {
			escapeFlag = true
		}
	}

	//Если escape последовательности нет
	if !escapeFlag {
		for i := 0; i < length; i++ {
			//Устанавливаем значение флагу (является ли он цифрой)
			flag = unicode.IsDigit(MyArrayRune[i])
			//Если нет, то мы просто добавляем элемент в резалт
			if !flag {
				result += string(MyArrayRune[i])
			} else {
				//Проверяем следующий элемент является ли цифрой
				if i+1 < length && unicode.IsDigit(MyArrayRune[i+1]) {
					return "", errors.New("Некорректный формат ")
				}
				//Узнаем сколько раз нужно элемент добавить в строку
				count = int(MyArrayRune[i] - '0')
				result += strings.Repeat(string(MyArrayRune[i-1]), count-1)
				count = 0
			}

		}

	} else {
		//Итерируемся по срезу и записываем в result руны до встречи  символа \
		for i := 0; i < length; i++ {
			flag = true
			if MyArrayRune[i] == '\\' {
				flag = false
			}
			//Когда мы его встретили начинается проверка является ли следующий элемент цифрой
			if !flag && i+1 < length && unicode.IsDigit(MyArrayRune[i+1]) {
				//Если все таки он является цифрой тогда проверяем дополнительно есть ли за цифрой еще одна цифра
				if i+2 < length && unicode.IsDigit(MyArrayRune[i+2]) {
					//Если да, тогда к переменной count добавляем эту цифру и заполняем срез предыдущей цифрой count раз
					count = int(MyArrayRune[i+2] - '0')
					result += strings.Repeat(string(MyArrayRune[i+1]), count)
					count = 0
					i = i + 2
				} else if !flag {
					if i+1 == length && unicode.IsDigit(MyArrayRune[i+1]) {
						count = int(MyArrayRune[i+1] - '0')
						result += strings.Repeat(string(MyArrayRune[i]), count)
						count = 0
						i++
					}
				} else { //Не придумал такой ситуации, ну а вдруг
					return str, errors.New("Некорректная строка ")
				}
			} else if flag {
				if i+1 == length && unicode.IsDigit(MyArrayRune[i]) {
					if unicode.IsDigit(MyArrayRune[i]) && unicode.IsDigit(MyArrayRune[i-2]) {
						result += strconv.Itoa(int(MyArrayRune[i] - '0'))
					} else if MyArrayRune[i-1] == '\\' {
						count = int(MyArrayRune[i] - '0')
						result += strings.Repeat(string(MyArrayRune[i-1]), count)
					} else if MyArrayRune[i-1] == '\\' && MyArrayRune[i-2] == '\\' {
						count = int(MyArrayRune[i] - '0')
						result += strings.Repeat(string(MyArrayRune[i-1]), count)
					}
				} else {
					result += string(MyArrayRune[i])
				}
			}
		}
	}
	return result, nil

}
