package controller

import (
	"net/http"

	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/RenanLourenco/financial-organizer/resume/usecase/month_resume"
	"github.com/gin-gonic/gin"
)

func Resume(c *gin.Context) {
	usecase := month_resume.MonthResume{
		Repository: database.DB,
	}

	var input month_resume.MonthResumeInput

	month := c.Params.ByName("month")
	year := c.Params.ByName("year")

	input.Month = month
	input.Year = year

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,output)

}