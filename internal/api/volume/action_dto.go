package volume

type attachVolumeRequest struct {
	MountPoint   string `json:"mountpoint"`
	InstanceUUID string `json:"instance_uuid"`
	HostName     string `json:"host_name"`
	Mode         string `json:"mode"`
}

type detachVolumeRequest struct {
	AttachmentID string `json:"attachment_id"`
}

type extendVolumeRequest struct {
	NewSize int `json:"new_size"`
}

type resetVolumeStatusRequest struct {
	Status       string `json:"status"`
	AttachStatus string `json:"attach_status"`
}
