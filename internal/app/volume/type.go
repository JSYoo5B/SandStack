package volume

import "errors"

var ErrVolumeTypeNotFound = errors.New("volume type not found")
var ErrExtraSpecNotFound = errors.New("extra spec not found")

func (s *Service) ListVolumeTypes() []VolumeType {
	volumeTypes := make([]VolumeType, 0, len(s.volumeTypes))
	volumeTypes = append(volumeTypes, s.volumeTypes...)

	return volumeTypes
}

func (s *Service) GetVolumeType(id string) (VolumeType, error) {
	for _, volumeType := range s.volumeTypes {
		if volumeType.ID == id {
			return volumeType, nil
		}
	}

	return VolumeType{}, ErrVolumeTypeNotFound
}

func (s *Service) ListExtraSpecs(volumeTypeID string) (map[string]string, error) {
	volumeType, err := s.GetVolumeType(volumeTypeID)
	if err != nil {
		return nil, err
	}

	return copyStringMap(volumeType.ExtraSpecs), nil
}

func (s *Service) GetExtraSpec(
	volumeTypeID string,
	key string,
) (map[string]string, error) {
	extraSpecs, err := s.ListExtraSpecs(volumeTypeID)
	if err != nil {
		return nil, err
	}

	value, ok := extraSpecs[key]
	if !ok {
		return nil, ErrExtraSpecNotFound
	}

	return map[string]string{key: value}, nil
}

func (s *Service) CreateExtraSpecs(
	volumeTypeID string,
	extraSpecs map[string]string,
) (map[string]string, error) {
	volumeType, index, err := s.findVolumeType(volumeTypeID)
	if err != nil {
		return nil, err
	}
	if volumeType.ExtraSpecs == nil {
		volumeType.ExtraSpecs = map[string]string{}
	}

	for key, value := range extraSpecs {
		volumeType.ExtraSpecs[key] = value
	}
	s.volumeTypes[index] = volumeType

	return copyStringMap(volumeType.ExtraSpecs), nil
}

func (s *Service) UpdateExtraSpec(
	volumeTypeID string,
	extraSpec map[string]string,
) (map[string]string, error) {
	volumeType, index, err := s.findVolumeType(volumeTypeID)
	if err != nil {
		return nil, err
	}
	if volumeType.ExtraSpecs == nil {
		volumeType.ExtraSpecs = map[string]string{}
	}

	for key, value := range extraSpec {
		volumeType.ExtraSpecs[key] = value
		s.volumeTypes[index] = volumeType
		return map[string]string{key: value}, nil
	}

	return nil, ErrExtraSpecNotFound
}

func (s *Service) DeleteExtraSpec(volumeTypeID string, key string) error {
	volumeType, index, err := s.findVolumeType(volumeTypeID)
	if err != nil {
		return err
	}
	if _, ok := volumeType.ExtraSpecs[key]; !ok {
		return ErrExtraSpecNotFound
	}

	delete(volumeType.ExtraSpecs, key)
	s.volumeTypes[index] = volumeType
	return nil
}

func (s *Service) findVolumeType(id string) (VolumeType, int, error) {
	for index, volumeType := range s.volumeTypes {
		if volumeType.ID == id {
			return volumeType, index, nil
		}
	}

	return VolumeType{}, 0, ErrVolumeTypeNotFound
}

func copyStringMap(values map[string]string) map[string]string {
	copied := map[string]string{}
	for key, value := range values {
		copied[key] = value
	}

	return copied
}
