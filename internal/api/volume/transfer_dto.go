package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type createTransferRequest struct {
	Transfer createTransferDocument `json:"transfer"`
}

type createTransferDocument struct {
	Name     string `json:"name"`
	VolumeID string `json:"volume_id"`
}

func (r createTransferRequest) createTransfer() appvolume.CreateTransfer {
	return appvolume.CreateTransfer{
		Name:     r.Transfer.Name,
		VolumeID: r.Transfer.VolumeID,
	}
}

type acceptTransferRequest struct {
	Accept struct {
		AuthKey string `json:"auth_key"`
	} `json:"accept"`
}

type transferListResponse struct {
	Transfers []transferDocument `json:"transfers"`
}

type transferResponse struct {
	Transfer transferDocument `json:"transfer"`
}

type transferDocument struct {
	ID        string              `json:"id"`
	AuthKey   string              `json:"auth_key,omitempty"`
	Name      string              `json:"name"`
	VolumeID  string              `json:"volume_id"`
	CreatedAt string              `json:"created_at"`
	Links     []map[string]string `json:"links"`
}

func toTransferDocuments(transfers []appvolume.Transfer) []transferDocument {
	documents := make([]transferDocument, 0, len(transfers))
	for _, transfer := range transfers {
		documents = append(documents, toTransferDocument(transfer))
	}

	return documents
}

func toTransferDocument(transfer appvolume.Transfer) transferDocument {
	return transferDocument{
		ID:        transfer.ID,
		AuthKey:   transfer.AuthKey,
		Name:      transfer.Name,
		VolumeID:  transfer.VolumeID,
		CreatedAt: transfer.CreatedAt,
		Links:     transfer.Links,
	}
}
