package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Function to split the input array into blocks of a given size
func createBlocks(array []byte, blockSize int) [][]byte {
	var blocks [][]byte
	for i := 0; i < len(array); i += blockSize {
		end := i + blockSize
		if end > len(array) {
			end = len(array)
		}
		blocks = append(blocks, array[i:end])
	}
	return blocks
}

// Function to compute the SHA-256 hash of a block
func hashBlock(block []byte) string {
	hash := sha256.Sum256(block)
	return hex.EncodeToString(hash[:])
}

// Function to verify if a block has been tampered with
func verifyBlock(block []byte, expectedHash string) bool {
	return hashBlock(block) == expectedHash
}

func main() {
	// Sample array to encrypt block by block
	data := []byte("This is a sample array of data to encrypt block by block with SHA-256.")

	// Define block size (e.g., 16 bytes)
	blockSize := 16

	
	blocks := createBlocks(data, blockSize)

	// Encrypt (hash) each block and store the hashes
	var hashedBlocks []string
	for i, block := range blocks {
		hash := hashBlock(block)
		fmt.Printf("Block %d (plaintext): %s\n", i+1, string(block))
		fmt.Printf("Block %d (hashed): %s\n", i+1, hash)
		hashedBlocks = append(hashedBlocks, hash)
	}

	// Verify if the blocks are unchanged (integrity check)
	fmt.Println("\nVerifying blocks...")
	data = []byte("This is a sample array of data to encrypt block by block with SHA-25.")
	blocks = createBlocks(data, blockSize)

	// Create blocks from the data
	for i, block := range blocks {
		if verifyBlock(block, hashedBlocks[i]) {
			fmt.Printf("Block %d is unchanged.\n", i+1)
		} else {
			fmt.Printf("Block %d has been tampered with!\n", i+1)
		}
	}
}
