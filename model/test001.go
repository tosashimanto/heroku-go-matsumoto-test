package model

// bbbb
type Test001 struct {
	ID             int64     `json:"testId"`
	RegisteredDate string    `json:"registeredDate"`
	TestArray      []Test002 `json:"testArray"`
}

// aaaa
type Test002 struct {
	Index        int64 `json:"index"`
	JudgedNumber int64 `json:"number"`
}
