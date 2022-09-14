package controllers

import (
	"ams/database"
	"ams/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const key string = "Content-Type"
const value string = "application/json"

func checkIfAssetExists(assetId string) bool {

	var asset entities.Asset

	database.Instance.First(&asset, assetId)

	if asset.AssetId == 0 {

		return false

	}

	return true
}

// CurrentUser godoc
// @Summary  Gives the information about the current user
// @Schemes
// @Description  Gives the information about the current user.
// @Tags         write
// @Accept       json
// @Produce      json
// @Success      200
//@Param id path int true "asset ID"
//@ID get-asset-by-id
// @Router       /api/getasset/{id}/ [get]
func GetAssetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(key, value)
	assetId := mux.Vars(r)["id"]
	if checkIfAssetExists(assetId) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Asset Not Found!")
		return
	}
	w.WriteHeader(200)
	var asset entities.Asset
	database.Instance.First(&asset, assetId)

	json.NewEncoder(w).Encode(asset)
}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	assetId := mux.Vars(r)["id"]
	if checkIfAssetExists(assetId) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Asset Not Found!")
		return
	}
	var asset entities.Asset
	database.Instance.First(&asset, assetId)
	json.NewDecoder(r.Body).Decode(&asset)
	database.Instance.Save(&asset)
	w.Header().Set(key, value)
	json.NewEncoder(w).Encode(asset)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	empId := mux.Vars(r)["id"]
	assetId := mux.Vars(r)["asset_id"]
	w.Header().Set(key, value)
	var emp entities.Employee
	database.Instance.First(&emp, empId)
	json.NewDecoder(r.Body).Decode(&emp)
	database.Instance.Model(&entities.Employee{}).Where(&emp).Update("asset_id", assetId)
	database.Instance.Save(emp)
	json.NewEncoder(w).Encode(map[string]string{"status": "true", "message": "data updated successfully!"})

}

// Register godoc
// @Summary  creates new Employee
// @Schemes
// @Description  Creates the new Employee.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param        input body  entities.Employee  true  "New Employee"
// @Success      200
// @Router       /api/createemployee [post]
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(key, value)

	var employee entities.Employee
	var response entities.ResponseEmployee

	json.NewDecoder(r.Body).Decode(&employee)

	errs := []string{}

	if employee.EmpName == "" {

		errs = append(errs, fmt.Errorf("EmpName is required").Error())

		json.NewEncoder(w).Encode(errs)

		return

	}

	if employee.Address == "" {

		errs = append(errs, fmt.Errorf("EmpAddress is required").Error())

		json.NewEncoder(w).Encode(errs)

		return

	}

	if CheckDuplicateEmpId(employee.EmpId) == true {

		json.NewEncoder(w).Encode("duplicated entry found!")

		return

	}

	if err := database.Instance.Create(&employee).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//response.Status = false
		//response.Messagae = "Duplicate Entry Found!"
		json.NewEncoder(w).Encode("Duplicate entry")
		return
	}

	response.Status = true
	response.Messagae = "User insert Successfully!"
	response.Data.EmpId = employee.EmpId
	response.Data.EmpName = employee.EmpName
	response.Data.Address = employee.Address
	response.Data.Phone = employee.Phone
	response.Data.AssetId = employee.AssetId
	json.NewEncoder(w).Encode(response)

}

// CurrentUser godoc
// @Summary  Gives the information about all the  user
// @Schemes
// @Description  Gives the information about all the user.
// @Tags         write
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/getusers/ [get]
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	var users []entities.User
	database.Instance.Find(&users)
	w.Header().Set(key, value)
	json.NewEncoder(w).Encode(users)

}

// Create godoc
// @Summary  creates new user
// @Schemes
// @Description  Create new user.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param        input body  entities.User  true  "New User"
// @Success      200
// @Router      /api/usercreate [post]
func Users(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(key, value)
	var user entities.User
	var response entities.ResponseUser
	json.NewDecoder(r.Body).Decode(&user)

	if err := database.Instance.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Status = false
		response.Messagae = "Duplicated Entity Found!"
		response.Data.UserId = user.UserId
		response.Data.UserPassword = user.UserPassword

		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = true
	response.Messagae = "User insert Successfully!"
	response.Data.UserId = user.UserId
	response.Data.UserPassword = user.UserPassword
	json.NewEncoder(w).Encode(response)
}

// Create godoc
// @Summary  creates new asset
// @Schemes
// @Description  Created new Asset.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param        input body  entities.Asset  true  "New Asset"
// @Success      200
// @Router       /api/createassets [post]
func Asset(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(key, value)
	var asset entities.Asset
	var response entities.ResponseAsset
	json.NewDecoder(r.Body).Decode(&asset)

	errs := []string{}

	if asset.AssetDescription == "" {

		errs = append(errs, fmt.Errorf("Asset Description is required").Error())

		json.NewEncoder(w).Encode(errs)

		return

	}

	if CheckDuplicateEmpId(int(asset.AssetId)) == true {

		json.NewEncoder(w).Encode("duplicated entry found!")

		return

	}

	if err := database.Instance.Create(&asset).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Status = false
		response.Messagae = "Duplicae entry"
		response.Data.AssetId = asset.AssetId
		response.Data.AssetDescription = asset.AssetDescription

		json.NewEncoder(w).Encode(response)
		return
	}
	response.Status = true
	response.Messagae = "Asset inserted Successfully!"
	response.Data.AssetId = asset.AssetId
	response.Data.AssetDescription = asset.AssetDescription

	json.NewEncoder(w).Encode(response)
}

func CheckDuplicateEmpId(id int) bool {

	var emp entities.Employee

	database.Instance.First(&emp, id)

	if emp.EmpId != 0 {

		return true

	}

	return false

}
