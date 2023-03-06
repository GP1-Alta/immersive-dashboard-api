package delivery

import "immersive-dashboard/features/classes"

type ClassHandler struct {
	menteeService classes.ClassServiceInterface
}

func New(srv classes.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		menteeService: srv,
	}
}
