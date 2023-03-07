package delivery

import "immersive-dashboard/features/status"

type StatusHandler struct {
	statusService status.StatusServiceInterface
}

func New(srv status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		statusService: srv,
	}
}
