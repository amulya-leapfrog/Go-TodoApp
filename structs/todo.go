package structs

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type CreateTodo struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}