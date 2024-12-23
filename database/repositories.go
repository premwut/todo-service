package database

import (
	"errors"
	"time"

	"gitnub.com/premwut/todo-service/domain"
)

var mockUsers = []domain.User{
	{
		ID:   "123",
		Name: "Jimi Hendrix",
		Age:  666,
	},
	{
		ID:   "124",
		Name: "Paul Gilbert",
		Age:  999,
	},
	{
		ID:   "125",
		Name: "Zack Wylde",
		Age:  777,
	},
	{
		ID:   "126",
		Name: "John Petrucci",
		Age:  888,
	},
}

var mockTasks = []domain.Task{
	{
		Id:     "1",
		Name:   "Create project entity",
		Status: "Created",
	},
	{
		Id:     "1",
		Name:   "Create project entity",
		Status: "Created",
	},
}

var mockProjects = []domain.Project{
	{
		Id:    "122",
		Name:  "My first proejct",
		Owner: "1",
		Tasks: mockTasks,
	},
	{
		Id:    "123",
		Name:  "My first proejct",
		Owner: "1",
		Tasks: mockTasks,
	},
}

type Database struct {
	users    []domain.User
	projects []domain.Project
}

var db = Database{
	users:    mockUsers,
	projects: mockProjects,
}

type UserRepository struct {
	// may have db instance here in real use case
}

func NewUserRepo() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUser(userId string) (*domain.User, error) {
	var err error
	time.Sleep(1 * time.Second)
	mockUser := mockUsers[0]
	return &mockUser, err
}

type ProjectRepository struct {
	db Database
}

func NewProjectRepository(db Database) ProjectRepository {
	r := ProjectRepository{db}
	return r
}

func (r ProjectRepository) Find(projectId string) (*domain.Project, error) {
	var err error
	time.Sleep(1 * time.Second)
	var project *domain.Project
	for _, p := range db.projects {
		if p.Id == projectId {
			project = &p
			break
		}
	}

	if project == nil {
		err = errors.New("project not found")
		return nil, err
	}

	return project, nil
}
