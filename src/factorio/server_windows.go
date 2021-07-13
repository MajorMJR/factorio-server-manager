package factorio

import (
	"log"
	"os"
	"strings"
	"syscall"
	"time"
)

func sendCtrlCToPid(pid int) {
	d, e := syscall.LoadDLL("kernel32.dll")
	if e != nil {
		log.Fatalf("LoadDLL: %v\n", e)
	}
	p, e := d.FindProc("GenerateConsoleCtrlEvent")
	if e != nil {
		log.Fatalf("FindProc: %v\n", e)
	}
	r, _, e := p.Call(uintptr(syscall.CTRL_C_EVENT), uintptr(pid))
	if r == 0 {
		log.Fatalf("GenerateConsoleCtrlEvent: %v\n", e)
	}
}

func setCtrlHandlingIsDisabledForThisProcess(disabled bool) {
	disabledInt := 0
	if disabled {
		disabledInt = 1
	}

	d, e := syscall.LoadDLL("kernel32.dll")
	if e != nil {
		log.Fatalf("LoadDLL: %v\n", e)
	}
	p, e := d.FindProc("SetConsoleCtrlHandler")
	if e != nil {
		log.Fatalf("FindProc: %v\n", e)
	}
	r, _, e := p.Call(uintptr(0), uintptr(disabledInt))
	if r == 0 {
		log.Fatalf("SetConsoleCtrlHandler: %v\n", e)
	}
}

func (server *Server) Kill() error {
	err := server.Cmd.Process.Signal(os.Kill)
	if err != nil {
		if err.Error() == "os: process already finished" {
			server.SetRunning(false)
			return err
		}
		log.Printf("Error sending SIGKILL to Factorio process: %s", err)
		return err
	}
	server.SetRunning(false)
	log.Println("Sent SIGKILL to Factorio process. Factorio forced to exit.")

	return nil
}

func (server *Server) Stop() error {
	// Disable our own handling of CTRL+C, so we don't close when we send it to the console.
	setCtrlHandlingIsDisabledForThisProcess(true)

	// Send CTRL+C to all processes attached to the console (ourself, and the factorio server instance)
	sendCtrlCToPid(0)
	log.Println("Sent SIGINT to Factorio process. Factorio shutting down...")
	time.Sleep(20 * time.Millisecond)
	// Re-enable handling of CTRL+C after we're sure that the factorio server is shut down.
	setCtrlHandlingIsDisabledForThisProcess(false)

	return nil
}

func (server *Server) checkProcessHealth(text string) {
	// check if the output indicates a server shutdown
	if strings.Contains(text, "ServerMultiplayerManager.cpp:783: updateTick(0) changing state from(Disconnected) to(Closed)") {
		// Somehow, the Factorio devs managed to code the game to react appropriately to CTRL+C, including
		// saving the game, but not actually exit. So, we still have to manually kill the process, and
		// for extra fun, there's no way to know when the server save has actually completed (unless we want
		// to inject filesystem logic into what should be a process-level Stop() routine), so our best option
		// is to just wait an arbitrary amount of time and hope that the save is successful in that time.
		server.Cmd.Process.Signal(os.Kill)
	}
}
