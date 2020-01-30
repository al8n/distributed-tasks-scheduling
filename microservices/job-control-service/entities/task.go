package entities


type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Expression string  `json:"expression"`
}
