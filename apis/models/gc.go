package models

import (
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	GCInputsModel struct {
		Id           int       `json:"id"`
		TenantId     int       `json:"tenant_id"`
		TenantType   string    `json:"tenant_type"`
		ComplaintFor string    `json:"complaint_for"`
		Subject      string    `json:"subject"`
		IsCC         bool      `json:"is_cc"`
		CCResponse   string    `json:"cc_response"`
		ProofFileId  int       `json:"proof_file_id"`
		Description  string    `json:"description"`
		VictimId     int       `json:"victim_id"`
		VictimType   string    `json:"victim_type"`
		IsActive     bool      `json:"is_active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func CreateNewGCInput(gc GCInputsModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into gc_inputs (tenant_id,tenant_type,complaint_for,subject,is_cc,cc_response,proof_file_id,description,victim_id,victim_type,is_active) values (?,?,?,?,?,?,?,?,?,?,?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(gc.TenantId, gc.TenantType, gc.ComplaintFor, gc.Subject, gc.IsCC, gc.CCResponse, gc.ProofFileId, gc.Description, gc.VictimId, gc.VictimType, gc.IsActive)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	id, err := resp.LastInsertId()

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	return int(id), nil

}
