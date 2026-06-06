package piscine

func ReverseMenuIndex(menu []string) []string {
	for i := 0; i < len(menu)-1; i++ {
		if i >= len(menu)/2 {
			menu[i], menu[len(menu)-1-i] = menu[len(menu)-1-i], menu[i]
		}
	}
	return menu
}
