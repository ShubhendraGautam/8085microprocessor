package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	// General Purpose Registers
	A, B, C, D, E, H, L byte

	// Special Purpose Registers
	PC uint16 // Program Counter
	SP uint16 // Stack Pointer

	// Flags
	Zero     bool // Zero flag
	Sign     bool // Sign flag
	Parity   bool // Parity flag
	Carry    bool // Carry flag
	AuxCarry bool // Auxiliary Carry flag

	Memory [65536]byte
}

var instructionMap = map[string]func(*CPU, []string) bool{
	"MOV": func(cpu *CPU, args []string) bool {
		return cpu.MOV(args)
	},
	"LXI": func(cpu *CPU, args []string) bool {
		return cpu.LXI(args)
	},
	"LDA": func(cpu *CPU, args []string) bool {
		return cpu.LDA(args)
	},
	"MVI": func(cpu *CPU, args []string) bool {
		return cpu.MVI(args)
	},
	"STA": func(cpu *CPU, args []string) bool {
		return cpu.STA(args)
	},
	"ADD": func(cpu *CPU, args []string) bool {
		return cpu.ADD(args)
	},
	"ADI": func(cpu *CPU, args []string) bool {
		return cpu.ADI(args)
	},
	"LHLD": func(cpu *CPU, args []string) bool {
		return cpu.LHLD(args)
	},
	"SHLD": func(cpu *CPU, args []string) bool {
		return cpu.SHLD(args)
	},
	"LDAX": func(cpu *CPU, args []string) bool {
		return cpu.LDAX(args)
	},
	"STAX": func(cpu *CPU, args []string) bool {
		return cpu.STAX(args)
	},
	"XCHG": func(cpu *CPU, args []string) bool {
		return cpu.XCHG(args)
	},
	"ADC": func(cpu *CPU, args []string) bool {
		return cpu.ADC(args)
	},
	"ACI": func(cpu *CPU, args []string) bool {
		return cpu.ACI(args)
	},
}

func InitializeCPU() *CPU {
	cpu := &CPU{
		A:        0, // Initialize to zero
		B:        0,
		C:        0,
		D:        0,
		E:        0,
		H:        0,
		L:        0,
		PC:       0, // Start at the beginning of memory
		SP:       0xFFFF,
		Zero:     false,
		Sign:     false,
		Parity:   false,
		Carry:    false,
		AuxCarry: false,
	}
	return cpu
}

func getRegisterValue(cpu *CPU, reg string) (byte, error) {
	switch reg {

	case "A":
		return cpu.A, nil
	case "B":
		return cpu.B, nil
	case "C":
		return cpu.C, nil
	case "D":
		return cpu.D, nil
	case "E":
		return cpu.E, nil
	case "H":
		return cpu.H, nil
	case "L":
		return cpu.L, nil
	case "M":
		memory_address := uint16(cpu.H)<<8 | uint16(cpu.L)
		return cpu.Memory[memory_address], nil
	default:
		return 0, errors.New("invalid register")
	}

}

func setRegisterValue(cpu *CPU, reg string, value byte) {

	switch reg {

	case "A":
		cpu.A = value
		// fmt.Println("I was here")
	case "B":
		cpu.B = value
	case "C":
		cpu.C = value
	case "D":
		cpu.D = value
	case "E":
		cpu.E = value
	case "H":
		cpu.H = value
	case "L":
		cpu.L = value
	default:
		fmt.Println("Unknown Type of Register")
	}

}
