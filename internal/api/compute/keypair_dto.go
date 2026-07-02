package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type createKeyPairRequest struct {
	KeyPair createKeyPairDocument `json:"keypair"`
}

type createKeyPairDocument struct {
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
}

func (r createKeyPairRequest) createKeyPair() appcompute.CreateKeyPair {
	return appcompute.CreateKeyPair{
		Name:      r.KeyPair.Name,
		UserID:    r.KeyPair.UserID,
		Type:      r.KeyPair.Type,
		PublicKey: r.KeyPair.PublicKey,
	}
}

type keyPairListResponse struct {
	KeyPairs []keyPairListDocument `json:"keypairs"`
}

type keyPairListDocument struct {
	KeyPair keyPairDocument `json:"keypair"`
}

type keyPairResponse struct {
	KeyPair keyPairDocument `json:"keypair"`
}

type keyPairDocument struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key,omitempty"`
	UserID      string `json:"user_id"`
	Type        string `json:"type"`
}

func toKeyPairListDocuments(
	keyPairs []appcompute.KeyPair,
) []keyPairListDocument {
	documents := make([]keyPairListDocument, 0, len(keyPairs))
	for _, keyPair := range keyPairs {
		documents = append(documents, keyPairListDocument{
			KeyPair: toKeyPairDocument(keyPair),
		})
	}

	return documents
}

func toKeyPairDocument(keyPair appcompute.KeyPair) keyPairDocument {
	return keyPairDocument{
		Name:        keyPair.Name,
		Fingerprint: keyPair.Fingerprint,
		PublicKey:   keyPair.PublicKey,
		PrivateKey:  keyPair.PrivateKey,
		UserID:      keyPair.UserID,
		Type:        keyPair.Type,
	}
}
