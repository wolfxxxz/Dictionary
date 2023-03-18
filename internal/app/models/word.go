package models

//Word model
type Word struct {
	ID         int    `json:"id"`
	English    string `json:"english"`
	Russian    string `json:"russian"`
	Theme      string `json:"theme"`
	RightAswer int    `json:"rightAnswer"`
}
