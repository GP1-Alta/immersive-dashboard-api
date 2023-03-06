package delivery

import "immersive-dashboard/features/mentees"

type MenteeHandler struct {
	menteeService mentees.MenteeServiceInterface
}

func New(srv mentees.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: srv,
	}
}
