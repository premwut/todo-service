package usecase

import (
	"gitnub.com/premwut/todo-service/domain"
)

type ProjectRepositoryInterface interface {
	Find(projectId string) (*domain.Project, error)
}

type ProjectService struct {
	projectRepository ProjectRepositoryInterface
}

func NewProjectService(r ProjectRepositoryInterface) *ProjectService {
	s := ProjectService{r}
	return &s
}

func (s *ProjectService) GetProject(projectId string) (*domain.Project, error) {
	var project, err = s.projectRepository.Find(projectId)

	if project == nil {
		return project, err
	}

	return project, nil
}
