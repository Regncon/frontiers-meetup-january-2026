package models

var TodoViewModeStrings = []string{"All", "Active", "Completed"}

type Todo struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TodoViewMode int

type TodoPageState struct {
	Todos      []*Todo      `json:"todos"`
	EditingIdx int          `json:"editingIdx"`
	Mode       TodoViewMode `json:"mode"`
}
