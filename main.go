package main

import (
	"fmt"
	libs "go-cheater/libs"
	pm "go-cheater/libs/process_manager"
)

func main() {
	pm.NewProcessManager()
	processList, err := libs.GetProcessList()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range processList {
		fmt.Printf("Name: %s, PID: %d\n", p.Name, p.PID)
	}
	pid, found := libs.FindProcessByName(
		processList,
		"eldenring.exe",
	)
	if !found {
		fmt.Println("Error: Process not found")
		return
	}
	//print the pid of eldenring.exe
	fmt.Println("PID of eldenring.exe:", pid.PIDHex)
	address, f := libs.MemoryReadInit(&pid)
	if !f {
		fmt.Println("Error:", f)
		return
	}
	//print the address of the first memory location
	fmt.Println("Address of first memory location:", address)
	//wait for input before exiting
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
