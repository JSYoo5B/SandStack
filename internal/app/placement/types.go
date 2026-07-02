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
