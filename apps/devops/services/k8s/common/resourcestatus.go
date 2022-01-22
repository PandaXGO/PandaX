package common

// ResourceStatus provides basic information about resources status on the list.
type ResourceStatus struct {
	// Number of resources that are currently in running state.
	Running int `json:"running"`

	// Number of resources that are currently in pending state.
	Pending int `json:"pending"`

	// Number of resources that are in failed state.
	Failed int `json:"failed"`

	// Number of resources that are in succeeded state.
	Succeeded int `json:"succeeded"`

	// Number of resources that are in unknown state.
	Unknown int `json:"unknown"`

	// Number of resources that are in terminating state.
	Terminating int `json:"terminating"`
}
