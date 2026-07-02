package volume

type AttachVolume struct {
	InstanceUUID string
	HostName     string
	MountPoint   string
	Mode         string
}

type DetachVolume struct {
	AttachmentID string
}

type ResetVolumeStatus struct {
	Status       string
	AttachStatus string
}

func (s *Service) Attach(id string, input AttachVolume) error {
	volume, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	volume.Status = "in-use"
	volume.UpdatedAt = s.clock.Now().UTC().Format(timestampFormat)
	_, err = s.repository.Update(volume)
	return err
}

func (s *Service) BeginDetaching(id string) error {
	return s.setVolumeStatus(id, "detaching")
}

func (s *Service) Detach(id string, input DetachVolume) error {
	return s.setVolumeStatus(id, "available")
}

func (s *Service) Reserve(id string) error {
	return s.setVolumeStatus(id, "reserved")
}

func (s *Service) Unreserve(id string) error {
	return s.setVolumeStatus(id, "available")
}

func (s *Service) ExtendSize(id string, newSize int) error {
	volume, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	volume.Size = newSize
	volume.UpdatedAt = s.clock.Now().UTC().Format(timestampFormat)
	_, err = s.repository.Update(volume)
	return err
}

func (s *Service) ResetStatus(id string, input ResetVolumeStatus) error {
	return s.setVolumeStatus(id, input.Status)
}

func (s *Service) setVolumeStatus(id string, status string) error {
	volume, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	volume.Status = status
	volume.UpdatedAt = s.clock.Now().UTC().Format(timestampFormat)
	_, err = s.repository.Update(volume)
	return err
}
