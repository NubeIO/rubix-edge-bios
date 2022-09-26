package model

type SystemCtlStatus struct {
	Status string `json:"status"`
}

type SystemCtlActiveStatus struct {
	Active bool   `json:"active"`
	Status string `json:"status"`
}

type SystemCtlProperty struct {
	Property string `json:"property"`
}

type SystemCtlState struct {
	State bool `json:"state"`
}
