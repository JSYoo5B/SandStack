package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type limitsResponse struct {
	Limits limitsDocument `json:"limits"`
}

type limitsDocument struct {
	Absolute absoluteLimitsDocument `json:"absolute"`
}

type absoluteLimitsDocument struct {
	MaxTotalCores           int `json:"maxTotalCores"`
	MaxImageMeta            int `json:"maxImageMeta"`
	MaxServerMeta           int `json:"maxServerMeta"`
	MaxPersonality          int `json:"maxPersonality"`
	MaxPersonalitySize      int `json:"maxPersonalitySize"`
	MaxTotalKeypairs        int `json:"maxTotalKeypairs"`
	MaxSecurityGroups       int `json:"maxSecurityGroups"`
	MaxSecurityGroupRules   int `json:"maxSecurityGroupRules"`
	MaxServerGroups         int `json:"maxServerGroups"`
	MaxServerGroupMembers   int `json:"maxServerGroupMembers"`
	MaxTotalFloatingIps     int `json:"maxTotalFloatingIps"`
	MaxTotalInstances       int `json:"maxTotalInstances"`
	MaxTotalRAMSize         int `json:"maxTotalRAMSize"`
	TotalCoresUsed          int `json:"totalCoresUsed"`
	TotalInstancesUsed      int `json:"totalInstancesUsed"`
	TotalFloatingIpsUsed    int `json:"totalFloatingIpsUsed"`
	TotalRAMUsed            int `json:"totalRAMUsed"`
	TotalSecurityGroupsUsed int `json:"totalSecurityGroupsUsed"`
	TotalServerGroupsUsed   int `json:"totalServerGroupsUsed"`
}

func toLimitsDocument(limits appcompute.Limits) limitsDocument {
	return limitsDocument{
		Absolute: absoluteLimitsDocument{
			MaxTotalCores:           limits.MaxTotalCores,
			MaxImageMeta:            limits.MaxImageMeta,
			MaxServerMeta:           limits.MaxServerMeta,
			MaxPersonality:          limits.MaxPersonality,
			MaxPersonalitySize:      limits.MaxPersonalitySize,
			MaxTotalKeypairs:        limits.MaxTotalKeypairs,
			MaxSecurityGroups:       limits.MaxSecurityGroups,
			MaxSecurityGroupRules:   limits.MaxSecurityGroupRules,
			MaxServerGroups:         limits.MaxServerGroups,
			MaxServerGroupMembers:   limits.MaxServerGroupMembers,
			MaxTotalFloatingIps:     limits.MaxTotalFloatingIps,
			MaxTotalInstances:       limits.MaxTotalInstances,
			MaxTotalRAMSize:         limits.MaxTotalRAMSize,
			TotalCoresUsed:          limits.TotalCoresUsed,
			TotalInstancesUsed:      limits.TotalInstancesUsed,
			TotalFloatingIpsUsed:    limits.TotalFloatingIpsUsed,
			TotalRAMUsed:            limits.TotalRAMUsed,
			TotalSecurityGroupsUsed: limits.TotalSecurityGroupsUsed,
			TotalServerGroupsUsed:   limits.TotalServerGroupsUsed,
		},
	}
}
