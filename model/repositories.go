package model

import (
	"errors"
	"strconv"
	"time"

	"gitnub.com/premwut/todo-service/domain"
	"gorm.io/gorm"
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
		Id:     "2",
		Name:   "Create project entity",
		Status: "Created",
	},
}

var mockProjects = []domain.Project{
	{
		Id:    "122",
		Name:  "My first project",
		Owner: "1",
		Tasks: mockTasks,
	},
	{
		Id:    "123",
		Name:  "My second project",
		Owner: "1",
		Tasks: mockTasks,
	},
}

type Database struct {
	users    []domain.User
	projects []domain.Project
}

type DatabaseInterface[Model any] interface {
	Create(value interface{}) *gorm.DB
	// FindUser(id string) domain.User
	// Find(id string) T
	// Create(data T) T
}

var db = &Database{
	users:    mockUsers,
	projects: mockProjects,
}

type UserRepository struct {
	// may have db instance here in real use case
	db *Database
}

func NewUserRepo() *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) GetUser(userId string) (*domain.User, error) {
	var err error
	time.Sleep(1 * time.Second)
	mockUser := mockUsers[0]
	return &mockUser, err
}

type ProjectRepository struct {
	db DatabaseInterface[Project]
}

func NewProjectRepository(db DatabaseInterface[Project]) *ProjectRepository {
	r := ProjectRepository{db}
	return &r
}

type TaskRepository struct {
	db DatabaseInterface[Task]
}

func NewTaskRepository(db DatabaseInterface[Task]) *TaskRepository {
	r := TaskRepository{db}
	return &r
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

func (r *ProjectRepository) CreateTask(projectId string, taskName string) (domain.Task, error) {
	project, err := r.Find(projectId)
	if project == nil {
		return domain.Task{}, errors.New("project not found")
	} else if err != nil {
		return domain.Task{}, err
	}

	taskNum := len(project.Tasks) + 1
	newTask := domain.Task{Id: strconv.Itoa(taskNum), Name: taskName, Status: "created"}
	project.Tasks = append(project.Tasks, newTask)

	// TODO: implement actual db

	return newTask, nil
}
