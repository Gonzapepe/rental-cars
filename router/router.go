package router

import (
	"net/http"
	"strconv"

	"github.com/Gonzapepe/cars-rental/controller"
	"github.com/Gonzapepe/cars-rental/helper"
	"github.com/Gonzapepe/cars-rental/model"
	"github.com/Gonzapepe/cars-rental/repository"
	"github.com/gin-gonic/gin"
)

func NewRouter(carRepository *repository.CarRepository) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	baseRouter := router.Group("/api")
	carRouter := baseRouter.Group("/cars")
	// makeRouter := baseRouter.Group("/makers")
	
	// CAR ROUTERS 
	carRouter.POST("/create", func(context *gin.Context) {
		var car model.Car

		err := context.ShouldBindJSON(&car)

		if err != nil {
			response := helper.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := controller.CreateCar(&car, *carRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		context.JSON(code, response)

		
	})

	carRouter.GET("/", func(context *gin.Context) {
		code := http.StatusOK

		response := controller.FindAllCars(*carRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	carRouter.GET("/show/:id", func(context *gin.Context) {
		id, err :=strconv.Atoi(context.Param("id"))
		
		if err != nil {
			response := helper.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := controller.FindOneCarById(id, *carRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		context.JSON(code, response)
	})

	carRouter.PUT("/update/:id", func(context *gin.Context) {
		id, err :=strconv.Atoi(context.Param("id"))
		
		if err != nil {
			response := helper.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}
		
		var car model.Car
		
		err = context.ShouldBindJSON(&car)

		if err != nil {
			response := helper.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := controller.UpdateCarById(id, &car, *carRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	carRouter.DELETE("/delete/:id", func(context *gin.Context) {
		id, err :=strconv.Atoi(context.Param("id"))
		
		if err != nil {
			response := helper.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}
		
		code := http.StatusOK

		response := controller.DeleteCarById(id, *carRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	carRouter.GET("/pagination", func(context *gin.Context) {
		code := http.StatusOK

		pagination := helper.GeneratePaginationRequest(context)

		response := controller.Pagination(*carRepository, context, pagination)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
	return router
}
