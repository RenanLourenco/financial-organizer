package search_income

	
import "github.com/RenanLourenco/financial-organizer/income/adapter/entity"	


type SearchIncomeInput struct {
	Search string 
}

type SearchIncomeOutput struct {
	Data []entity.Income `json:"data"`
}