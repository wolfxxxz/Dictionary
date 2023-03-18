package library

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
)

type Slovarick struct {
	Words []*models.Word
}

/*type Library struct {
	English string `json:"english"`
	Russian string `json:"russian"`
	Theme   string `json:"theme"`
}*/

// Create Library
func NewLibrary(newId int, newEnglish string, newRussian string, newTheme string) *models.Word {
	d := &models.Word{ID: newId, English: newEnglish, Russian: newRussian, Theme: newTheme}
	return d
}

// Сохранить в json file
// Marshal
func Savejson(l []*models.Word, file string) {
	s := Slovarick{
		Words: l,
	}
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(byteArr))             0664
	err = os.WriteFile(file, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}

// Достать с json file
// Unmarshal
func Takejson(file string) []*models.Word {
	//1. Создадим файл дескриптор
	filejson, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer filejson.Close()
	//fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go

	f := make([]byte, 64)
	var data2 string

	for {
		n, err := filejson.Read(f)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data2 = data2 + string(f[:n])
	}

	data := []byte(data2)
	var users Slovarick

	// Теперь задача - перенести все из data в users - это и есть десериализация!
	json.Unmarshal(data, &users)

	Usersu := users.Words

	return Usersu
}
