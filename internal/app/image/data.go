package image

import "time"

func (s *Service) UploadData(id string, data []byte) error {
	image, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	now := s.clock.Now().UTC()
	image.Status = "active"
	image.SizeBytes = int64(len(data))
	image.UpdatedAt = now.Format(time.RFC3339)

	if _, err := s.repository.Update(image); err != nil {
		return err
	}
	s.dataRepository.Put(id, data)

	return nil
}

func (s *Service) StageData(id string, data []byte) error {
	image, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	now := s.clock.Now().UTC()
	image.Status = "uploading"
	image.SizeBytes = int64(len(data))
	image.UpdatedAt = now.Format(time.RFC3339)

	if _, err := s.repository.Update(image); err != nil {
		return err
	}
	s.dataRepository.Put(id, data)

	return nil
}

func (s *Service) ImportData(id string) error {
	image, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	now := s.clock.Now().UTC()
	image.Status = "active"
	image.UpdatedAt = now.Format(time.RFC3339)

	_, err = s.repository.Update(image)
	return err
}

func (s *Service) DownloadData(id string) ([]byte, error) {
	if _, err := s.repository.Get(id); err != nil {
		return nil, err
	}

	return s.dataRepository.Get(id)
}
