package task

type TaskCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AssigneeID  string `json:"assignee_id"`
}
