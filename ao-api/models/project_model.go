package models

type Project struct {
	Id          int    `db:"id" json:"-"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	AccountId   string `db:"account_id" json:"-"`
}
