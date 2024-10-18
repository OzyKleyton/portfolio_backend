package service

import (
	"github.com/ozykt4/portfolio_backend/internal/model"
	"github.com/ozykt4/portfolio_backend/internal/repository"
)

type IProjectService interface {
	CreateProject(project *model.ProjectReq) *model.Response
}

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) IProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (p *ProjectService) CreateProject(project *model.ProjectReq) *model.Response {
	projectReq := project.ToProject()

	createProject, err := p.repo.Create(projectReq)
	if err != nil {
		return model.NewErrorResponse(err, 400)
	}

	return model.NewSuccessResponse(createProject.ToProjectRes())
}
