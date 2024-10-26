package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modkernel32           = windows.NewLazySystemDLL("kernel32.dll")
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")
)

func ReadProcessMemory(process windows.Handle, baseAddress uintptr, buffer []byte) (int, error) {
	var bytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(process),
		baseAddress,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	return int(bytesRead), nil
}

func getFirstReadableAddr(pid uint32) (uintptr, error) {
	hProcess, err := windows.OpenProcess(windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		return 0, fmt.Errorf("error opening process: %v", err)
	}
	defer windows.CloseHandle(hProcess)

	// For demonstration, we'll use a static memory address
	// NOTE: This address must be valid for the target process
	addr := uintptr(0x0) // Replace with a valid address

	// Buffer to hold data read from memory
	buffer := make([]byte, 4) // Change size as necessary

	// Read memory from the specified address
	bytesRead, err := ReadProcessMemory(hProcess, addr, buffer)
	if err != nil {
		return 0, fmt.Errorf("error reading memory: %v", err)
	}

	// Convert bytes to uint32 (or other type as needed)
	value := *(*uint32)(unsafe.Pointer(&buffer[0]))
	fmt.Printf("Read %d bytes from address 0x%x: %d\n", bytesRead, addr, value)

	return addr, nil
}

func main() {
	pid := uint32(0x8a8) // Replace with the target PID
	addr, err := getFirstReadableAddr(pid)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read from address 0x%x for PID %d\n", addr, pid)
	}
}
