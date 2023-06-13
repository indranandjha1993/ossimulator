package main

import (
	"fmt"
	"sync"
)

// Process represents a process in the system
type Process struct {
	ID       int
	State    string
	Priority int
	// Other process attributes
}

// ProcessScheduler manages the execution of processes
type ProcessScheduler struct {
	Processes []Process
	// Other scheduler attributes
}

// Mutex for process synchronization
var mutex sync.Mutex

// Function to execute a process
func executeProcess(p *Process) {
	// Logic to execute the process
	fmt.Printf("Executing Process ID: %d\n", p.ID)
}

// Function to synchronize access to shared resources
func accessSharedResource(p *Process) {
	// Acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// Logic to access shared resources
	fmt.Printf("Process ID %d accessing shared resource\n", p.ID)
}

func main() {
	// Initialize the process scheduler
	scheduler := ProcessScheduler{
		Processes: []Process{
			{ID: 1, State: "Ready", Priority: 1},
			{ID: 2, State: "Ready", Priority: 2},
			{ID: 3, State: "Ready", Priority: 3},
		},
	}

	// Execute processes in parallel
	var wg sync.WaitGroup
	for _, p := range scheduler.Processes {
		wg.Add(1)
		go func(process Process) {
			defer wg.Done()
			executeProcess(&process)
		}(p)
	}

	// Wait for all processes to finish execution
	wg.Wait()

	// Access shared resources
	var wg2 sync.WaitGroup
	for _, p := range scheduler.Processes {
		wg2.Add(1)
		go func(process Process) {
			defer wg2.Done()
			accessSharedResource(&process)
		}(p)
	}

	// Wait for all processes to finish accessing shared resources
	wg2.Wait()

	fmt.Println("All processes finished execution")
}
