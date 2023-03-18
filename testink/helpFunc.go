package testink

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
)

// Scan bufio
func Scan() string {
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
	//Mistake
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
}
