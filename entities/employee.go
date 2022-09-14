package entities

type Employee struct {
	EmpId   int    `json:"emp_id" binding:"Required"`
	EmpName string `json:"emp_name" binding:"Required"`
	Address string `json:"address" binding:"Required"`
	Phone   int64  `json:"phone" binding:"Required"`
	AssetId int64  `json:"asset_id" binding:"Required"`
}

type ResponseEmployee struct {
	Status   bool
	Messagae string
	Data     Employee
}
