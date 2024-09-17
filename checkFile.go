package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	// Enter the path to the file to be checked
	var input string
	fmt.Print("Enter file name (without extension): \n")
	fmt.Scanln(&input)

	// Use fmt.Sprintf for formatting the filename
	fileName := fmt.Sprintf("%s.txt", input)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Generate the hash file name
	hashFileName := fmt.Sprintf("Hash_%s.txt", input)

	// Check if the hash file exists
	if _, err := os.Stat(hashFileName); os.IsNotExist(err) {
		// If the hash file does not exist, encode and create it
		encode(file, hashFileName)
	} else {
		// Check if the file has changed by comparing the hash
		checkChange(hashFileName, file)
	}
}

// checkChange compares the hash of the file with the existing hash in the hash file
func checkChange(hashFileName string, file *os.File) {
	fmt.Println("Checking if file has changed...")
	hashFile, err := os.Open(hashFileName)
	if err != nil {
		log.Fatalf("Failed to open hash file: %s", err)
	}
	defer hashFile.Close()

	// Read the stored hash from the hash file
	scannerHash := bufio.NewScanner(hashFile)
	scannerFile := bufio.NewScanner(file)

	storedHash := ""
	if scannerHash.Scan() {
		storedHash = scannerHash.Text()
	}

	// Compute the hash of the current file
	currentHash := computeHash(scannerFile)

	// Compare the stored hash with the current hash
	if storedHash != currentHash {
		fmt.Println("File has changed!")
	} else {
		fmt.Println("File has not changed.")
	}
}

// encode creates the hash of the file and stores it in the hash file
func encode(file *os.File, hashFileName string) {
	fmt.Println("Creating new hash file...")

	// Create or overwrite the hash file
	hashFile, err := os.Create(hashFileName)
	if err != nil {
		log.Fatalf("Failed to create hash file: %s", err)
	}
	defer hashFile.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Compute the hash of the file
	hashString := computeHash(scanner)

	// Write the computed hash to the hash file
	_, err = hashFile.WriteString(hashString)
	if err != nil {
		log.Fatalf("Error writing hash to file: %s", err)
	}

	fmt.Println("SHA-256 hash written to file successfully.")
}

// computeHash computes the SHA-256 hash of the file content, line by line
func computeHash(scanner *bufio.Scanner) string {
	// Initialize a new SHA-256 hash
	hash := sha256.New()

	// Read the file line by line and update the hash
	for scanner.Scan() {
		line := scanner.Text() + "\n"
		hash.Write([]byte(line)) // Incremental hash computation per line
	}

	// Compute the SHA-256 checksum
	hashInBytes := hash.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}
