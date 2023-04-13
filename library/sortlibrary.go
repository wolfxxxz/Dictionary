package library

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
)

var LibraryTXT string = "txt/library.txt"

func SortLibraryAnswer(rttr []*models.Word, i string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].RightAswer < rttr[j].RightAswer
	})
	SaveTXT(rttr, LibraryTXT)
	Savejson(rttr, i)
}

func SortLibraryTheme(rttr []*models.Word, i string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].Theme < rttr[j].Theme
	})
	SaveTXT(rttr, LibraryTXT)
	Savejson(rttr, i)
}

func SortLibraryEnglisch(rttr []*models.Word, i string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].English < rttr[j].English
	})
	SaveTXT(rttr, LibraryTXT)
	Savejson(rttr, i)
}

func SortLibraryRussian(rttr []*models.Word, i string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].Russian < rttr[j].Russian
	})
	SaveTXT(rttr, LibraryTXT)
	Savejson(rttr, i)
}

func MixUp(rttr []*models.Word, i string) {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(rttr))
	for i, j := range perm {
		rttr[i], rttr[j] = rttr[j], rttr[i]
	}
	SaveTXT(rttr, LibraryTXT)
	Savejson(rttr, i)
}

// i file.txt
func SortLibrary(l []*models.Word, i string) {
	fmt.Println("sort Theme 1 || sort Englisch 2 || sort Russian 3 || mix up 4 || sort RightAnswer 5")
	c, _ := Scan()
	cc, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println("Incorect, please enter number")
	}
	if cc == 1 {
		SortLibraryTheme(l, i)
	} else if cc == 2 {
		SortLibraryEnglisch(l, i)
	} else if cc == 3 {
		SortLibraryRussian(l, i)
	} else if cc == 4 {
		MixUp(l, i)
	} else if cc == 5 {
		SortLibraryAnswer(l, i)
	} else {
		fmt.Println("You are lazy")
	}
}

func MixUpTwo(l []*models.Word) []*models.Word {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(l))
	for i, j := range perm {
		l[i], l[j] = l[j], l[i]
	}
	return l
}

func MixUpTime(rttr []*models.Word) {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(rttr))
	for i, j := range perm {
		rttr[i], rttr[j] = rttr[j], rttr[i]
	}
}
