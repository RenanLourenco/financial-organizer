package list_incomes

import "github.com/RenanLourenco/financial-organizer/income/adapter/entity"


type ListIncomesOutput struct {
	Data []entity.Income
}