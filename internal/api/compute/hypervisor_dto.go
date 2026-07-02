package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type hypervisorListResponse struct {
	Hypervisors []hypervisorDocument `json:"hypervisors"`
}

type hypervisorResponse struct {
	Hypervisor hypervisorDocument `json:"hypervisor"`
}

type hypervisorStatisticsResponse struct {
	HypervisorStatistics hypervisorStatisticsDocument `json:"hypervisor_statistics"`
}

type hypervisorDocument struct {
	ID                 string                    `json:"id"`
	CPUInfo            hypervisorCPUInfoDocument `json:"cpu_info"`
	CurrentWorkload    int                       `json:"current_workload"`
	Status             string                    `json:"status"`
	State              string                    `json:"state"`
	DiskAvailableLeast int                       `json:"disk_available_least"`
	HostIP             string                    `json:"host_ip"`
	FreeDiskGB         int                       `json:"free_disk_gb"`
	FreeRAMMB          int                       `json:"free_ram_mb"`
	HypervisorHostname string                    `json:"hypervisor_hostname"`
	HypervisorType     string                    `json:"hypervisor_type"`
	HypervisorVersion  int                       `json:"hypervisor_version"`
	LocalGB            int                       `json:"local_gb"`
	LocalGBUsed        int                       `json:"local_gb_used"`
	MemoryMB           int                       `json:"memory_mb"`
	MemoryMBUsed       int                       `json:"memory_mb_used"`
	RunningVMs         int                       `json:"running_vms"`
	Service            hypervisorServiceDocument `json:"service"`
	VCPUs              int                       `json:"vcpus"`
	VCPUsUsed          int                       `json:"vcpus_used"`
}

type hypervisorCPUInfoDocument struct {
	Vendor   string                        `json:"vendor"`
	Arch     string                        `json:"arch"`
	Model    string                        `json:"model"`
	Features []string                      `json:"features"`
	Topology hypervisorCPUTopologyDocument `json:"topology"`
}

type hypervisorCPUTopologyDocument struct {
	Cells   int `json:"cells"`
	Sockets int `json:"sockets"`
	Cores   int `json:"cores"`
	Threads int `json:"threads"`
}

type hypervisorServiceDocument struct {
	ID             string `json:"id"`
	Host           string `json:"host"`
	DisabledReason string `json:"disabled_reason"`
}

type hypervisorStatisticsDocument struct {
	Count              int `json:"count"`
	CurrentWorkload    int `json:"current_workload"`
	DiskAvailableLeast int `json:"disk_available_least"`
	FreeDiskGB         int `json:"free_disk_gb"`
	FreeRAMMB          int `json:"free_ram_mb"`
	LocalGB            int `json:"local_gb"`
	LocalGBUsed        int `json:"local_gb_used"`
	MemoryMB           int `json:"memory_mb"`
	MemoryMBUsed       int `json:"memory_mb_used"`
	RunningVMs         int `json:"running_vms"`
	VCPUs              int `json:"vcpus"`
	VCPUsUsed          int `json:"vcpus_used"`
}

func toHypervisorDocuments(
	hypervisors []appcompute.Hypervisor,
) []hypervisorDocument {
	documents := make([]hypervisorDocument, 0, len(hypervisors))
	for _, hypervisor := range hypervisors {
		documents = append(documents, toHypervisorDocument(hypervisor))
	}

	return documents
}

func toHypervisorDocument(
	hypervisor appcompute.Hypervisor,
) hypervisorDocument {
	return hypervisorDocument{
		ID:                 hypervisor.ID,
		CPUInfo:            defaultCPUInfo(),
		CurrentWorkload:    hypervisor.CurrentWorkload,
		Status:             hypervisor.Status,
		State:              hypervisor.State,
		DiskAvailableLeast: hypervisor.DiskAvailableLeast,
		HostIP:             hypervisor.HostIP,
		FreeDiskGB:         hypervisor.FreeDiskGB,
		FreeRAMMB:          hypervisor.FreeRAMMB,
		HypervisorHostname: hypervisor.Hostname,
		HypervisorType:     hypervisor.Type,
		HypervisorVersion:  hypervisor.Version,
		LocalGB:            hypervisor.LocalGB,
		LocalGBUsed:        hypervisor.LocalGBUsed,
		MemoryMB:           hypervisor.MemoryMB,
		MemoryMBUsed:       hypervisor.MemoryMBUsed,
		RunningVMs:         hypervisor.RunningVMs,
		Service: hypervisorServiceDocument{
			ID:   hypervisor.ServiceID,
			Host: hypervisor.Hostname,
		},
		VCPUs:     hypervisor.VCPUs,
		VCPUsUsed: hypervisor.VCPUsUsed,
	}
}

func toHypervisorStatisticsDocument(
	hypervisors []appcompute.Hypervisor,
) hypervisorStatisticsDocument {
	stats := hypervisorStatisticsDocument{Count: len(hypervisors)}
	for _, hypervisor := range hypervisors {
		stats.CurrentWorkload += hypervisor.CurrentWorkload
		stats.DiskAvailableLeast += hypervisor.DiskAvailableLeast
		stats.FreeDiskGB += hypervisor.FreeDiskGB
		stats.FreeRAMMB += hypervisor.FreeRAMMB
		stats.LocalGB += hypervisor.LocalGB
		stats.LocalGBUsed += hypervisor.LocalGBUsed
		stats.MemoryMB += hypervisor.MemoryMB
		stats.MemoryMBUsed += hypervisor.MemoryMBUsed
		stats.RunningVMs += hypervisor.RunningVMs
		stats.VCPUs += hypervisor.VCPUs
		stats.VCPUsUsed += hypervisor.VCPUsUsed
	}

	return stats
}

func defaultCPUInfo() hypervisorCPUInfoDocument {
	return hypervisorCPUInfoDocument{
		Vendor:   "SandStack",
		Arch:     "x86_64",
		Model:    "compat",
		Features: []string{},
		Topology: hypervisorCPUTopologyDocument{
			Cells:   1,
			Sockets: 1,
			Cores:   1,
			Threads: 1,
		},
	}
}
