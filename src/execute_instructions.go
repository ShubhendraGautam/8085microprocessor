package main

func ExecuteInstruction(instruction string, args []string, cpu *CPU) bool {
	function_to_execute := instructionMap[instruction]
	function_output := function_to_execute(cpu, args)
	if !function_output {
		return false
	}
	cpu.PC++
	return true
}
