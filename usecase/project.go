package usecase

import (
	"errors"

	"gitnub.com/premwut/todo-service/domain"
	"gitnub.com/premwut/todo-service/model"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	Find(projectId string) (*domain.Project, error)
	CreateTask(projectId string, taskName string) (domain.Task, error)
}

type ProjectService struct {
	// projectRepository ProjectRepositoryInterface
	// TODO: figure out how to define interface for this DB object
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	s := ProjectService{db}
	return &s
}

func (s *ProjectService) GetProject(projectId string) (*model.Project, error) {
	var project model.Project
	var err error
	result := s.db.First(&project, projectId) // * Get with the row's id

	// * Example: Find with condition (WHERE statement)
	// * result := s.db.Where("id = ?", projectId).Find(&project)
	if result.Error != nil {
		err = errors.New("project not found")
	}

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) CreateProject(name string) (*model.Project, error) {
	project := model.Project{Name: name, UserID: 1}
	s.db.Create(&project)

	return &project, nil
}

// TODO: implement creating and save task to DB
func (s *ProjectService) CreateTask(projectId string, taskName string) (model.Task, error) {
	// var task, err = s.projectRepository.CreateTask(projectId, taskName)
	var err error
	task := model.Task{}

	if err != nil {
		return model.Task{}, err
	}

	return task, nil

}
