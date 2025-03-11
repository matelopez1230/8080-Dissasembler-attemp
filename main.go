package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello there")
	buff, err := os.ReadFile("invaders.h")

	if err != nil {
		log.Fatal("Mate you fucked")
	}

	for pc := 0; pc < len(buff); pc++ {
		op := DecodeOpCode(buff[pc])

		switch op.ByteSize {
		case 1:
			fmt.Printf("0x%04x %s\n", pc, op.Name)
		case 2:
			fmt.Printf("0x%04x %s #$%02x\n", pc, op.Name, buff[pc+1])
			pc += 1
		case 3:
			fmt.Printf("0x%04x %s $%02x%02x\n", pc, op.Name, buff[pc+2], buff[pc+1])
			pc += 2
		default:
			log.Fatalf("Unknown Instruction %d %02x\n", pc, op)
		}
	}
}
