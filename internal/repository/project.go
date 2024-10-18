package repository

import (
	"github.com/ozykt4/portfolio_backend/internal/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *model.Project) (*model.Project, error)
}

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepo{
		db: db,
	}
}

func (p *ProjectRepo) Create(project *model.Project) (*model.Project, error) {
	if err := p.db.Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}
