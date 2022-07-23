package types

type (
	RequestHeaders struct {
		TenantId int `json:"tenant_id"`
		TenantType int `json:"tenant_type"`
	}
)