package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	BlockSize     = 512
	SeparatorLen  = 1 // for the "1" separator
	SizeReserve   = 64 // reserve for length encoding in binary
	wordInSchdule = 64
)

// Padding function
func padding(message *string) {
	// Determine the length of the message in bits
	messageLen := len(*message)    // message length in bits
	messageLenWithInfo := messageLen + SeparatorLen + SizeReserve
	
	// Calculate the required block size in bits (multiple of 512)
	closestBlock := int(math.Ceil(float64(messageLenWithInfo) / BlockSize))
	closestFullLen := closestBlock * BlockSize

	// Number of zeros needed for padding
	zeroNeeded := closestFullLen - messageLenWithInfo
	
	// Convert the original message length to binary (64-bit)
	binaryRep := strconv.FormatUint(uint64(messageLen), 2)
	
	// Pad with '1', followed by required zeros, then append the binary length
	*message += "1" + strings.Repeat("0", zeroNeeded) + fmt.Sprintf("%064s", binaryRep)
}

// Convert message to binary string
func convertMessageTobin(message []byte) string {
	res := ""
	for _, ch := range message {
		// Get the binary representation as a string and pad to 8 bits
		binaryRep := strconv.FormatUint(uint64(ch), 2)
		res += fmt.Sprintf("%08s", binaryRep)
	}
	return res
}
func createSechdule(message string,start int)[]string{
	schedule := make([]string, wordInSchdule)
	for i := 0; i < 512; i += 32 {
		schedule[i/32] = message[start + i: start + i+32]
	}
	return schedule
}
func main() {
	var input string
	fmt.Print("Enter a string: \n")
	fmt.Scanln(&input)
	// Convert the input string to binary
	message := convertMessageTobin([]byte(input))
	
	fmt.Printf("============= Start sha256 =============\n")
	fmt.Printf("message\n----------\n intput: %s \n binary intput: %s \n",input,message)
	// Apply padding
	padding(&message)
	// Print out the binary message in blocks of 512 bits
	// Correct slicing without going out of bounds
	printPadded(message)
	fmt.Printf("============= create schedule =============\n")

	for i := 0; i < len(message); i+= 512{
		schedule := createSechdule(message,i)
		for j := 0 ; j < 16; j++{
			if j <= 9 {
				fmt.Printf( "Word  %d:%s \n", j,schedule[j])

			} else {
				fmt.Printf( "Word %d:%s \n", j,schedule[j])

			}
		}
		// for i := 16; i <  wordInSchdule;i++{
		// 	s0 := (rightRotate(schedule[i-15], 7)) ^ (rightRotate(schedule[i-15], 18)) ^ (rightShift(schedule[i-15], 7))
			
        // 	s1 := (rightRotate(schedule[i-2] , 17)) ^ (rightRotate(schedule[i-2] , 19)) ^ (rightShift(schedule[i-2],10))
			 
		// }
        
	}
	
}

func printPadded(res string) {
	fmt.Printf("============== padding ==============\n")
	for i := 0; i < len(res)/512; i++ {
		fmt.Printf("Block %d\n", i+1)
		fmt.Println(res[i*512 : (i+1)*512])
	}
}

//Initialize hash values:
//(first 32 bits of the fractional parts of the square roots of the first primes 2..19):

func rightRotate(word string, shift int) uint64 {
    wordLen := len(word)
    number, _ := strconv.ParseUint(word, 2, 64)  // Use ParseUint for unsigned numbers
    shift %= wordLen  // To handle shifts greater than the word length
    rotated := (number >> shift) | (number << (wordLen - shift)) & ((1 << wordLen) - 1)  // Rotate and mask
    return rotated
}
func rightShift(word string,shift int) (uint64 ){
	number, _ := strconv.ParseInt(word, 2, 64)
	
	number >>= uint64(shift)
	return uint64(number)

}

// h0 := 0x6a09e667
// h1 := 0xbb67ae85
// h2 := 0x3c6ef372
// h3 := 0xa54ff53a
// h4 := 0x510e527f
// h5 := 0x9b05688c
// h6 := 0x1f83d9ab
// h7 := 0x5be0cd19
// //Initialize array of round constants:
// //(first 32 bits of the fractional parts of the cube roots of the first 64 primes 2..311):


// primeSlice :=[]int{

//    0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
//    0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
//    0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
//    0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
//    0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
//    0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
//    0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
//    0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}
