package controller

import (
	"fmt"

	"github.com/Gonzapepe/cars-rental/dtos"
	"github.com/Gonzapepe/cars-rental/model"
	"github.com/Gonzapepe/cars-rental/repository"
	"github.com/gin-gonic/gin"
)

func CreateCar(car *model.Car, repository repository.CarRepository) dtos.Response {
	operationResult := repository.Save(car)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Car)

	return dtos.Response{Success: true, Data: data}
}

func FindAllCars(repository repository.CarRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Car)

	return dtos.Response{Success: true, Data: data}
}

func FindOneCarById(id int, repository repository.CarRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Car)

	return dtos.Response{Success: true, Data: data}
}

func UpdateCarById(id int, car *model.Car, repository repository.CarRepository) dtos.Response {
	existingCarResponse := FindOneCarById(id, repository)

	if !existingCarResponse.Success {
		return existingCarResponse
	}

	existingCar := existingCarResponse.Data.(*model.Car)

	existingCar.Model = car.Model
	existingCar.Drive = car.Drive
	existingCar.FuelType = car.FuelType
	existingCar.Make = car.Make
	existingCar.MakeID = car.MakeID
	existingCar.Transmission = car.Transmission
	existingCar.Year = car.Year

	operationResult := repository.Save(existingCar)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	
	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteCarById(id int, repository repository.CarRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func Pagination(repository repository.CarRepository, context *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := repository.Pagination(pagination)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	// Get current url path
	urlPath := context.Request.URL.Path

	// Search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return dtos.Response{Success: true, Data: data}
}