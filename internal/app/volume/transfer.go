package volume

import "errors"

var ErrTransferNotFound = errors.New("transfer not found")

func (s *Service) CreateTransfer(input CreateTransfer) Transfer {
	transfer := Transfer{
		ID:        "transfer-" + s.idGen.Hex(16),
		AuthKey:   "auth-" + s.idGen.Hex(16),
		Name:      input.Name,
		VolumeID:  input.VolumeID,
		CreatedAt: s.clock.Now().UTC().Format(timestampFormat),
		Links:     []map[string]string{},
	}

	return s.transferRepository.Create(transfer)
}

func (s *Service) ListTransfers() []Transfer {
	return s.transferRepository.List()
}

func (s *Service) GetTransfer(id string) (Transfer, error) {
	return s.transferRepository.Get(id)
}

func (s *Service) DeleteTransfer(id string) error {
	return s.transferRepository.Delete(id)
}

func (s *Service) AcceptTransfer(id string, authKey string) (Transfer, error) {
	transfer, err := s.transferRepository.Get(id)
	if err != nil {
		return Transfer{}, err
	}
	if transfer.AuthKey != authKey {
		return Transfer{}, ErrTransferNotFound
	}
	if err := s.transferRepository.Delete(id); err != nil {
		return Transfer{}, err
	}

	return transfer, nil
}
