package dep

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	DueDateTime string `json:"dueDateTime"`
	Time        int64  `json:"time"`
}
