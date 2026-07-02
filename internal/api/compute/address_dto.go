package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type serverAddressListResponse struct {
	Addresses map[string][]serverAddressDocument `json:"addresses"`
}

type serverAddressDocument struct {
	Version int    `json:"version"`
	Address string `json:"addr"`
}

func toServerAddressDocumentsByNetwork(
	addresses map[string][]appcompute.ServerAddress,
) map[string][]serverAddressDocument {
	documents := make(map[string][]serverAddressDocument, len(addresses))
	for network, values := range addresses {
		documents[network] = toServerAddressDocuments(values)
	}

	return documents
}

func toServerAddressDocuments(
	addresses []appcompute.ServerAddress,
) []serverAddressDocument {
	documents := make([]serverAddressDocument, 0, len(addresses))
	for _, address := range addresses {
		documents = append(documents, serverAddressDocument{
			Version: address.Version,
			Address: address.Address,
		})
	}

	return documents
}
