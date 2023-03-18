package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
	"github.com/Wolfxxxz/Dictionary/library"
	"github.com/Wolfxxxz/Dictionary/testink"
)

var Rttr []*models.Word

func init() {

	//-----Зачитываем содержимое файла библиотека-----
	Rttr = library.Takejson("txt/library.json")
	/*
		rttr = library.TakeTXT("library.txt")
		library.Savejson(rttr, "library.json")
		library.WriteArr(rttr, "library.txt")
		library.SortLibraryTheme(rttr, "library.txt")
	*/
}

func main() {

	fmt.Println("Тест    знаний    введите   1")
	fmt.Println("Добавить или изменить слова 2")
	fmt.Println("Запустить сервер            3")
	f, _ := library.Scan()
	ff, _ := strconv.Atoi(f)
	if ff == 1 {
		fmt.Println("Тест слов по количеству    1")
		fmt.Println("Тест знаний по теме        2")

		f, _ := library.Scan()
		fff, _ := strconv.Atoi(f)
		if fff == 1 {
			testink.TestKnowlig(Rttr)
		} else if fff == 2 {
			testink.ThemesOfWords(Rttr)
		} else {
			fmt.Println("YOU ARE LAZY")
		}
	} else if ff == 2 {
		fmt.Println("Добавить новые слова СПИСКОМ - введите 1")
		fmt.Println("Ввести новые слова в ручном режиме     2")
		fmt.Println("Добавить тему или изменить слова       3")
		fmt.Println("Сортировать библиотеку порядок изменится  4")
		f, _ := library.Scan()
		fff, _ := strconv.Atoi(f)
		if fff == 1 {
			library.UpdateLibrary("txt/newWords.txt", Rttr)
			//library.NewWordsTXT(rttr)
		} else if fff == 2 {
			library.NewWordRukamy(Rttr)
		} else if fff == 3 {
			library.AddTheme(Rttr)
		} else if fff == 4 {
			fmt.Println("start sort")
			library.SortLibrary(Rttr, "txt/library.json")
		} else {
			fmt.Println("YOU ARE LAZY")
		}
	} else if ff == 3 {
		StartServer()
	} else {
		fmt.Println("YOU ARE LAZY")
	}
	time.Sleep(5 * time.Second)
}
