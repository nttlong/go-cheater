package libs

import (
	"github.com/mitchellh/go-ps"
	//import sort for sorting processList by Name
	"fmt"
	"path/filepath"
	"sort"

	"github.com/0xrawsec/golang-win32/win32"
	kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
)

type ProcessInfo struct {
	Name string `json:"name"`
	PID  int    `json:"pid"`
	//PID in hex format
	PIDHex string `json:"pid_hex"`
}

func GetProcessList() ([]ProcessInfo, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	var processList []ProcessInfo
	for _, p := range processes {
		processList = append(processList, ProcessInfo{
			Name: p.Executable(),
			PID:  p.Pid(),
			//PID in hex format
			PIDHex: fmt.Sprintf("%x", p.Pid()),
		})
	}
	// sort processList by Name
	sort.Slice(processList, func(i, j int) bool {
		return processList[i].Name < processList[j].Name
	})

	return processList, nil
}

// this function is find pid of process by name in list of ProcessInfo
func FindPidByName(processList []ProcessInfo, name string) int {
	for _, p := range processList {
		if p.Name == name {
			return p.PID
		}
	}
	return -1
}

// this function will find ProcessInfo by Name and return it if not found it will return nil
func FindProcessByName(processList []ProcessInfo, name string) (process ProcessInfo, found bool) {
	for _, p := range processList {
		if p.Name == name {
			process = p // Copy the matching ProcessInfo struct
			found = true
			return // Early return since process is found
		}
	}
	found = false
	return // Return default values if not found
}

// This funcion will get the first address of the process by process id
// Example: GetProcessAddress(1234) will return the first address of the process with pid 1234 looks like 07FFC000
func MemoryReadInit(p *ProcessInfo) (int64, bool) {
	pid := p.PID
	ExeFile := p.Name
	win32handle, _ := kernel32.OpenProcess(0x0010|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, win32.BOOL(0), win32.DWORD(pid))
	moduleHandles, _ := kernel32.EnumProcessModules(win32handle)
	for _, moduleHandle := range moduleHandles {
		s, _ := kernel32.GetModuleFilenameExW(win32handle, moduleHandle)
		targetModuleFilename := ExeFile
		if filepath.Base(s) == targetModuleFilename {
			info, _ := kernel32.GetModuleInformation(win32handle, moduleHandle)
			return int64(info.LpBaseOfDll), true
		}
	}
	return 0, false
}

//This function will get the first address of the process by process name
