package storage

type Todo struct {
	Item   string
	IfDone bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func TodoItemsInit() *PageData {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install Go", IfDone: true},
			{Item: "Learn HTMX", IfDone: false},
			{Item: "Revise Physics", IfDone: false},
		},
	}
	return &data
}
