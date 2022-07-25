package types

type (
	DatabaseConfig struct {
		DbName     string `json:"db_name"`
		DbHost     string `json:"db_host"`
		DbUser     string `json:"db_user"`
		DbPassword string `json:"db_pass"`
	}
)
