package model

type Test001 struct {
	Id             int64     `json:"testId"`
	RegisteredDate string    `json:"registeredDate"`
	TestArray      []Test002 `json:"testArray"`
}

type Test002 struct {
	Index        int64 `json:"index"`
	JudgedNumber int64 `json:"number"`
}
