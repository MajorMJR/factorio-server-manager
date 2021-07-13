package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/OpenFactorioServerManager/factorio-server-manager/bootstrap"
	"github.com/OpenFactorioServerManager/factorio-server-manager/factorio"
	"github.com/gorilla/sessions"

	"github.com/gorilla/mux"
)

const readHttpBodyError = "Could not read the Request Body."

type JSONResponseFileInput struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,string"`
	Error     string      `json:"error"`
	ErrorKeys []int       `json:"errorkeys"`
}

func WriteResponse(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error writing response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ReadRequestBody(w http.ResponseWriter, r *http.Request) (body []byte, resp interface{}, err error) {
	if r.Body == nil {
		resp = fmt.Sprintf("%s: no request body", readHttpBodyError)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		err = errors.New("no request body")
		return
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		resp = fmt.Sprintf("%s: %s", readHttpBodyError, err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func ReadSessionStore(w http.ResponseWriter, r *http.Request, name string) (session *sessions.Session, resp interface{}, err error) {
	session, err = sessionStore.Get(r, name)
	if err != nil {
		resp = fmt.Sprintf("Error reading session cookie [%s]: %s", name, err)
		log.Println(resp)
		if session != nil {
			session.Options.MaxAge = -1
			err2 := session.Save(r, w)
			if err2 != nil {
				log.Printf("Error deleting session cookie: %s", err2)
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
	return
}

func SaveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) (resp interface{}, err error) {
	err = session.Save(r, w)
	if err != nil {
		resp = fmt.Sprintf("Error saving session cookie: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

// Lists all save files in the factorio/saves directory
func ListSaves(w http.ResponseWriter, r *http.Request) {
	var resp interface{}
	config := bootstrap.GetConfig()
	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	savesList, err := factorio.ListSaves(config.FactorioSavesDir)
	if err != nil {
		resp = fmt.Sprintf("Error listing save files: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loadLatest := factorio.Save{Name: "Load Latest"}
	savesList = append(savesList, loadLatest)

	resp = savesList
}

func DLSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	config := bootstrap.GetConfig()
	vars := mux.Vars(r)
	save := vars["save"]
	saveName := filepath.Join(config.FactorioSavesDir, save)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", save))
	log.Printf("%s downloading: %s", r.Host, saveName)

	http.ServeFile(w, r, saveName)
}

func UploadSave(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	log.Println("Uploading save file")

	r.ParseMultipartForm(32 << 20)
	config := bootstrap.GetConfig()

	for _, saveFile := range r.MultipartForm.File["savefile"] {
		ext := filepath.Ext(saveFile.Filename)
		if ext != ".zip" {
			// Only zip-files allowed
			resp = fmt.Sprintf("Fileformat {%s} is not allowed", ext)
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		file, err := saveFile.Open()
		if err != nil {
			resp = fmt.Sprintf("Error opening uploaded saveFile: %s", err)
			log.Println(resp)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		out, err := os.Create(filepath.Join(config.FactorioSavesDir, saveFile.Filename))
		if err != nil {
			resp = fmt.Sprintf("Error creating new savefile to copy uploaded on to: %s", err)
			log.Println(resp)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			resp = fmt.Sprintf("Error coping uploaded file to created file on disk: %s", err)
			log.Println(resp)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	resp = "Uploading files successful"
}

// Deletes provided save
func RemoveSave(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	name := vars["save"]

	save, err := factorio.FindSave(name)
	if err != nil {
		resp = fmt.Sprintf("Error finding save {%s}: %s", name, err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = save.Remove()
	if err != nil {
		resp = fmt.Sprintf("Error removing save {%s}: %s", name, err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// save was removed
	resp = fmt.Sprintf("Removed save: %s", save.Name)
}

// Launches Factorio server binary with --create flag to create save
// Url must include save name for creation of savefile
func CreateSaveHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	vars := mux.Vars(r)
	saveName := vars["save"]

	if saveName == "" {
		resp = fmt.Sprintf("Error creating save, no save name provided: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	config := bootstrap.GetConfig()
	saveFile := filepath.Join(config.FactorioSavesDir, saveName)
	cmdOut, err := factorio.CreateSave(saveFile)
	if err != nil {
		resp = fmt.Sprintf("Error creating save {%s}: %s", saveName, err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp = fmt.Sprintf("Save %s created successfully. Command output: \n%s", saveName, cmdOut)
}

// LogTail returns last lines of the factorio-current.log file
func LogTail(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	config := bootstrap.GetConfig()
	resp, err = factorio.TailLog()
	if err != nil {
		resp = fmt.Sprintf("Could not tail %s: %s", config.FactorioLog, err)
		return
	}
}

// LoadConfig returns JSON response of config.ini file
func LoadConfig(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	config := bootstrap.GetConfig()
	configContents, err := factorio.LoadConfig(config.FactorioConfigFile)
	if err != nil {
		resp = fmt.Sprintf("Could not retrieve config.ini: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp = configContents

	log.Printf("Sent config.ini response")
}

func StartServer(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}
	var server = factorio.GetFactorioServer()
	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if server.GetRunning() {
		resp = "Factorio server is already running"
		w.WriteHeader(http.StatusConflict)
		return
	}

	log.Printf("Starting Factorio server.")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}

	log.Printf("Starting Factorio server with settings: %v", string(body))

	err = json.Unmarshal(body, &server)
	if err != nil {
		resp = fmt.Sprintf("Error unmarshalling server settings JSON: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if savefile was submitted with request to start server.
	if server.Savefile == "" {
		resp = "Error starting Factorio server: No save file provided"
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		err = server.Run()
		if err != nil {
			log.Printf("Error starting Factorio server: %+v", err)
			return
		}
	}()

	timeout := 0
	for timeout <= 3 {
		time.Sleep(1 * time.Second)
		if server.GetRunning() {
			log.Printf("Running Factorio server detected")
			break
		} else {
			log.Printf("Did not detect running Factorio server attempt: %+v", timeout)
		}

		timeout++
	}

	if server.GetRunning() == false {
		resp = fmt.Sprintf("Error starting Factorio server: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp = fmt.Sprintf("Factorio server with save: %s started on port: %d", server.Savefile, server.Port)
	log.Println(resp)
}

func StopServer(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var server = factorio.GetFactorioServer()
	if server.GetRunning() {
		err := server.Stop()
		if err != nil {
			resp = fmt.Sprintf("Error stopping factorio server: %s", err)
			log.Println(resp)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp = fmt.Sprintf("Factorio server stopped")
		log.Println(resp)
	} else {
		resp = "Factorio server is not running"
		w.WriteHeader(http.StatusConflict)
		return
	}
}

func KillServer(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var server = factorio.GetFactorioServer()
	if server.GetRunning() {
		err := server.Kill()
		if err != nil {
			resp = fmt.Sprintf("Error killing factorio server: %s", err)
			log.Println(resp)
			return
		}

		log.Printf("Killed Factorio server.")
		resp = fmt.Sprintf("Factorio server killed")
	} else {
		resp = "Factorio server is not running"
		w.WriteHeader(http.StatusBadRequest)
	}
}

func CheckServer(w http.ResponseWriter, r *http.Request) {
	defer func() {
		WriteResponse(w, factorio.GetFactorioServer())
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
}

func FactorioVersion(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var server = factorio.GetFactorioServer()
	resp["version"] = server.Version.String()
	resp["base_mod_version"] = server.BaseModVersion
}

// Unmarshall the User object from the given bytearray
// This function has side effects (it will write to resp and to w, in case of an error)
func UnmarshallUserJson(body []byte, w http.ResponseWriter) (user User, resp interface{}, err error) {
	err = json.Unmarshal(body, &user)
	if err != nil {
		resp = fmt.Sprintf("Unable to parse the request body: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

// Handler for the Login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	// add resp to the response
	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}

	user, resp, err := UnmarshallUserJson(body, w)
	if err != nil {
		return
	}

	log.Printf("Logging in user: %s", user.Username)

	err = auth.checkPassword(user.Username, user.Password)
	if err != nil {
		resp = fmt.Sprintf("Password for user %s wrong", user.Username)
		log.Println(resp)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session, resp, err := ReadSessionStore(w, r, "authentication")
	if err != nil {
		return
	}

	session.Values["username"] = user.Username

	resp, err = SaveSession(w, r, session)
	if err != nil {
		return
	}

	log.Printf("User: %s, logged in successfully", user.Username)

	user.Password = ""
	resp = user
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	session, resp, err := ReadSessionStore(w, r, "authentication")
	if err != nil {
		return
	}

	delete(session.Values, "username")

	resp, err = SaveSession(w, r, session)
	if err != nil {
		return
	}

	resp = "User logged out successfully."
}

func GetCurrentLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp interface{}

	// add resp to the response
	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	session, resp, err := ReadSessionStore(w, r, "authentication")
	if err != nil {
		return
	}

	username := session.Values["username"].(string)

	user, err := auth.getUser(username)
	if err != nil {
		resp = fmt.Sprintf("Error getting user: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Password = ""

	resp = user
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	users, err := auth.listUsers()
	if err != nil {
		resp = fmt.Sprintf("Error listing users: %s", err)
		log.Println(resp)
		return
	}

	resp = users
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}

	user, resp, err := UnmarshallUserJson(body, w)
	if err != nil {
		return
	}

	err = auth.addUser(user)
	if err != nil {
		resp = fmt.Sprintf("Error in adding user {%s}: %s", user.Username, err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp = fmt.Sprintf("User: %s successfully added.", user.Username)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}

	user, resp, err := UnmarshallUserJson(body, w)
	if err != nil {
		return
	}

	err = auth.deleteUser(user.Username)
	if err != nil {
		resp = fmt.Sprintf("Error in removing user {%s}, error: %s", user.Username, err)
		log.Println(resp)
		return
	}

	resp = fmt.Sprintf("User: %s successfully removed.", user.Username)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}

	var user struct {
		OldPassword        string `json:"old_password"`
		NewPassword        string `json:"new_password"`
		NewPasswordConfirm string `json:"new_password_confirmation"`
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		resp = fmt.Sprintf("Unable to parse the request body: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// only allow to change its own password
	// get username from session cookie
	session, resp, err := ReadSessionStore(w, r, "authentication")
	if err != nil {
		return
	}

	username := session.Values["username"].(string)

	// check if password for user is correct
	err = auth.checkPassword(username, user.OldPassword)
	if err != nil {
		resp = fmt.Sprintf("Password for user %s wrong", username)
		log.Println(resp)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// only run, when confirmation correct
	if user.NewPassword != user.NewPasswordConfirm {
		resp = fmt.Sprintf("Password confirmation incorrect")
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = auth.changePassword(username, user.NewPassword)
	if err != nil {
		resp = fmt.Sprintf("Error changing password: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp = true
}

// GetServerSettings returns JSON response of server-settings.json file
func GetServerSettings(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var server = factorio.GetFactorioServer()
	resp = server.Settings

	log.Printf("Sent server settings response")
}

func UpdateServerSettings(w http.ResponseWriter, r *http.Request) {
	var resp interface{}

	defer func() {
		WriteResponse(w, resp)
	}()

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, resp, err := ReadRequestBody(w, r)
	if err != nil {
		return
	}
	log.Printf("Received settings JSON: %s", body)
	var server = factorio.GetFactorioServer()

	// Race Condition while unmarshal possible
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err = json.Unmarshal(body, &server.Settings)
		wg.Done()
	}()

	// Wait for unmarshal to avoid race condition
	wg.Wait()

	if err != nil {
		resp = fmt.Sprintf("Error unmarhaling server settings JSON: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	settings, err := json.MarshalIndent(&server.Settings, "", "  ")
	if err != nil {
		resp = fmt.Sprintf("Failed to marshal server settings: %s", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	config := bootstrap.GetConfig()
	err = ioutil.WriteFile(config.SettingsFile, settings, 0644)
	if err != nil {
		resp = fmt.Sprintf("Failed to save server settings: %v\n", err)
		log.Println(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Saved Factorio server settings in server-settings.json")

	if (server.Version.Greater(factorio.Version{0, 17, 0})) {
		// save admins to adminJson
		admins, err := json.MarshalIndent(server.Settings["admins"], "", "  ")
		if err != nil {
			resp = fmt.Sprintf("Failed to marshal admins-Setting: %s", err)
			log.Println(resp)
			return
		}

		err = ioutil.WriteFile(config.FactorioAdminFile, admins, 0664)
		if err != nil {
			resp = fmt.Sprintf("Failed to save admins: %s", err)
			log.Println(resp)
			return
		}
	}

	resp = fmt.Sprintf("Settings successfully saved")
}
