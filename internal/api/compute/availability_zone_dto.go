package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type availabilityZoneListResponse struct {
	AvailabilityZoneInfo []availabilityZoneDocument `json:"availabilityZoneInfo"`
}

type availabilityZoneDocument struct {
	Hosts     map[string]map[string]availabilityZoneServiceState `json:"hosts"`
	ZoneName  string                                             `json:"zoneName"`
	ZoneState availabilityZoneState                              `json:"zoneState"`
}

type availabilityZoneState struct {
	Available bool `json:"available"`
}

type availabilityZoneServiceState struct {
	Active    bool   `json:"active"`
	Available bool   `json:"available"`
	UpdatedAt string `json:"updated_at"`
}

func toAvailabilityZoneDocuments(
	zones []appcompute.AvailabilityZone,
) []availabilityZoneDocument {
	documents := make([]availabilityZoneDocument, 0, len(zones))
	for _, zone := range zones {
		documents = append(documents, toAvailabilityZoneDocument(zone))
	}

	return documents
}

func toAvailabilityZoneDocument(
	zone appcompute.AvailabilityZone,
) availabilityZoneDocument {
	return availabilityZoneDocument{
		Hosts:    map[string]map[string]availabilityZoneServiceState{},
		ZoneName: zone.Name,
		ZoneState: availabilityZoneState{
			Available: zone.Available,
		},
	}
}
