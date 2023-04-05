package model

type Notify struct {
	From  string `json:"from"`
	ID    int64  `json:"id"`
	Testo string `json:"testo"`
}
