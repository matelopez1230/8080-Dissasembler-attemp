package main

type CpuFlags struct {
	Parity   uint8
	Zero     uint8
	Sing     uint8
	Carry    uint8
	AuxCarry uint8
}

type CpuState struct {
	RegA      uint8
	RegB      uint8
	RegC      uint8
	RegD      uint8
	RegE      uint8
	RegH      uint8
	RegL      uint8
	SP        uint16
	PC        uint16
	Memory    *uint8
	IntEnable bool
	Condition CpuFlags
}
