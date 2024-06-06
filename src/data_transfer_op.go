package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (cpu *CPU) LDA(args []string) bool {
	// fmt.Println("Inside the LDA function")
	memory_address, err := strconv.ParseUint(args[0], 16, 16)
	if err != nil {
		fmt.Println("	Error in ParseUint in STA function", err)
		return false
	}
	fmt.Printf("	Value at the memory address %X is %X\n ", memory_address, cpu.Memory[memory_address])
	cpu.A = cpu.Memory[memory_address]
	return true

}

func (cpu *CPU) LXI(args []string) bool {
	// fmt.Println("Inside the LXI function")
	split_args := strings.Split(args[0], ",")
	split_args_length := len(split_args)

	if split_args_length != 2 {
		fmt.Println("	Invalid Number of Arguments in LXI")
		return false
	} else {
		hexString := split_args[1]

		// Convert the hexadecimal string to uint16
		value, err := strconv.ParseUint(hexString, 16, 16)
		if err != nil {
			fmt.Println("	Error in ParseUint in LXI function", err)
			return false
		}

		memory_address := uint16(value)

		if split_args[0] == "B" {
			cpu.B = byte(memory_address >> 8)   // Store the first 8 bits in register B. We use right shift the address by 8 to get the first half
			cpu.C = byte(memory_address & 0xFF) // This will take only the last 8 bits
		} else if split_args[0] == "D" {
			cpu.D = byte(memory_address >> 8)
			cpu.E = byte(memory_address & 0xFF)
		} else if split_args[0] == "H" {
			cpu.H = byte(memory_address >> 8)
			cpu.L = byte(memory_address & 0xFF)
		} else {
			fmt.Println("	Invalid Register Pair recieved in LXI")
			return false
		}

	}
	return true
}
func (cpu *CPU) STA(args []string) bool {
	// fmt.Println("Inside the STA function")
	memory_address, err := strconv.ParseUint(args[0], 16, 16)
	if err != nil {
		fmt.Println("	Error in ParseUint in STA function", err)
		return false
	}
	cpu.Memory[memory_address] = cpu.A
	fmt.Printf("	Value at the memory address %X is %X\n ", memory_address, cpu.Memory[memory_address])
	return true

}

func (cpu *CPU) MOV(args []string) bool {
	// fmt.Println("Inside the MOV function")
	split_args := strings.Split(args[0], ",")
	split_args_length := len(split_args)

	if split_args_length != 2 {
		fmt.Println("	Invalid Number of Arguments in MOV")
		return false
	} else {
		dest, src := split_args[0], split_args[1]
		source_register_value, err := getRegisterValue(cpu, src)
		if err != nil {
			printError(err)
			return false
		}

		if dest == "M" {
			memory_address := uint16(cpu.H)<<8 | uint16(cpu.L)
			cpu.Memory[memory_address] = source_register_value
		} else {
			setRegisterValue(cpu, dest, source_register_value)
		}

	}
	return true
}

func (cpu *CPU) MVI(args []string) bool {

	split_args := strings.Split(args[0], ",")
	split_args_length := len(split_args)

	if split_args_length != 2 {
		fmt.Println("	Invalid Number of Arguments in MVI")
		return false
	} else {
		dest, value := split_args[0], split_args[1]
		source_value, err := strconv.ParseUint(value, 16, 8) // Parse hexadecimal string to uint8
		if err != nil {
			fmt.Println("	Error in ParseUint in MVI function", err)
			return false
		}
		if dest == "M" {
			memory_address := uint16(cpu.H)<<8 | uint16(cpu.L)
			cpu.Memory[memory_address] = byte(source_value)
			fmt.Printf("	Value at memory address %X is %X\n", memory_address, byte(source_value))
		} else {
			setRegisterValue(cpu, dest, byte(source_value))
		}
	}
	return true

}

func (cpu *CPU) LHLD(args []string) bool {
	memory_address_for_H, err := strconv.ParseUint(args[0], 16, 16)

	if err != nil {
		fmt.Println("	Error in ParseUint in LHLD function", err)
		return false
	}
	memory_address_for_L := memory_address_for_H + 1

	cpu.H = cpu.Memory[memory_address_for_H]
	cpu.L = cpu.Memory[memory_address_for_L]

	// fmt.Println(memory_address_for_H, memory_address_for_L)

	return true
}

func (cpu *CPU) SHLD(args []string) bool {
	memory_address_for_H, err := strconv.ParseUint(args[0], 16, 16)

	if err != nil {
		fmt.Println("	Error in ParseUint in SHLD function", err)
		return false
	}
	memory_address_for_L := memory_address_for_H + 1

	cpu.Memory[memory_address_for_H] = cpu.H
	cpu.Memory[memory_address_for_L] = cpu.L

	fmt.Printf("	Value at memory address %X is %X\n", memory_address_for_H, cpu.Memory[memory_address_for_H])
	fmt.Printf("	Value at memory address %X is %X\n", memory_address_for_L, cpu.Memory[memory_address_for_L])

	// fmt.Println(memory_address_for_H, memory_address_for_L)

	return true
}

func (cpu *CPU) LDAX(args []string) bool {

	if len(args) == 0 {
		fmt.Println("No arguments provided in LDAX")
		return false
	}

	if args[0] == "B" {
		memory_address := uint16(cpu.B)<<8 | uint16(cpu.C)
		cpu.A = cpu.Memory[memory_address]
	} else if args[0] == "D" {
		memory_address := uint16(cpu.D)<<8 | uint16(cpu.E)
		cpu.A = cpu.Memory[memory_address]
	} else {
		fmt.Println("Invalid argument in LDAX")
		return false
	}

	return true
}

func (cpu *CPU) STAX(args []string) bool {

	if len(args) == 0 {
		fmt.Println("	No arguments provided in STAX")
		return false
	}

	if args[0] == "B" {
		memory_address := uint16(cpu.B)<<8 | uint16(cpu.C)
		cpu.Memory[memory_address] = cpu.A
	} else if args[0] == "D" {
		memory_address := uint16(cpu.D)<<8 | uint16(cpu.E)
		cpu.Memory[memory_address] = cpu.A
	} else {
		fmt.Println("	Invalid argument in STAX")
		return false
	}

	return true
}

func (cpu *CPU) XCHG(args []string) bool {

	tempH := cpu.H
	tempL := cpu.L

	cpu.H = cpu.D
	cpu.L = cpu.E

	cpu.D = tempH
	cpu.E = tempL

	return true
}
