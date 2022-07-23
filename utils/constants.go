package utils

var (
	ValidTentantTypes = [...]string{"STUDENT", "TEACHER", "PARENT", "COLLEGE_ADMIN", "SUPER_ADMIN", "HEIA"}
	TenantMap         = map[string]string{
		"STUDENT":       "STUDENT",
		"TEACHER":       "TEACHER",
		"PARENT":        "PARENT",
		"COLLEGE_ADMIN": "COLLEGE_ADMIN",
		"SUPER_ADMIN":   "SUPER_ADMIN",
		"HEIA":          "HEIA",
	}
)
