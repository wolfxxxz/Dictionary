package testink

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
)

// Scan bufio
func Scan() string {
	fmt.Print("    ")
	in := bufio.NewScanner(os.Stdin)
	var nn string
	if in.Scan() {
		nn = in.Text()
	}
	return nn
}

// Сравнение строк / пробелы между словами "_"
func Compare(l models.Word) (yes int, not int) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := ""
	//Игнорировать пробелы
	for _, v := range l.English {
		if v != ' ' {
			c = c + string(v)
		}
	}
	var a string
	s := ""
	//Mistake-----------------------------------------------------------
	a = Scan()
	for _, v := range a {
		if v != ' ' {
			s = s + string(v)
		}
	}

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not
	/* Если захочется игнорировать одну ошибку в слове
	if moreThanOneMistake(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not*/
}

/*
// Игнорировать одну ошибку в словах
func moreThanOneMistake(first, second string) bool {
	first = strings.ToLower(first)
	second = strings.ToLower(second)

	lenFirst, lenSecond := len(first), len(second)
	if strings.EqualFold(first, second) {
		//fmt.Println("100% duplicates")
		return true
	}

	if (lenFirst-lenSecond) >= 2 || (lenSecond-lenFirst) >= 2 {
		//fmt.Println("try more, over time")
		return false
	}

	if (lenSecond - lenFirst) == 1 {
		return quantityMistakes(first, second)
	}

	if (lenFirst - lenSecond) == 1 {
		return quantityMistakes(second, first)
	}

	return true
}

func quantityMistakes(a, b string) bool {
	sizeMistake := 0
	for i, v := range a {
		if v != rune(b[i+sizeMistake]) {
			sizeMistake++
			if sizeMistake >= 2 {
				return false
			}
		}
	}

	return true
}*/

func ScanTime(a *string) {
	fmt.Print("    ")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		*a = in.Text()
	}
}

// Сравнение строк / пробелы между словами "_"
func CompareTime(l models.Word) (yes int, not int) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := ""
	//Игнорировать пробелы
	for _, v := range l.English {
		if v != ' ' {
			c = c + string(v)
		}
	}
	var a string
	s := ""
	//Mistake-----------------------------------------------------------
	go ScanTime(&a)
	time.Sleep(10 * time.Second)
	for _, v := range a {
		if v != ' ' {
			s = s + string(v)
		}
	}

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not
	/* Если захочется игнорировать одну ошибку в слове
	if moreThanOneMistake(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not*/
}
