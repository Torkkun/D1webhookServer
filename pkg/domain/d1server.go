package domain

type DoorOpenRequest struct {
	Time string `json:"time"`
}

type ProjectorRequest struct {
	HandlerName  string `json:"handlerName"`
	Notification bool   `json:"notification"`
	Slot         int    `json:"slot,omitempty"`
}

type MonitorRequest struct {
	HandlerName  string `json:"handlerName"`
	Notification bool   `json:"notification"`
	Slot         int    `json:"slot,omitempty"`
}
