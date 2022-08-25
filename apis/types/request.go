package types

import "time"

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

	GetFeedbackTeachersBody struct {
		CourseId int `json:"course_id"`
		Year     int `json:"year"`
	}

	STSubmitFeedBackBody struct {
		DriveId   int `json:"drive_id"`
		Feedbacks []struct {
			TeacherID int `json:"teacher_id"`
			Mcq       []struct {
				QuestionId int `json:"question_id"`
				AnswerId   int `json:"answer_id"`
			} `json:"mcq"`
			TextFeedback string `json:"text_feedback"`
		} `json:"feedbacks"`
	}
	SCSubmitFeedBackBody struct {
		DriveId  int `json:"drive_id"`
		Feedback struct {
			Mcq []struct {
				QuestionId int `json:"question_id"`
				AnswerId   int `json:"answer_id"`
			} `json:"mcq"`
			TextFeedback string `json:"text_feedback"`
		} `json:"feedback"`
	}

	PCSubmitFeedBackBody struct {
		DriveId   int `json:"drive_id"`
		CollegeId int `json:"college_id"`
		Feedback  struct {
			Mcq []struct {
				QuestionId int `json:"question_id"`
				AnswerId   int `json:"answer_id"`
			} `json:"mcq"`
			TextFeedback string `json:"text_feedback"`
		} `json:"feedback"`
	}

	AndroidSubmitFeedbackBody struct {
		DriveId   int `json:"drive_id"`
		TeacherID int `json:"teacher_id"`
		CollegeId int `json:"college_id"`
		Feedback  struct {
			Mcq []struct {
				QuestionId int `json:"question_id"`
				AnswerId   int `json:"answer_id"`
			} `json:"mcq"`
			TextFeedback string `json:"text_feedback"`
		} `json:"feedback"`
	}

	HCSubmitFeedBackBody struct {
	}

	KBCGraphBody struct {
		TeacherId  int `json:"teacher_id"`
		QuestionId int `json:"question_id"`
	}

	SAGraphBody struct {
		CollegeId int `json:"college_id"`
		TeacherId int `json:"teacher_id"`
	}

	OnboardingAPIBody struct {
		Students []struct {
			Name             string    `json:"name"`
			Email            string    `json:"email"`
			Phone            string    `json:"phone"`
			CourseId         int       `json:"course_id"`
			Year             int       `json:"year"`
			Batch            string    `json:"batch"`
			Password         string    `json:"password"`
			EnrollmentNumber string    `json:"enrollment"`
			Dob              time.Time `json:"dob"`
			FathersName      string    `json:"fathers_name"`
			MotherName       string    `json:"mother_name"`
			GuardianEmail    string    `json:"guardian_email"`
			GuardianPhone    string    `json:"guardian_phone"`
			IsActive         bool      `json:"is_active"`
		}
		Teachers []struct {
			Name           string `json:"name"`
			OfficialEmail  string `json:"email"`
			AlternateEmail string `json:"alt_email"`
			Department     string `json:"department"`
			CourseId       int    `json:"course_id"`
			Designation    string `json:"designation"`
			Password       string `json:"password"`
			IsActive       bool   `json:"is_active"`
		}
	}

	GCSubmitBody struct {
		ComplaintFor string `json:"complaint_for"`
		Subject string `json:"subject"`
		IsCC bool `json:"is_cc"`
		CCResponse string `json:"cc_response"`
		ProofFileId int `json:"proof_file_id"`
		Description string `json:"description"`
	}
)
