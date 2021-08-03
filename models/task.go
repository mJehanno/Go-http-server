package model

type Task struct {
	Id       int
	Name     string
	Priority int
	Done     bool
	Archived bool
}
