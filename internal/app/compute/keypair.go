package compute

import "errors"

var ErrKeyPairNotFound = errors.New("key pair not found")

func (s *Service) CreateKeyPair(input CreateKeyPair) KeyPair {
	keyType := input.Type
	if keyType == "" {
		keyType = "ssh"
	}

	publicKey := input.PublicKey
	privateKey := ""
	if publicKey == "" {
		publicKey = "ssh-rsa sandstack-generated-key"
		privateKey = "-----BEGIN PRIVATE KEY-----\nsandstack\n-----END PRIVATE KEY-----"
	}

	keyPair := KeyPair{
		Name:        input.Name,
		Fingerprint: "sandstack:" + s.idGen.Hex(8),
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
		UserID:      input.UserID,
		Type:        keyType,
	}

	return s.keyPairRepository.Create(keyPair)
}

func (s *Service) ListKeyPairs() []KeyPair {
	return s.keyPairRepository.List()
}

func (s *Service) GetKeyPair(name string) (KeyPair, error) {
	return s.keyPairRepository.Get(name)
}

func (s *Service) DeleteKeyPair(name string) error {
	return s.keyPairRepository.Delete(name)
}
