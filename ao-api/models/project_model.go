package models

type Project struct {
	Id          int    `db:"id" json:"-"`
	Name        string `db:"name" json:"name" binding:"regexp=^[a-zA-Z][a-zA-Z0-9_]*$,min=2,max=20"`
	Description string `db:"description" json:"description"`
	AccountId   string `db:"account_id" json:"-"`
	Tag         string `db:"tag" json:"tag"`
}
