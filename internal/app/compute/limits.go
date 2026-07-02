package compute

const (
	defaultMaxTotalCores         = 200
	defaultMaxImageMeta          = 128
	defaultMaxServerMeta         = 128
	defaultMaxPersonality        = 5
	defaultMaxPersonalitySize    = 10240
	defaultMaxTotalKeypairs      = 100
	defaultMaxSecurityGroups     = 10
	defaultMaxSecurityGroupRules = 20
	defaultMaxServerGroups       = 10
	defaultMaxServerGroupMembers = 10
	defaultMaxTotalFloatingIps   = 50
	defaultMaxTotalInstances     = 100
	defaultMaxTotalRAMSize       = 512000
)

func (s *Service) GetLimits() Limits {
	servers := s.serverRepository.List()

	return Limits{
		MaxTotalCores:           defaultMaxTotalCores,
		MaxImageMeta:            defaultMaxImageMeta,
		MaxServerMeta:           defaultMaxServerMeta,
		MaxPersonality:          defaultMaxPersonality,
		MaxPersonalitySize:      defaultMaxPersonalitySize,
		MaxTotalKeypairs:        defaultMaxTotalKeypairs,
		MaxSecurityGroups:       defaultMaxSecurityGroups,
		MaxSecurityGroupRules:   defaultMaxSecurityGroupRules,
		MaxServerGroups:         defaultMaxServerGroups,
		MaxServerGroupMembers:   defaultMaxServerGroupMembers,
		MaxTotalFloatingIps:     defaultMaxTotalFloatingIps,
		MaxTotalInstances:       defaultMaxTotalInstances,
		MaxTotalRAMSize:         defaultMaxTotalRAMSize,
		TotalCoresUsed:          s.totalServerCores(servers),
		TotalInstancesUsed:      len(servers),
		TotalFloatingIpsUsed:    0,
		TotalRAMUsed:            s.totalServerRAM(servers),
		TotalSecurityGroupsUsed: 0,
		TotalServerGroupsUsed:   0,
	}
}

func (s *Service) totalServerCores(servers []Server) int {
	total := 0
	for _, server := range servers {
		flavor, err := s.GetFlavor(server.FlavorID)
		if err == nil {
			total += flavor.VCPUs
		}
	}

	return total
}

func (s *Service) totalServerRAM(servers []Server) int {
	total := 0
	for _, server := range servers {
		flavor, err := s.GetFlavor(server.FlavorID)
		if err == nil {
			total += flavor.RAM
		}
	}

	return total
}
