package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var program_memory []string

func printError(err error) {
	fmt.Println("Error Encountered")
	fmt.Println("The error is: ", err)
}

func printBoxedMessage(message string) {
	messageLength := len(message)
	border := strings.Repeat("-", messageLength+4)
	fmt.Println(border)
	fmt.Printf("| %s |\n", message)
	fmt.Println(border)
}

func printCpuState(c *CPU) {
	header := "| A  | B  | C  | D  | E  | H  | L  |  PC  |  SP  | Sign | Zero | Parity | Carry |  AuxCarry  |"
	border := strings.Repeat("-", len(header))
	fmt.Println(border)
	fmt.Println(header)
	fmt.Println(border)
	fmt.Printf("| %02X | %02X | %02X | %02X | %02X | %02X | %02X | %04X | %04X | %5v| %5v| %6v | %5v |    %5v   |\n",
		c.A, c.B, c.C, c.D, c.E, c.H, c.L, c.PC, c.SP, c.Sign, c.Zero, c.Parity, c.Carry, c.AuxCarry)
	fmt.Println(border)
}

/*

	Function To Read the filename
	Use to get interactive filename

*/
// func readFileName() string {
// 	// User input for the test file
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Printf("Input the file name: ")
// 	fileName, err := reader.ReadString('\n')
// 	if err != nil {
// 		return "Error"
// 	}
// 	fileName = strings.TrimSpace(fileName) // Trim newline and any surrounding spaces
// 	fmt.Println("Filename is : ", fileName)
// 	return fileName

// }

func executeFile(fileName string) {
	// Reading the file line by line
	file, err := os.Open(fileName)
	if err != nil {
		printError(err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function returns

	fmt.Printf("\nParsing the file...\n\n")

	scanner := bufio.NewScanner(file)

	// Initialzing the CPU
	newCPU := InitializeCPU()
	newCPU.Memory[20480] = 20

	// Wrting the file in memory

	for scanner.Scan() {
		line := scanner.Text()
		program_memory = append(program_memory, line)
	}
	// fmt.Println("Data at address 20480 is ", newCPU.Memory[20480])
	// fmt.Println("Data at address 8256 is ", newCPU.Memory[8256])
	for newCPU.PC < uint16(len(program_memory)) {
		line := program_memory[newCPU.PC]
		fmt.Println("	CURRENT CPU STATE")
		printCpuState(newCPU)
		fmt.Printf("\n\nInstruction number: %v, : Recieved: %v\n\n", newCPU.PC, line)
		instruction, args := DecodeInstruction(line)
		execution_output := ExecuteInstruction(instruction, args, newCPU)
		if !execution_output {
			os.Exit(1)
		}
	}
	fmt.Println("	FINAL CPU STATE")
	printCpuState(newCPU)

	// fmt.Println("Data at address 20480 is ", newCPU.Memory[20480])
	// fmt.Println("Data at address 8256 is ", newCPU.Memory[8256])

	if err := scanner.Err(); err != nil {
		printError(err)
		return
	}

}

func main() {

	// Welcome message
	printBoxedMessage("Simple 8085 Microprocessor Demo in Golang")

	// fileName := readFileName()
	// if fileName == "Error" {
	// 	fmt.Println("Invalid file name.")
	// }
	executeFile("test.asm")

}
