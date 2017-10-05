package main

import (
	// Import the entire framework (including bundled verilog)
	_ "sdaccel"

	aximemory "axi/memory"
	axiprotocol "axi/protocol"

	"github.com/joshbohde/add"
)

func Top(
	a uint32,
	b uint32,
	addr uintptr,

	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Since we're not reading anything from memory, disable those reads
	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	// Calculate the value
	val := add.Add(a, b)

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
