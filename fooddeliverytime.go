package piscine

type food struct {
	nameFood string
	preptime int
}

func FoodDeliveryTime(order string) int {
	answer := 0

	myCatalogFoodTime := []food{
		{nameFood: "burger", preptime: 15},
		{nameFood: "chips", preptime: 10},
		{nameFood: "nuggets", preptime: 12},
	}

	for _, food := range myCatalogFoodTime {
		if order == food.nameFood {
			answer = food.preptime
		}
	}

	if answer == 0 {
		answer = 404
	}

	return answer
}
