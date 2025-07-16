package main

import "fmt"

func main() {
	fmt.Println("********************************")
	fmt.Println("Welcome to myy quiz game")
	fmt.Println("********************************")

	fmt.Printf("Enter your name: \n")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, YOU ARE IN THE GAME NOWW!!, BEST OF LUCK\n", name)
	fmt.Println("********************************")
	fmt.Print("\n")
	fmt.Println("Enter your age :")
	var age int
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You are eligible to play the game")
	} else {
		fmt.Println("Sorry, you are not eligible to play the game")
		return
	}

	var score int = 0
	var question int = 0
	var answer string
	var answer2 string

	fmt.Print("\n\n")
	fmt.Println("Let's start the quiz")
	fmt.Print("********************************")

	// Q1
	fmt.Println("First Question is: ")
	fmt.Printf("What is better, the RTX 3080 or the RTX 3090?\n")
	question += 1
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" || answer+" "+answer2 == "rtX 3090" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q2
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Let's move to the next question")
	fmt.Println("Second Question is: ")
	fmt.Printf("How many cores does the Ryzen 9 have?\n")
	question += 1
	var cores uint
	fmt.Scan(&cores)
	if cores == 12 {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q3
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Third Question is: ")
	fmt.Printf("What does CPU stand for?\n")
	question += 1
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "Central Processing" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q4
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Fourth Question is: ")
	fmt.Printf("What is the latest version of Windows as of 2025?\n")
	question += 1
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "Windows 11" || answer+" "+answer2 == "windows 11" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q5
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Fifth Question is: ")
	fmt.Printf("Which company manufactures the iPhone?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "Apple" || answer == "apple" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q6
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Sixth Question is: ")
	fmt.Printf("How many bits are in a byte?\n")
	question += 1
	var bits int
	fmt.Scan(&bits)
	if bits == 8 {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q7
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Seventh Question is: ")
	fmt.Printf("Which is faster: SSD or HDD?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "SSD" || answer == "ssd" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q8
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Eighth Question is: ")
	fmt.Printf("How many GB is equal to 1 TB?\n")
	question += 1
	var gb int
	fmt.Scan(&gb)
	if gb == 1024 {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q9
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Ninth Question is: ")
	fmt.Printf("What language is primarily used for Android development?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "Kotlin" || answer == "kotlin" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q10
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Tenth Question is: ")
	fmt.Printf("What does GPU stand for?\n")
	question += 1
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "Graphics Processing" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q11
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Eleventh Question is: ")
	fmt.Printf("What is the capital of India?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "Delhi" || answer == "delhi" || answer == "NewDelhi" || answer == "newdelhi" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q12
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Twelfth Question is: ")
	fmt.Printf("Which company makes GeForce GPUs?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "Nvidia" || answer == "nvidia" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q13
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Thirteenth Question is: ")
	fmt.Printf("What year was Go language released? (Only year)\n")
	question += 1
	var year int
	fmt.Scan(&year)
	if year == 2009 {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q14
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Fourteenth Question is: ")
	fmt.Printf("What does RAM stand for?\n")
	question += 1
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "Random Access" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Q15
	fmt.Print("\n\n")
	fmt.Println("********************************")
	fmt.Println("Fifteenth Question is: ")
	fmt.Printf("Which language is used to build iOS apps?\n")
	question += 1
	fmt.Scan(&answer)
	if answer == "Swift" || answer == "swift" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Incorrect!!")
	}

	// Result
	fmt.Print("\n\n")
	fmt.Printf("You scored points %v out of %v, Well Done:\n)", score, question)
	percent := (float64(score) / float64(question)) * 100
	fmt.Printf("You scored percentage of %v%%\n", percent)
}
