package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var foodResult food
	if order == "burger" {
		foodResult.preptime = 15
	} else if order == "chips" {
		foodResult.preptime = 10
	} else if order == "nuggets" {
		foodResult.preptime = 12
	} else {
		return 404
	}
	return foodResult.preptime
}
