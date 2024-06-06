package main

import (
	"strings"
)

func DecodeInstruction(line string) (string, []string) {

	parts := strings.Fields(line)

	instruction := parts[0]
	args := parts[1:]

	// fmt.Println("Instruction and args", instruction, args)
	return instruction, args

}
