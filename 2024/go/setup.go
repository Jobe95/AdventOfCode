package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var folderName string
	var day string

	fmt.Print("Enter the folder name: ")
	fmt.Scanln(&folderName)

	fmt.Print("Enter the day number: ")
	fmt.Scanln(&day)

	err := os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating folder: %v\n", err)
		return
	}

	mainFilePath := filepath.Join(folderName, "main.go")
	mainContent := `package main

import "advent-of-code-2024/utils"

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

func partOne() int {
	output := 0
	return output
}

func partTwo() int {
	output := 0
	return output
}
`
	err = os.WriteFile(mainFilePath, []byte(mainContent), 0644)
	if err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		return
	}

	exampleFilePath := filepath.Join(folderName, "example.txt")
	err = os.WriteFile(exampleFilePath, []byte(""), 0644)
	if err != nil {
		fmt.Printf("Error creating example.txt: %v\n", err)
		return
	}

	inputData := getInput(day)
	if inputData == "" {
		fmt.Println("No input data fetched. Check your session cookie or the day number.")
		return
	}

	inputFilePath := filepath.Join(folderName, "input.txt")
	err = os.WriteFile(inputFilePath, []byte(inputData), 0644)
	if err != nil {
		fmt.Printf("Error creating input.txt: %v\n", err)
		return
	}

	fmt.Printf("Setup complete! Folder '%s' created with main.go, example.txt, and input.txt\n", day)
}

func getInput(day string) string {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		fmt.Println("AOC_SESSION environment variable is not set. Please set it with your session cookie.")
		os.Exit(1)
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error fetching input: %s\nResponse body: %s\n", resp.Status, string(body))
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	return string(body)
}
