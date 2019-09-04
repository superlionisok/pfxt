package models

type ResultModel struct {
	Status int
	Msg    string
	Count  int
	Data   interface{}
}
