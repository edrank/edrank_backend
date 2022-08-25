package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/edrank/edrank_backend/apis/db"
	"github.com/edrank/edrank_backend/apis/utils"
)

type (
	FileRegistryModel struct {
		Id        int       `json:"id"`
		Cid       int       `json:"cid"`
		Location  string    `json:"location"`
		Url       string    `json:"url"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func GetFileRegistriesByField(fieldName string, fieldValue any) ([]FileRegistryModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from file_registry where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return nil, errors.New("Cannot find file registry")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return nil, err
	}

	var file_registrys []FileRegistryModel
	for rows.Next() {
		var fr FileRegistryModel

		if err := rows.Scan(&fr.Id, &fr.Cid, &fr.Location, &fr.Url, &fr.Type, &fr.CreatedAt, &fr.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return nil, err
		}
		file_registrys = append(file_registrys, fr)
	}
	if len(file_registrys) == 0 {
		return nil, errors.New("Cannot find file registry")
	}
	return file_registrys, nil
}

func GetFileRegistryByField(fieldName string, fieldValue any) (FileRegistryModel, error) {
	database := db.GetDatabase()
	rows, err := database.Query(fmt.Sprintf("select * from file_registry where %s = ?", fieldName), fieldValue)
	if err == sql.ErrNoRows {
		return FileRegistryModel{}, errors.New("Cannot find file registry")
	}
	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return FileRegistryModel{}, err
	}

	var file_registrys []FileRegistryModel
	for rows.Next() {
		var fr FileRegistryModel

		if err := rows.Scan(&fr.Id, &fr.Cid, &fr.Location, &fr.Url, &fr.Type, &fr.CreatedAt, &fr.UpdatedAt); err != nil {
			utils.PrintToConsole(err.Error(), "red")
			return FileRegistryModel{}, err
		}
		file_registrys = append(file_registrys, fr)
	}
	if len(file_registrys) == 0 {
		return FileRegistryModel{}, errors.New("Cannot find file registry")
	}
	return file_registrys[0], nil
}

func CreateFileRegistry(fr FileRegistryModel) (int, error) {
	database := db.GetDatabase()
	query := "insert into file_registry (cid, location, url, type) values (?, ?, ?, ?);"

	stmt, err := database.Prepare(query)

	if err != nil {
		utils.PrintToConsole(err.Error(), "red")
		return -1, err
	}

	resp, err := stmt.Exec(fr.Cid, fr.Location, fr.Url, fr.Type)

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
