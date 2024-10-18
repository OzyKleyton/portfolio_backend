package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozykt4/portfolio_backend/internal/api/router"
	"github.com/ozykt4/portfolio_backend/internal/model"
	"github.com/ozykt4/portfolio_backend/internal/service"
)

type ProjectHandler struct {
	service service.IProjectService
}

func NewProjectHandler(service service.IProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

func (ph *ProjectHandler) Routes() router.Router {
	return func(route fiber.Router) {
		project := route.Group("projects")
		project.Post("/", ph.CreateProjectHandler)
	}
}

func (ph *ProjectHandler) CreateProjectHandler(ctx *fiber.Ctx) error {
	projectReq := new(model.ProjectReq)
	if err := ctx.BodyParser(projectReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse(err, fiber.ErrBadRequest))
	}

	res := ph.service.CreateProject(projectReq)

	return ctx.Status(res.Status).JSON(res)
}
