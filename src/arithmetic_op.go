package main

import (
	"fmt"
	"strconv"
)

func getParity(b byte) byte {
	count := 0
	for b != 0 {
		count += int(b & 1)
		b >>= 1
	}
	if count%2 == 0 {
		return 1 // Even parity
	} else {
		return 0 // Odd parity
	}
}

func getMSB(b byte) byte {
	return (b >> 7) & 0x01
}

func (cpu *CPU) ADD(args []string) bool {
	// fmt.Println("Inside the LDA function")
	// fmt.Println(args)
	sourceRegisterValue, err := getRegisterValue(cpu, args[0])
	if err != nil {
		fmt.Println("	Unable to get value of Register or Memory in ADD function", err)
		return false
	}
	result := uint16(cpu.A) + uint16(sourceRegisterValue)
	cpu.Carry = result > 0xFF
	cpu.AuxCarry = ((cpu.A & 0x0F) + (sourceRegisterValue & 0x0F)) > 0x0F
	cpu.A = byte(result)

	cpu.Zero = (cpu.A == 0)
	cpu.Sign = (getMSB(cpu.A) == 1)
	cpu.Parity = (getParity(cpu.A) == 1)

	return true

}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (cpu *CPU) ADI(args []string) bool {
	// fmt.Println("Inside the ADI function")
	// fmt.Println(args)
	value, err := strconv.ParseUint(args[0], 16, 8)
	if err != nil {
		fmt.Println("	Error in ParseUint in ADI function", err)
		return false
	}
	result := uint16(cpu.A) + uint16(byte(value))
	cpu.Carry = result > 0xFF
	cpu.AuxCarry = ((cpu.A & 0x0F) + (byte(value) & 0x0F)) > 0x0F
	cpu.A = byte(result)

	cpu.Zero = (cpu.A == 0)
	cpu.Sign = (getMSB(cpu.A) == 1)
	cpu.Parity = (getParity(cpu.A) == 1)

	return true

}

func (cpu *CPU) ADC(args []string) bool {
	sourceRegisterValue, err := getRegisterValue(cpu, args[0])
	if err != nil {
		fmt.Println("	Unable to get value of Register or Memory in ADC function", err)
		return false
	}

	// Calculate result including carry
	result := uint16(cpu.A) + uint16(sourceRegisterValue) + uint16(boolToInt(cpu.Carry))
	cpu.Carry = result > 0xFF
	cpu.AuxCarry = ((cpu.A & 0x0F) + (sourceRegisterValue & 0x0F) + uint8(boolToInt(cpu.Carry))) > 0x0F
	cpu.A = byte(result)

	cpu.Zero = (cpu.A == 0)
	cpu.Sign = (getMSB(cpu.A) == 1)
	cpu.Parity = (getParity(cpu.A) == 1)

	return true
}

func (cpu *CPU) ACI(args []string) bool {
	// Parse immediate value from argument
	value, err := strconv.ParseUint(args[0], 16, 8)
	if err != nil {
		fmt.Println("	Error in ParseUint in ACI function", err)
		return false
	}

	// Calculate result including carry and immediate value
	result := uint16(cpu.A) + uint16(value) + uint16(boolToInt(cpu.Carry))
	cpu.Carry = result > 0xFF
	cpu.AuxCarry = ((cpu.A & 0x0F) + uint8(value) + uint8(boolToInt(cpu.Carry))) > 0x0F
	cpu.A = byte(result)

	cpu.Zero = (cpu.A == 0)
	cpu.Sign = (getMSB(cpu.A) == 1)
	cpu.Parity = (getParity(cpu.A) == 1)

	return true
}
