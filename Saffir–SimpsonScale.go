/*The Saffir–Simpson hurricane wind scale (SSHWS),
formerly the Saffir–Simpson hurricane scale (SSHS), classifies hurricanes*/

package main

import (
	"fmt"
)

func main() {
	var windSpeed float32

	for windSpeed >= 0 {
		fmt.Println("Please enter the wind speed of interest in meter/second (m/s): ")
		fmt.Println("Or enter a negative number if you want to stop running the program")
		fmt.Scanln(&windSpeed)

		switch {

		case windSpeed < 0:
			fmt.Println("You have stopped the program")
			fmt.Println("Thank you for using!")
			break

		case windSpeed >= 0 && windSpeed < 33:
			fmt.Printf("The wind speed of %vm/s is not even classified as Catergory one hurricane\n ", windSpeed)

		case windSpeed >= 33 && windSpeed < 43:
			fmt.Printf("The wind speed of %vm/s is classified as Catergory one hurricane\n ", windSpeed)

		case windSpeed >= 43 && windSpeed < 50:
			fmt.Printf("The wind speed of %vm/s is classified as Catergory two hurricane\n ", windSpeed)

		case windSpeed >= 50 && windSpeed < 59:
			fmt.Printf("The wind speed of %v/s is classified as Catergory three hurricane\n ", windSpeed)

		case windSpeed >= 60 && windSpeed < 70:
			fmt.Printf("The wind speed of %v/s is classified as Catergory four hurricane\n ", windSpeed)

		case windSpeed >= 70 && windSpeed <= 103.3:
			fmt.Printf("The wind speed %vm/s is classified as Catergory five hurricane\n ", windSpeed)
		case windSpeed > 103.3:
			exceed := windSpeed - 103.3
			fmt.Printf("You have entered %[2]vm/s which exceeds the highest wind speed ever recorded on earth(103.3m/s) by %[1]v m/s\n ", exceed, windSpeed)
		}
	}
	fmt.Println("Please come again!")
}
