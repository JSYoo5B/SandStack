package volume

import "errors"

var ErrVolumeTypeNotFound = errors.New("volume type not found")

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
