package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type createAttachmentRequest struct {
	Attachment struct {
		VolumeUUID   string         `json:"volume_uuid"`
		InstanceUUID string         `json:"instance_uuid"`
		Connector    map[string]any `json:"connector"`
		Mode         string         `json:"mode"`
	} `json:"attachment"`
}

type updateAttachmentRequest struct {
	Attachment struct {
		Connector map[string]any `json:"connector"`
	} `json:"attachment"`
}

type attachmentActionRequest struct {
	Complete any `json:"os-complete"`
}

type attachmentListResponse struct {
	Attachments []attachmentDocument `json:"attachments"`
}

type attachmentResponse struct {
	Attachment attachmentDocument `json:"attachment"`
}

type attachmentDocument struct {
	ID             string         `json:"id"`
	VolumeID       string         `json:"volume_id"`
	Instance       string         `json:"instance"`
	AttachedAt     string         `json:"attached_at"`
	DetachedAt     string         `json:"detached_at,omitempty"`
	Status         string         `json:"status"`
	AttachMode     string         `json:"attach_mode"`
	ConnectionInfo map[string]any `json:"connection_info"`
}

func (r createAttachmentRequest) createAttachment() appvolume.CreateAttachment {
	return appvolume.CreateAttachment{
		VolumeID:   r.Attachment.VolumeUUID,
		InstanceID: r.Attachment.InstanceUUID,
		Connector:  r.Attachment.Connector,
		Mode:       r.Attachment.Mode,
	}
}

func (r updateAttachmentRequest) updateAttachment() appvolume.UpdateAttachment {
	return appvolume.UpdateAttachment{
		Connector: r.Attachment.Connector,
	}
}

func toAttachmentDocuments(
	attachments []appvolume.Attachment,
) []attachmentDocument {
	documents := make([]attachmentDocument, 0, len(attachments))
	for _, attachment := range attachments {
		documents = append(documents, toAttachmentDocument(attachment))
	}

	return documents
}

func toAttachmentDocument(
	attachment appvolume.Attachment,
) attachmentDocument {
	return attachmentDocument{
		ID:             attachment.ID,
		VolumeID:       attachment.VolumeID,
		Instance:       attachment.Instance,
		AttachedAt:     attachment.AttachedAt,
		DetachedAt:     attachment.DetachedAt,
		Status:         attachment.Status,
		AttachMode:     attachment.AttachMode,
		ConnectionInfo: attachment.ConnectionInfo,
	}
}
