package helper

import (
	"fmt"
	"food-ordering/pkg/utils"
)

func GenerateUserID() string {
	timestamp := utils.GenerateTimestamp()
	return fmt.Sprintf("UI%sUI", timestamp)
}

func GenerateMenuID(category string) string {
	timestamp := utils.GenerateTimestamp()
	var categoryCode string

	switch category {
	case "main course":
		categoryCode = "001"
	case "drink":
		categoryCode = "002"
	case "desert":
		categoryCode = "003"
	case "ala carte":
		categoryCode = "004"
	case "steak":
		categoryCode = "005"
	default:
		categoryCode = "000"
	}

	return fmt.Sprintf("MI%s%sMI", timestamp, categoryCode)
}

func GenerateCartID() string {
	timestamp := utils.GenerateTimestamp()
	return fmt.Sprintf("CI%sCI", timestamp)
}

func GenerateOrderID() string {
	timestamp := utils.GenerateTimestamp()
	return fmt.Sprintf("OI%sOI", timestamp)
}

func GeneratePaymentID() string {
	timestamp := utils.GenerateTimestamp()
	randomStr := utils.GenerateRandomString(3)
	return fmt.Sprintf("PIFO%s%s", timestamp, randomStr)
}
