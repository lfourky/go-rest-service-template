package domain

type User struct {
	PrimaryKey
	Name  string
	Email string
	Items []Item
}
