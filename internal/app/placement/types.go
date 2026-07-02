package placement

type ResourceProvider struct {
	UUID               string
	Name               string
	Generation         int
	ParentProviderUUID string
	RootProviderUUID   string
}

type CreateResourceProvider struct {
	UUID               string
	Name               string
	ParentProviderUUID string
}

type Inventory struct {
	ResourceClass   string
	AllocationRatio float32
	MaxUnit         int
	MinUnit         int
	Reserved        int
	StepSize        int
	Total           int
}

type Inventories struct {
	ResourceProviderGeneration int
	Inventories                map[string]Inventory
}

type UpdateInventories struct {
	ResourceProviderGeneration int
	Inventories                map[string]Inventory
}

type UpdateInventory struct {
	ResourceProviderGeneration int
	Inventory                  Inventory
}

type Traits struct {
	ResourceProviderGeneration int
	Traits                     []string
}

type UpdateTraits struct {
	ResourceProviderGeneration int
	Traits                     []string
}

type Aggregates struct {
	ResourceProviderGeneration *int
	Aggregates                 []string
}

type UpdateAggregates struct {
	ResourceProviderGeneration *int
	Aggregates                 []string
}

type Usages struct {
	ResourceProviderGeneration int
	Usages                     map[string]int
}

type Allocation struct {
	Resources map[string]int
}

type Allocations struct {
	ResourceProviderGeneration int
	Allocations                map[string]Allocation
}

type AllocationCandidateQuery struct {
	Resources map[string]int
}

type AllocationCandidates struct {
	AllocationRequests []AllocationRequest
	ProviderSummaries  map[string]ProviderSummary
}

type AllocationRequest struct {
	Allocations map[string]Allocation
}

type ProviderSummary struct {
	Resources          map[string]ProviderSummaryResource
	Traits             []string
	ParentProviderUUID string
	RootProviderUUID   string
}

type ProviderSummaryResource struct {
	Capacity int
	Used     int
}

type ResourceClass struct {
	Name string
}
