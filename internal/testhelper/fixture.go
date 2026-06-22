package testhelper

import (
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/gophercloud/gophercloud/v2"
)

func DefaultConfig() config.Config {
	return config.Config{
		Port:        "9696",
		PublicURL:   "",
		Region:      "RegionOne",
		Username:    "admin",
		Password:    "password",
		UserID:      "admin",
		ProjectID:   "demo",
		ProjectName: "demo",
		RoleName:    "admin",
	}
}

func PasswordAuthOptions() *gophercloud.AuthOptions {
	return &gophercloud.AuthOptions{
		Username:   "admin",
		Password:   "password",
		DomainID:   "default",
		TenantName: "demo",
	}
}

func ServiceClient(url string) *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{},
		Endpoint:       url + "/",
	}
}
