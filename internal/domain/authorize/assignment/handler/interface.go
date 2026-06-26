package handler

import "github.com/hoerilahyar/go-clean/internal/domain/authorize/assignment/usecase"

type AssignmentHandler struct {
	usecase usecase.AssignmentUsecase
}

func NewAssignmentHandler(
	assignmentUsecase usecase.AssignmentUsecase,
) *AssignmentHandler {
	return &AssignmentHandler{
		usecase: assignmentUsecase,
	}
}
