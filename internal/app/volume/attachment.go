package volume

import "errors"

var ErrAttachmentNotFound = errors.New("attachment not found")

func (s *Service) CreateAttachment(input CreateAttachment) Attachment {
	now := s.clock.Now().UTC().Format(timestampFormat)
	mode := input.Mode
	if mode == "" {
		mode = "rw"
	}

	attachment := Attachment{
		ID:             "att-" + s.idGen.Hex(16),
		VolumeID:       input.VolumeID,
		Instance:       input.InstanceID,
		AttachedAt:     now,
		Status:         "reserved",
		AttachMode:     mode,
		ConnectionInfo: map[string]any{},
		Connector:      input.Connector,
	}
	if len(input.Connector) > 0 {
		attachment.Status = "attaching"
	}

	return s.attachmentRepository.Create(attachment)
}

func (s *Service) ListAttachments() []Attachment {
	return s.attachmentRepository.List()
}

func (s *Service) GetAttachment(id string) (Attachment, error) {
	return s.attachmentRepository.Get(id)
}

func (s *Service) UpdateAttachment(
	id string,
	input UpdateAttachment,
) (Attachment, error) {
	attachment, err := s.attachmentRepository.Get(id)
	if err != nil {
		return Attachment{}, err
	}

	attachment.Connector = input.Connector
	attachment.ConnectionInfo = map[string]any{
		"driver_volume_type": "sandstack",
	}
	attachment.Status = "attaching"

	return s.attachmentRepository.Update(attachment)
}

func (s *Service) CompleteAttachment(id string) error {
	attachment, err := s.attachmentRepository.Get(id)
	if err != nil {
		return err
	}

	attachment.Status = "attached"
	_, err = s.attachmentRepository.Update(attachment)

	return err
}

func (s *Service) DeleteAttachment(id string) error {
	attachment, err := s.attachmentRepository.Get(id)
	if err != nil {
		return err
	}

	attachment.Status = "detached"
	attachment.DetachedAt = s.clock.Now().UTC().Format(timestampFormat)

	return s.attachmentRepository.Delete(id)
}
