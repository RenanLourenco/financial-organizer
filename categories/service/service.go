package service

import (
	"errors"

	"github.com/RenanLourenco/financial-organizer/categories/adapter/entity"
	"gorm.io/gorm"
)

type CategoryService struct {
	Repository *gorm.DB
}

func (c *CategoryService) Categorize(categoryDescription string) (entity.Category, error) {
	var category entity.Category

	

	switch categoryDescription {
	case "Food":
		c.Repository.Where(&entity.Category{Description: "Food"}).First(&category)

		if category.ID == 0 {

			category.Description = "Food"

			c.Repository.Create(&category)

		}

		
	case "Health":
		c.Repository.Where(&entity.Category{Description: "Health"}).First(&category)

		if category.ID == 0 {

			category.Description = "Health"

			c.Repository.Create(&category)

		}

		

	case "Housing":
		c.Repository.Where(&entity.Category{Description: "Housing"}).First(&category)

		if category.ID == 0 {

			category.Description = "Housing"

			c.Repository.Create(&category)

		}

		
	case "Transport":
		c.Repository.Where(&entity.Category{Description: "Transport"}).First(&category)

		if category.ID == 0 {

			category.Description = "Transport"

			c.Repository.Create(&category)

		}
		
	case "Education":
		c.Repository.Where(&entity.Category{Description: "Education"}).First(&category)

		if category.ID == 0 {

			category.Description = "Education"

			c.Repository.Create(&category)

		}
		
	case "Entertainment":
		c.Repository.Where(&entity.Category{Description: "Entertainment"}).First(&category)

		if category.ID == 0 {

			category.Description = "Entertainment"

			c.Repository.Create(&category)

		}
		
	case "Contingencies":
		c.Repository.Where(&entity.Category{Description: "Contingencies"}).First(&category)

		if category.ID == 0 {

			category.Description = "Contingencies"

			c.Repository.Create(&category)

		}
		
	case "Others":
		c.Repository.Where(&entity.Category{Description: "Others"}).First(&category)

		if category.ID == 0 {

			category.Description = "Others"

			c.Repository.Create(&category)

		}
		

	default:
		return entity.Category{},errors.New("Error in Categorize, verify the category")
	}

	return category, nil
}

