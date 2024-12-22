package domain

type Project struct {
	Name  string `json:"name"`
	Owner string `json:"userId"`
	// TODO: implement Task entity
	Tasks []string `json:"tasks"`
}

// TODO: create tests for methods
func (p *Project) SetName(name string) {
	p.Name = name
}

func (p *Project) CreateTask(task string) {
	p.Tasks = append(p.Tasks, task)
}
