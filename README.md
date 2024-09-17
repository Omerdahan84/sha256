# Go-based Hash Check Project(work in progress)

## Overview
This project is designed to provide tools for verifying and checking file integrity using SHA-256 hashing. The project serves as both a practical utility and an educational exercise in learning about the SHA-256 hashing algorithm. The `sha256.go` script attempts to visualize the process of hashing a file, helping users better understand how the SHA-256 algorithm works.

## Files in the Project:
1. **Hash_test1.txt**: Contains a precomputed SHA-256 hash:  
   `c2a4f4903509957d138e216a6d2c0d7867235c61088c02ca5cf38f2332407b00`

2. **HashcheckFile.txt**: Contains a precomputed SHA-256 hash:  
   `0f46738ebed370c5c52ee0ad96dec8f459fb901c2ca4e285211eddf903bf1598`

3. **Hash_checkFile.txt**: Contains a hash to verify:  
   `0f46738ebed370c5c52ee0ad96dec8f459fb901c2ca4e285211eddf903bf1598`

4. **test1.txt**: Test file for computing hashes, with the content:  
   `"test text"`

5. **sha256.go**: A Go script that computes the SHA-256 hash of a given file. This script not only computes the hash but also helps visualize the process, making it easier to understand how the SHA-256 algorithm works step-by-step.

6. **checkFile.go**: A Go script that compares the hash of a given file with a predefined hash value to validate file integrity.

7. **cryptArr.go**: A Go script related to cryptographic operations, likely used to support hash computation and comparison functionalities.

## Learning Focus: SHA-256 Algorithm
This project is not only a practical utility but also an attempt to deepen understanding of the SHA-256 hashing algorithm. The `sha256.go` script in particular is designed to make the hashing process more transparent, allowing users to see how input data is transformed into a fixed-length hash.

### What is SHA-256?
SHA-256 (Secure Hash Algorithm 256-bit) is part of the SHA-2 family of cryptographic hash functions. It generates a fixed-size 256-bit (32-byte) hash for any input, which is nearly impossible to reverse or duplicate for different inputs. This makes it highly secure for data integrity and cryptographic purposes.

## How to Use

### 1. Compute File Hash:
- To compute the hash of a file, use `sha256.go`. It reads the content of a file and prints the corresponding SHA-256 hash.

### 2. Verify File Hash:
- Use `checkFile.go` to verify the integrity of a file by comparing its hash with a precomputed hash. If the computed hash matches the given one, the file is verified.

### 3. Example:
- File: `test1.txt`
- Computed Hash (using `sha256.go`): Match this with the hash in `Hash_test1.txt` for verification.

## Setup and Installation
1. Install Go from [https://golang.org/](https://golang.org/).
2. Clone this repository or download the required Go files.
3. Compile the Go scripts using the command:
   ```bash
   go build sha256.go checkFile.go cryptArr.go
