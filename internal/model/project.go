package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Link        string
	Description string
}

type ProjectReq struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func (p *ProjectReq) ToProject() *Project {
	return &Project{
		Name:        p.Name,
		Link:        p.Link,
		Description: p.Description,
	}
}

type ProjectRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func (pr *Project) ToProjectRes() *ProjectRes {
	return &ProjectRes{
		ID:          pr.ID,
		Name:        pr.Name,
		Link:        pr.Link,
		Description: pr.Description,
	}
}
