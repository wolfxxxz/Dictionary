package testink

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
	"github.com/Wolfxxxz/Dictionary/library"
)

// Сравнение слов
func Work(s []*models.Word) (yes int, not int, wordsIncorrect []*models.Word, wordsRight []*models.Word) {
	fmt.Println("                     START")
	for _, v := range s {
		y, n := Compare(*v)
		if y > 0 {
			yes++
			v.RightAswer += 1
			wordsRight = append(wordsRight, v)
		} else if n > 0 {
			not++
			wordsIncorrect = append(wordsIncorrect, v)
		}
	}
	return yes, not, wordsIncorrect, wordsRight
}

// repair mistakes
func WorkMistake(s []*models.Word) {
	fmt.Println("                 Work with mistakes")
	//fmt.Println("--- Check and repair ---")
	for {
		if len(s) == 0 {
			break
		}
		v := s[len(s)-1]
		y, _ := Compare(*v)

		if y > 0 && len(s) != 1 {
			s = s[:len(s)-1]
			fmt.Println(len(s))
		} else if len(s) == 1 && y > 0 {
			s = []*models.Word{}
		}
	}
}

// Тест по количеству
func TestKnowlig(l []*models.Word) {
	//fmt.Println("Flag-----------------------------------------------")
	log.Println("                       Start")

	//Scan quantity words for test
	//----------------------------------------------
	fmt.Println("Количество слов для теста")
	var quantity int
	lenLibrary := len(l)
	for {
		cc := Scan()
		i, err := strconv.Atoi(cc)
		if err != nil {
			fmt.Println("Incorect, please enter number")
		} else if i >= lenLibrary {
			fmt.Printf("Incorect, please enter less number. Len Library is: %v\n", lenLibrary)
		} else {
			quantity = i
			break
		}
	}

	//Test_1
	// Cute some
	TestWords := l[:quantity]
	s, e, incorectWords, rightWords := Work(TestWords)
	if len(incorectWords) >= 1 {
		library.Print(incorectWords)
		//Test_2
		WorkMistake(incorectWords)
		fmt.Println(s, e)
	} else {
		fmt.Println("    БЕЗ ОШИБОК !!!")
		fmt.Printf(" Right answers is: %v\n", s)
	}
	//write used words in the end
	rttt := l[quantity:]
	rttt = append(rttt, rightWords...)
	incorectWords = append(incorectWords, rttt...)
	//Сохранить в txt file
	library.SaveTXT(incorectWords, "txt/library.txt")
	//Сохранить в json file
	library.Savejson(incorectWords, "txt/library.json")

	fmt.Println("  All the words in a dictionary: ", len(incorectWords)+1)
	log.Println("Final")
}

// Тест по темам
func ThemesOfWords(l []*models.Word) {
	//Упорядочить по теме
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Theme > l[j].Theme
	})
	//library.Savejson(l, "library.json")
	library.SaveTXT(l, "txt/library.txt")

	var quantity int = 1
	for i, v := range l {
		if i == len(l)-1 {
			fmt.Print("Осталось заполнить: ", quantity)
			break
		} else if v.Theme == l[i+1].Theme {
			quantity++
		} else if v.Theme != l[i+1].Theme {
			fmt.Print(v.Theme, ": ", quantity, " ")
			quantity = 1
		}
	}
	fmt.Println()

	fmt.Println("Для теста введите название темы")
	themes := Scan()
	ThemeSlice := []*models.Word{}
	if themes == "" {
		fmt.Println("You are lazy")
	} else {
		for _, v := range l {
			if v.Theme == themes {
				ThemeSlice = append(ThemeSlice, v)
			}
		}

		ThemeSlice = library.MixUpTwo(ThemeSlice)
		if len(ThemeSlice) >= 20 {
			ThemeSlice = ThemeSlice[:20]
		}

		s, e, incorectWords, _ := Work(ThemeSlice)
		if len(incorectWords) >= 1 {

			library.Print(incorectWords)
			//Test_2

			fmt.Println(s, e)
		} else {
			fmt.Println("    БЕЗ ОШИБОК !!!")
			fmt.Printf(" Right answers is: %v\n", s)
		}
	}
}
