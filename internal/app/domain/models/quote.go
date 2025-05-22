package models

type Quote struct {
	Id      int    `db:"id"`
	Author  string `db:"author"`
	Content string `db:"content"`
}
