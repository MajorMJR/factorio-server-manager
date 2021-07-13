package factorio

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/OpenFactorioServerManager/factorio-server-manager/api/websocket"
	"github.com/OpenFactorioServerManager/factorio-server-manager/bootstrap"
	"github.com/OpenFactorioServerManager/rcon"
)

type Server struct {
	Cmd            *exec.Cmd              `json:"-"`
	Savefile       string                 `json:"savefile"`
	Latency        int                    `json:"latency"`
	BindIP         string                 `json:"bindip"`
	Port           int                    `json:"port"`
	Running        bool                   `json:"running"`
	Version        Version                `json:"fac_version"`
	BaseModVersion string                 `json:"base_mod_version"`
	StdOut         io.ReadCloser          `json:"-"`
	StdErr         io.ReadCloser          `json:"-"`
	StdIn          io.WriteCloser         `json:"-"`
	Settings       map[string]interface{} `json:"-"`
	Rcon           *rcon.RemoteConsole    `json:"-"`
	LogChan        chan []string          `json:"-"`
}

var instantiated Server
var once sync.Once

func (server *Server) SetRunning(newState bool) {
	if server.Running != newState {
		log.Println("new state, will also send to correct room")
		server.Running = newState
		wsRoom := websocket.WebsocketHub.GetRoom("server_status")
		response, _ := json.Marshal(server)
		wsRoom.Send(string(response))
	}
}

func (server *Server) GetRunning() bool {
	return server.Running
}

func (server *Server) autostart() {
	var err error
	if server.BindIP == "" {
		server.BindIP = "0.0.0.0"

	}
	if server.Port == 0 {
		server.Port = 34197
	}
	server.Savefile = "Load Latest"

	err = server.Run()

	if err != nil {
		log.Printf("Error starting Factorio server: %+v", err)
		return
	}

}

func SetFactorioServer(server Server) {
	instantiated = server
}

func NewFactorioServer() (err error) {
	server := Server{}
	server.Settings = make(map[string]interface{})
	config := bootstrap.GetConfig()
	if err = os.MkdirAll(config.FactorioConfigDir, 0755); err != nil {
		log.Printf("failed to create config directory: %v", err)
		return
	}

	settingsPath := config.SettingsFile
	var settings *os.File

	if _, err = os.Stat(settingsPath); os.IsNotExist(err) {
		// copy example settings to supplied settings file, if not exists
		log.Printf("Server settings at %s not found, copying example server settings.\n", settingsPath)

		examplePath := filepath.Join(config.FactorioDir, "data", "server-settings.example.json")

		var example *os.File
		example, err = os.Open(examplePath)
		if err != nil {
			log.Printf("failed to open example server settings: %v", err)
			return
		}
		defer example.Close()

		settings, err = os.Create(settingsPath)
		if err != nil {
			log.Printf("failed to create server settings file: %v", err)
			return
		}
		defer settings.Close()

		_, err = io.Copy(settings, example)
		if err != nil {
			log.Printf("failed to copy example server settings: %v", err)
			return
		}

		err = example.Close()
		if err != nil {
			log.Printf("failed to close example server settings: %s", err)
			return
		}
	} else {
		// otherwise, open file normally
		settings, err = os.Open(settingsPath)
		if err != nil {
			log.Printf("failed to open server settings file: %v", err)
			return
		}
		defer settings.Close()
	}

	// before reading reset offset
	if _, err = settings.Seek(0, 0); err != nil {
		log.Printf("error while seeking in settings file: %v", err)
		return
	}

	if err = json.NewDecoder(settings).Decode(&server.Settings); err != nil {
		log.Printf("error reading %s: %v", settingsPath, err)
		return
	}

	log.Printf("Loaded Factorio settings from %s\n", settingsPath)

	out := []byte{}
	//Load factorio version
	if config.GlibcCustom == "true" {
		out, err = exec.Command(config.GlibcLocation, "--library-path", config.GlibcLibLoc, config.FactorioBinary, "--version").Output()
	} else {
		out, err = exec.Command(config.FactorioBinary, "--version").Output()
	}

	if err != nil {
		log.Printf("error on loading factorio version: %s", err)
		return
	}

	reg := regexp.MustCompile("Version.*?((\\d+\\.)?(\\d+\\.)?(\\*|\\d+)+)")
	found := reg.FindStringSubmatch(string(out))
	err = server.Version.UnmarshalText([]byte(found[1]))
	if err != nil {
		log.Printf("could not parse version: %v", err)
		return
	}

	//Load baseMod version
	baseModInfoFile := filepath.Join(config.FactorioBaseModDir, "info.json")
	bmifBa, err := ioutil.ReadFile(baseModInfoFile)
	if err != nil {
		log.Printf("couldn't open baseMods info.json: %s", err)
		return
	}
	var modInfo ModInfo
	err = json.Unmarshal(bmifBa, &modInfo)
	if err != nil {
		log.Printf("error unmarshalling baseMods info.json to a modInfo: %s", err)
		return
	}

	server.BaseModVersion = modInfo.Version

	// load admins from additional file
	if (server.Version.Greater(Version{0, 17, 0})) {
		if _, err = os.Stat(config.FactorioAdminFile); os.IsNotExist(err) {
			//save empty admins-file
			err = ioutil.WriteFile(config.FactorioAdminFile, []byte("[]"), 0664)
			server.Settings["admins"] = make([]string, 0)
		} else {
			var data []byte
			data, err = ioutil.ReadFile(config.FactorioAdminFile)
			if err != nil {
				log.Printf("Error loading FactorioAdminFile: %s", err)
				return
			}

			var jsonData interface{}
			err = json.Unmarshal(data, &jsonData)
			if err != nil {
				log.Printf("Error unmarshalling FactorioAdminFile: %s", err)
				return
			}

			server.Settings["admins"] = jsonData
		}
	}

	SetFactorioServer(server)

	// autostart factorio is configured to do so
	if config.Autostart == "true" {
		go instantiated.autostart()
	}

	return
}

func GetFactorioServer() (f *Server) {
	return &instantiated
}

func (server *Server) Run() error {
	var err error
	config := bootstrap.GetConfig()
	data, err := json.MarshalIndent(server.Settings, "", "  ")
	if err != nil {
		log.Println("Failed to marshal FactorioServerSettings: ", err)
	} else {
		ioutil.WriteFile(config.SettingsFile, data, 0644)
	}

	saves, err := ListSaves(config.FactorioSavesDir)
	if err != nil {
		log.Println("Failed to get saves list: ", err)
	}

	if len(saves) == 0 {
		return errors.New("No savefile exists on the server")
	}

	args := []string{}

	//The factorio server refenences its executable-path, since we execute the ld.so file and pass the factorio binary as a parameter
	//the game would use the path to the ld.so file as it's executable path and crash, to prevent this the parameter "--executable-path" is added
	if config.GlibcCustom == "true" {
		log.Println("Custom glibc selected, glibc.so location:", config.GlibcLocation, " lib location:", config.GlibcLibLoc)
		args = append(args, "--library-path", config.GlibcLibLoc, config.FactorioBinary, "--executable-path", config.FactorioBinary)
	}

	args = append(args,
		"--bind", server.BindIP,
		"--port", strconv.Itoa(server.Port),
		"--server-settings", config.SettingsFile,
		"--rcon-port", strconv.Itoa(config.FactorioRconPort),
		"--rcon-password", config.FactorioRconPass)

	if (server.Version.Greater(Version{0, 17, 0})) {
		args = append(args, "--server-adminlist", config.FactorioAdminFile)
	}

	if server.Savefile == "Load Latest" {
		args = append(args, "--start-server-load-latest")
	} else {
		args = append(args, "--start-server", filepath.Join(config.FactorioSavesDir, server.Savefile))
	}

	if config.GlibcCustom == "true" {
		log.Println("Starting server with command: ", config.GlibcLocation, args)
		server.Cmd = exec.Command(config.GlibcLocation, args...)
	} else {
		log.Println("Starting server with command: ", config.FactorioBinary, args)
		server.Cmd = exec.Command(config.FactorioBinary, args...)
	}

	server.StdOut, err = server.Cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error opening stdout pipe: %s", err)
		return err
	}

	server.StdIn, err = server.Cmd.StdinPipe()
	if err != nil {
		log.Printf("Error opening stdin pipe: %s", err)
		return err
	}

	server.StdErr, err = server.Cmd.StderrPipe()
	if err != nil {
		log.Printf("Error opening stderr pipe: %s", err)
		return err
	}

	go server.parseRunningCommand(server.StdOut)
	go server.parseRunningCommand(server.StdErr)

	err = server.Cmd.Start()
	if err != nil {
		log.Printf("Factorio process failed to start: %s", err)
		return err
	}
	server.SetRunning(true)

	err = server.Cmd.Wait()
	log.Printf("Factorio process is closed")
	server.SetRunning(false)
	if err != nil {
		log.Printf("Factorio process exited with error: %s", err)
		return err
	}

	return nil
}

func (server *Server) parseRunningCommand(std io.ReadCloser) (err error) {
	stdScanner := bufio.NewScanner(std)
	for stdScanner.Scan() {
		text := stdScanner.Text()

		log.Printf("Factorio Server: %s", text)
		if err := server.writeLog(text); err != nil {
			log.Printf("Error: %s", err)
		}

		// send the reported line per websocket
		wsRoom := websocket.WebsocketHub.GetRoom("gamelog")
		go wsRoom.Send(text)

		line := strings.Fields(text)
		// Ensure logline slice is in bounds
		if len(line) > 1 {
			// Check if Factorio Server reports any errors if so handle it
			if line[1] == "Error" {
				err := server.checkLogError(line)
				if err != nil {
					log.Printf("Error checking Factorio Server Error: %s", err)
				}
			}
			// If rcon port opens indicated in log connect to rcon
			rconLog := "Starting RCON interface at IP"
			// check if slice index is greater than 2 to prevent panic
			if len(line) > 2 {
				// log line for opened rcon connection
				if strings.Contains(text, rconLog) {
					log.Printf("Rcon running on Factorio Server")
					err = connectRC()
					if err != nil {
						log.Printf("Error: %s", err)
					}
				}

				server.checkProcessHealth(text)
			}
		}
	}
	if err := stdScanner.Err(); err != nil {
		log.Printf("Error reading std buffer: %s", err)
		return err
	}
	return nil
}

func (server *Server) writeLog(logline string) error {
	config := bootstrap.GetConfig()
	logfileName := config.ConsoleLogFile
	file, err := os.OpenFile(logfileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Cannot open logfile %s for appending Factorio Server output: %s", logfileName, err)
		return err
	}
	defer file.Close()

	logline = logline + "\n"

	if _, err = file.WriteString(logline); err != nil {
		log.Printf("Error appending to %s: %s", logfileName, err)
		return err
	}

	return nil
}

func (server *Server) checkLogError(logline []string) error {
	// TODO Handle errors generated by running Factorio Server
	log.Println(logline)

	return nil
}

func init() {
	websocket.WebsocketHub.RegisterControlHandler <- serverWebsocketControl
}

// react to websocket control messages and run the command if it is requested
func serverWebsocketControl(controls websocket.WsControls) {
	log.Println(controls)
	if controls.Type == "command" {
		command := controls.Value
		server := GetFactorioServer()
		if server.GetRunning() {
			log.Printf("Received command: %v", command)

			reqId, err := server.Rcon.Write(command)
			if err != nil {
				log.Printf("Error sending rcon command: %s", err)
				return
			}

			log.Printf("Command send to Factorio: %s, with rcon request id: %v", command, reqId)
		}
	}
}
