package domain

type Project struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"userId"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"` // * Should be enum of created, in_progress, done, rejected
}

// TODO: create tests for methods
func (p *Project) SetName(name string) {
	p.Name = name
}

func (p *Project) CreateTask(taskName string) {
	task := Task{Name: taskName, Status: "created"}
	p.Tasks = append(p.Tasks, task)
}
