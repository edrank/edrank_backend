package types

type (
	// Common
	ChangePasswordBody struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	GetLinkingDataBody struct {
	}

	OnBoardCollegeBody struct {
		OnboardingFile string   `json:"onboarding_file"`
		FileRegistryId string   `json:"file_registry_id"`
		Courses        []string `json:"courses"`
	}

	GetOnboardingSheetBody struct {
		College struct {
			Name           string `json:"name"`
			Email          string `json:"email"`
			Phone          string `json:"phone" `
			Website        string `json:"website "`
			UniversityName string `json:"university_name"`
			CollegeType    string `json:"collge_type"`
			City           string `json:"city"`
			State          string `json:"state"`
		} `json:"college"`
		Courses []string `json:"courses"`
	}

	CreateNewCollegeAdminBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	Top3TeachersBody struct {
		RequestType string `json:"request_type"`
		Cid         int    `json:"cid"`
		City        string `json:"city"`
		State       string `json:"state"`
		N           int    `json:"n"`
	}
)
