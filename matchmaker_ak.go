package main

import (
	"fmt"
	"math"
	"os"
)

// constants

const DESIREONE = 4
const DESIRETWO = 1
const DESIRETHREE = 5
const DESIREFOUR = 2
const DESIREFIVE = 3
const WEIGHTONE = 1
const WEIGHTTWO = 2
const WEIGHTTHREE = 3
const COMPATBEST = 89
const COMPATMIDDLE = 75
const COMPATWORST = 70

func validate(x int) {
	if (x != 1) && (x != 2) && (x != 3) && (x != 4) && (x != 5) {
		fmt.Println("Error. Type in only 1,2,3,4,5 please!")
		os.Exit(0)
	}
}
func main() {
	fmt.Println("Welcome to my Matchmaker Program.")
	fmt.Println(" ")
	fmt.Println("Let's see if we are soulmates!")
	fmt.Println("Read the question, then enter a number from one to five.")
	fmt.Println("one being completely disagree and five being completely agree.")
	fmt.Println("")

	fmt.Println("My favorite color is Pink")
	var first int
	fmt.Scanln(&first)
	validate(first)
	proxione := DESIREONE - first
	proxone := int(math.Abs(float64(proxione)))
	sum1 := proxone * WEIGHTONE
	fmt.Println(" ")
	fmt.Println("Compatibility Score: ", proxone)
	fmt.Println("Weighted Score : ", sum1)
	fmt.Println(" ")

	fmt.Println("I hate Water")
	var second int
	fmt.Scanln(&second)
	validate(second)
	proxitwo := DESIRETWO - second
	proxtwo := int(math.Abs(float64(proxitwo)))
	sum2 := proxtwo * WEIGHTTHREE
	fmt.Println(" ")
	fmt.Println("Compatibility Score: ", proxtwo)
	fmt.Println("Weighted Score : ", sum2)
	fmt.Println(" ")

	fmt.Println("I think the two-party political system is bad")
	var third int
	fmt.Scanln(&third)
	validate(third)
	proxithree := DESIRETHREE - third
	proxthree := int(math.Abs(float64(proxithree)))
	sum3 := proxthree * WEIGHTTHREE
	fmt.Println(" ")
	fmt.Println("Compatibility Score: ", proxthree)
	fmt.Println("Weighted Score : ", sum3)
	fmt.Println(" ")

	fmt.Println("I love math!")
	var fourth int
	fmt.Scanln(&fourth)
	validate(fourth)
	proxifour := DESIREFOUR - fourth
	proxfour := int(math.Abs(float64(proxifour)))
	sum4 := proxfour * WEIGHTTWO
	fmt.Println(" ")
	fmt.Println("Compatibility Score: ", proxfour)
	fmt.Println("Weighted Score : ", sum4)
	fmt.Println(" ")

	fmt.Println("Programming is fun.")
	var fifth int
	fmt.Scanln(&fifth)
	validate(fifth)
	proxifive := DESIREFIVE - fifth
	proxfive := int(math.Abs(float64(proxifive)))
	sum5 := proxfive * WEIGHTONE
	fmt.Println(" ")
	fmt.Println("Compatibility Score: ", proxfive)
	fmt.Println("Weighted Score : ", sum5)
	fmt.Println(" ")

	compatibility := sum1 + sum2 + sum3 + sum4 + sum5
	finalcompatscore := 100 - compatibility
	fmt.Printf("Final Compatibility Score: %d %% ", finalcompatscore)

	if finalcompatscore >= COMPATBEST {
		fmt.Println("")
		fmt.Println("Wow, we really are soulmates.")
	} else if finalcompatscore >= COMPATMIDDLE && finalcompatscore < COMPATBEST {
		fmt.Println("")
		fmt.Println("Eh, we'd get along somewhat")
	} else if finalcompatscore < COMPATMIDDLE && finalcompatscore <= COMPATWORST {
		fmt.Println("")
		fmt.Println("Ugh, you can't sit with us.")
	}

}
