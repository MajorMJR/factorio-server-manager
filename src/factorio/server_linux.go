// use this file only when compiling not windows (all unix systems)
// +build !windows

package factorio

import (
	"log"
	"os"
)

// Stubs for windows-only functions

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
	log.Printf("Sent SIGKILL to Factorio process. Factorio forced to exit.")

	err = server.Rcon.Close()
	if err != nil {
		log.Printf("Error close rcon connection: %s", err)
	}

	return nil
}

func (server *Server) Stop() error {
	err := server.Cmd.Process.Signal(os.Interrupt)
	if err != nil {
		if err.Error() == "os: process already finished" {
			server.SetRunning(false)
			return err
		}
		log.Printf("Error sending SIGINT to Factorio process: %s", err)
		return err
	}
	log.Printf("Sent SIGINT to Factorio process. Factorio shutting down...")

	err = server.Rcon.Close()
	if err != nil {
		log.Printf("Error close rcon connection: %s", err)
	}

	return nil
}

func (server *Server) checkProcessHealth(text string) {
	// ignore
}
