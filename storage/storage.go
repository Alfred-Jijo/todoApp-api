package storage

type Todo struct {
	Item   string
	IfDone bool
}

type PageData struct {
	Title string
	Todos []Todo
}
