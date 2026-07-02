package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type availabilityZoneListResponse struct {
	AvailabilityZoneInfo []availabilityZoneDocument `json:"availabilityZoneInfo"`
}

type availabilityZoneDocument struct {
	ZoneName  string                `json:"zoneName"`
	ZoneState availabilityZoneState `json:"zoneState"`
}

type availabilityZoneState struct {
	Available bool `json:"available"`
}

func toAvailabilityZoneDocuments(
	zones []appvolume.AvailabilityZone,
) []availabilityZoneDocument {
	documents := make([]availabilityZoneDocument, 0, len(zones))
	for _, zone := range zones {
		documents = append(documents, toAvailabilityZoneDocument(zone))
	}

	return documents
}

func toAvailabilityZoneDocument(
	zone appvolume.AvailabilityZone,
) availabilityZoneDocument {
	return availabilityZoneDocument{
		ZoneName: zone.Name,
		ZoneState: availabilityZoneState{
			Available: zone.Available,
		},
	}
}
