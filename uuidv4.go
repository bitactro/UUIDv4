package uuidv4

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var randomBits [16]byte

func generateUUID4() string {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<16; i++ {
		r1 := rand.Intn(255)
		randomBits[i]=byte(r1) 
	}
	randomBits[6] = (randomBits[6] & 0x0f) | 0x40 // Version 4
	randomBits[8] = (randomBits[8] & 0x3f) | 0x80 // Variant is 10

	var uuid4 string

	part1 := hex.EncodeToString(randomBits[:4])
	uuid4=part1
	uuid4 += "-"
	part2 := hex.EncodeToString(randomBits[4:6])
	uuid4 += part2
	uuid4 += "-"
	part3 := hex.EncodeToString(randomBits[6:8])
	uuid4 += part3
	uuid4 += "-"
	part4 := hex.EncodeToString(randomBits[8:10])
	uuid4 += part4
	uuid4 += "-"
	part5 := hex.EncodeToString(randomBits[10:])
	uuid4 += part5
	
	fmt.Print(uuid4)
	return ""
}