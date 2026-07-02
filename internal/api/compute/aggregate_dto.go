package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type createAggregateRequest struct {
	Aggregate createAggregateDocument `json:"aggregate"`
}

type createAggregateDocument struct {
	Name             string `json:"name"`
	AvailabilityZone string `json:"availability_zone"`
}

type aggregateListResponse struct {
	Aggregates []aggregateDocument `json:"aggregates"`
}

type aggregateResponse struct {
	Aggregate aggregateDocument `json:"aggregate"`
}

type aggregateDocument struct {
	AvailabilityZone string            `json:"availability_zone"`
	Hosts            []string          `json:"hosts"`
	ID               int               `json:"id"`
	Metadata         map[string]string `json:"metadata"`
	Name             string            `json:"name"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	DeletedAt        *string           `json:"deleted_at"`
	Deleted          bool              `json:"deleted"`
	UUID             string            `json:"uuid"`
}

func (request createAggregateRequest) createAggregate() appcompute.CreateAggregate {
	return appcompute.CreateAggregate{
		Name:             request.Aggregate.Name,
		AvailabilityZone: request.Aggregate.AvailabilityZone,
	}
}

func toAggregateDocuments(
	aggregates []appcompute.Aggregate,
) []aggregateDocument {
	documents := make([]aggregateDocument, 0, len(aggregates))
	for _, aggregate := range aggregates {
		documents = append(documents, toAggregateDocument(aggregate))
	}

	return documents
}

func toAggregateDocument(aggregate appcompute.Aggregate) aggregateDocument {
	var deletedAt *string
	if aggregate.DeletedAt != "" {
		deletedAt = &aggregate.DeletedAt
	}

	return aggregateDocument{
		AvailabilityZone: aggregate.AvailabilityZone,
		Hosts:            aggregate.Hosts,
		ID:               aggregate.ID,
		Metadata:         aggregate.Metadata,
		Name:             aggregate.Name,
		CreatedAt:        aggregate.CreatedAt,
		UpdatedAt:        aggregate.UpdatedAt,
		DeletedAt:        deletedAt,
		Deleted:          aggregate.Deleted,
		UUID:             aggregate.UUID,
	}
}
