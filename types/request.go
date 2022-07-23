package types

type (
	// Common
	ChangePasswordTypes struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	ForgetPasswordTypes struct {
		Email string
	}
)
