package solutions

import "fmt"

func CountingSundays() {
	count := 0
	firstDayOfTheMonth := 1 // 1 denotes Monday, 7 denotes Sunday

	for year := 1900; year <= 2000; year ++ {
		february := 28

		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			february = 29
		}

		for month := 1; month <= 12; month ++ {
			if firstDayOfTheMonth == 7 && year != 1900 {
				count ++
			}

			var advance int

			switch month {
				case 1, 3, 5, 7, 8, 10, 12:
					advance = 31
				case 4, 6, 9, 11:
					advance = 30
				case 2:
					advance = february
			}

			firstDayOfTheMonth = (firstDayOfTheMonth + advance)%7 + 1
		}
	}
	
	fmt.Println(count)
}
