package models

func NewQuotaFields(name string, memory int64, InstanceMemoryLimit int64, routes int, services int, nonbasicservices bool) (q QuotaFields) {
	q.Name = name
	q.MemoryLimit = memory
	q.InstanceMemoryLimit = InstanceMemoryLimit
	q.RoutesLimit = routes
	q.ServicesLimit = services
	q.NonBasicServicesAllowed = nonbasicservices
	return
}

func NewQuotaUsage(name string, memoryLimit int64, InstanceMemoryLimit int64, routesLimit int, serviceLimit int, nonbasicservices bool, routes int, services int, memory int64) (q QuotaUsage) {
	q.Name = name
	q.MemoryLimit = memoryLimit
	q.InstanceMemoryLimit = InstanceMemoryLimit
	q.RoutesLimit = routesLimit
	q.ServicesLimit = serviceLimit
	q.NonBasicServicesAllowed = nonbasicservices
	q.OrgUsage.Routes = routes
	q.OrgUsage.Services = services
	q.OrgUsage.Memory = memory
	return
}

type OrgUsageFields struct {
	Routes   int   `json:"routes"`
	Services int   `json:"services"`
	Memory   int64 `json:"memory"`
}

type QuotaUsage struct {
	QuotaFields
	OrgUsage OrgUsageFields `json:"org_usage"`
}

type QuotaFields struct {
	Guid                    string `json:"guid,omitempty"`
	Name                    string `json:"name"`
	MemoryLimit             int64  `json:"memory_limit"`          // in Megabytes
	InstanceMemoryLimit     int64  `json:"instance_memory_limit"` // in Megabytes
	RoutesLimit             int    `json:"total_routes"`
	ServicesLimit           int    `json:"total_services"`
	NonBasicServicesAllowed bool   `json:"non_basic_services_allowed"`
}
