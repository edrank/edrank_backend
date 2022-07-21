package models

type (
	CollegeAdminModel struct {
		Id int
		Cid int
		Name string
		Email string
		Password string
		CreatedAt string
		UpdatedAt string
	}
)

func (collgeAdmin *CollegeAdminModel) GetAllCollegeAdminsOfCollege() []CollegeAdminModel {
	return nil
}

func (collgeAdmin *CollegeAdminModel) GetAllCollegeAdminsById() CollegeAdminModel {
	return CollegeAdminModel{}
}

func (collgeAdmin *CollegeAdminModel) UpdateCollegeAdminById() CollegeAdminModel {
	return CollegeAdminModel{}
}