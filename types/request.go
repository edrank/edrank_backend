package types

type (
	// Common
	ChangePasswordBody struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	GetLinkingDataBody struct {
		
	}
)
