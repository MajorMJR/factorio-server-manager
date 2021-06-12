package api

import (
	"github.com/OpenFactorioServerManager/factorio-server-manager/api/websocket"
	"github.com/OpenFactorioServerManager/factorio-server-manager/factorio"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	ServerOff   bool // Set to `true' if factorio server has to be turned off to call this
}

type Routes []Route

func ServerOffMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// only run if server is turned off
		server := factorio.GetFactorioServer()
		if server.GetRunning() {
			http.Error(w, "factorio server still running", http.StatusLocked)
		} else {
			next.ServeHTTP(w, r)
		}
		return
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// create subrouter for authenticated calls
	sr := r.NewRoute().Subrouter()
	sr.Use(AuthMiddleware)

	// API subrouter
	// Serves all JSON REST handlers prefixed with /api
	s := r.PathPrefix("/api").Subrouter()
	s.Use(AuthMiddleware)

	// use subrouter for calls, that run only, when server is turned off
	so := s.NewRoute().Subrouter()
	so.Use(ServerOffMiddleware)

	s.NewRoute().Subrouter()
	for _, route := range apiRoutes {
		var router *mux.Router
		if route.ServerOff {
			router = so
		} else {
			router = s
		}
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// The login handler does not check for authentication.
	r.Path("/api/login").
		Methods("POST").
		Name("LoginUser").
		HandlerFunc(LoginUser)

	// Route for initializing websocket connection
	// Clients connecting to /ws establish websocket connection by upgrading
	// HTTP session.
	// Ensure user is logged in with the AuthorizeHandler middleware
	sr.Path("/ws").
		Methods("GET").
		Name("Websocket").
		Handler(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					websocket.ServeWs(w, r)
				},
			),
		)

	// Serves the frontend application from the app directory
	// Uses basic file server to serve index.html and Javascript application
	// Routes match the ones defined in React frontend application
	r.Path("/login").
		Methods("GET").
		Name("Login").
		Handler(http.StripPrefix("/login", http.FileServer(http.Dir("./app/"))))

	sr.Path("/saves").
		Methods("GET").
		Name("Saves").
		Handler(http.StripPrefix("/saves", http.FileServer(http.Dir("./app/"))))
	sr.Path("/map-generator").
		Methods("GET").
		Name("MapGenerator").
		Handler(http.StripPrefix("/map-generator", http.FileServer(http.Dir("./app/"))))
	sr.Path("/mods").
		Methods("GET").
		Name("Mods").
		Handler(http.StripPrefix("/mods", http.FileServer(http.Dir("./app/"))))
	sr.Path("/server-settings").
		Methods("GET").
		Name("Server settings").
		Handler(http.StripPrefix("/server-settings", http.FileServer(http.Dir("./app/"))))
	sr.Path("/game-settings").
		Methods("GET").
		Name("Game settings").
		Handler(http.StripPrefix("/game-settings", http.FileServer(http.Dir("./app/"))))
	sr.Path("/console").
		Methods("GET").
		Name("Console").
		Handler(http.StripPrefix("/console", http.FileServer(http.Dir("./app/"))))
	sr.Path("/logs").
		Methods("GET").
		Name("Logs").
		Handler(http.StripPrefix("/logs", http.FileServer(http.Dir("./app/"))))
	sr.Path("/user-management").
		Methods("GET").
		Name("User management").
		Handler(http.StripPrefix("/user-management", http.FileServer(http.Dir("./app/"))))
	sr.Path("/help").
		Methods("GET").
		Name("Help").
		Handler(http.StripPrefix("/help", http.FileServer(http.Dir("./app/"))))

	// catch all route
	r.PathPrefix("/").
		Methods("GET").
		Name("Index").
		Handler(http.FileServer(http.Dir("./app/")))

	return r
}

// Defines all API REST endpoints
// All routes are prefixed with /api
var apiRoutes = Routes{
	{
		"ListSaves",
		"GET",
		"/saves/list",
		ListSaves,
		false,
	}, {
		"DlSave",
		"GET",
		"/saves/dl/{save}",
		DLSave,
		false,
	}, {
		"UploadSave",
		"POST",
		"/saves/upload",
		UploadSave,
		false,
	}, {
		"RemoveSave",
		"GET",
		"/saves/rm/{save}",
		RemoveSave,
		false,
	}, {
		"CreateSave",
		"POST",
		"/saves/create/{save}",
		CreateSaveHandler,
		true,
	}, {
		"GenerateMapPreview",
		"POST",
		"/saves/preview",
		GenerateMapPreview,
		true,
	}, {
		"DefaultMapSettings",
		"GET",
		"/saves/default-map-settings",
		DefaultMapSettings,
		false,
	}, {
		"DefaultMapGenSettings",
		"GET",
		"/saves/default-map-gen-settings",
		DefaultMapGenSettings,
		false,
	}, {
		"LoadModsFromSave",
		"POST",
		"/saves/mods",
		LoadModsFromSaveHandler,
		true,
	}, {
		"LogTail",
		"GET",
		"/log/tail",
		LogTail,
		false,
	}, {
		"LoadConfig",
		"GET",
		"/config",
		LoadConfig,
		false,
	}, {
		"StartServer",
		"POST",
		"/server/start",
		StartServer,
		true,
	}, {
		"StopServer",
		"GET",
		"/server/stop",
		StopServer,
		false,
	}, {
		"KillServer",
		"GET",
		"/server/kill",
		KillServer,
		false,
	}, {
		"RunningServer",
		"GET",
		"/server/status",
		CheckServer,
		false,
	}, {
		"FactorioVersion",
		"GET",
		"/server/facVersion",
		FactorioVersion,
		false,
	}, {
		"LogoutUser",
		"GET",
		"/logout",
		LogoutUser,
		false,
	}, {
		"StatusUser",
		"GET",
		"/user/status",
		GetCurrentLogin,
		false,
	}, {
		"ListUsers",
		"GET",
		"/user/list",
		ListUsers,
		false,
	}, {
		"AddUser",
		"POST",
		"/user/add",
		AddUser,
		false,
	}, {
		"RemoveUser",
		"POST",
		"/user/remove",
		RemoveUser,
		false,
	}, {
		"ChangePassword",
		"POST",
		"/user/password",
		ChangePassword,
		false,
	}, {
		"GetServerSettings",
		"GET",
		"/settings",
		GetServerSettings,
		false,
	}, {
		"UpdateServerSettings",
		"POST",
		"/settings/update",
		UpdateServerSettings,
		false,
	},
	// Mod Portal Stuff
	{
		"ModPortalListAllMods",
		"GET",
		"/mods/portal/list",
		ModPortalListModsHandler,
		false,
	}, {
		"ModPortalGetModInfo",
		"GET",
		"/mods/portal/info/{mod}",
		ModPortalModInfoHandler,
		false,
	}, {
		"ModPortalInstallMod",
		"POST",
		"/mods/portal/install",
		ModPortalInstallHandler,
		true,
	}, {
		"ModPortalLogin",
		"POST",
		"/mods/portal/login",
		ModPortalLoginHandler,
		false,
	}, {
		"ModPortalLoginStatus",
		"GET",
		"/mods/portal/loginstatus",
		ModPortalLoginStatusHandler,
		false,
	}, {
		"ModPortalLogout",
		"GET",
		"/mods/portal/logout",
		ModPortalLogoutHandler,
		false,
	}, {
		"ModPortalInstallMultiple",
		"POST",
		"/mods/portal/install/multiple",
		ModPortalInstallMultipleHandler,
		true,
	},
	// Mods Stuff
	{
		"ListInstalledMods",
		"GET",
		"/mods/list",
		ListInstalledModsHandler,
		false,
	}, {
		"ToggleMod",
		"POST",
		"/mods/toggle",
		ModToggleHandler,
		true,
	}, {
		"DeleteMod",
		"POST",
		"/mods/delete",
		ModDeleteHandler,
		true,
	}, {
		"DeleteAllMods",
		"POST",
		"/mods/delete/all",
		ModDeleteAllHandler,
		true,
	}, {
		"UpdateMod",
		"POST",
		"/mods/update",
		ModUpdateHandler,
		true,
	}, {
		"UploadMod",
		"POST",
		"/mods/upload",
		ModUploadHandler,
		true,
	}, {
		"DownloadMods",
		"GET",
		"/mods/download",
		ModDownloadHandler,
		false,
	},
	// Mod Packs
	{
		"ModPacksList",
		"GET",
		"/mods/packs/list",
		ModPackListHandler,
		false,
	}, {
		"ModPackCreate",
		"POST",
		"/mods/packs/create",
		ModPackCreateHandler,
		false,
	}, {
		"ModPackDelete",
		"POST",
		"/mods/packs/{modpack}/delete",
		ModPackDeleteHandler,
		false,
	}, {
		"ModPackDownload",
		"GET",
		"/mods/packs/{modpack}/download",
		ModPackDownloadHandler,
		false,
	}, {
		"LoadModPack",
		"POST",
		"/mods/packs/{modpack}/load",
		ModPackLoadHandler,
		true,
	},
	// Mods inside Mod Packs
	{
		"ModPackListMods",
		"GET",
		"/mods/packs/{modpack}/list",
		ModPackModListHandler,
		false,
	}, {
		"ModPackToggleMod",
		"POST",
		"/mods/packs/{modpack}/mod/toggle",
		ModPackModToggleHandler,
		false,
	}, {
		"ModPackDeleteMod",
		"POST",
		"/mods/packs/{modpack}/mod/delete",
		ModPackModDeleteHandler,
		false,
	}, {
		"ModPackDeleteAllMod",
		"POST",
		"/mods/packs/{modpack}/mod/delete/all",
		ModPackModDeleteAllHandler,
		false,
	}, {
		"ModPackUpdateMod",
		"POST",
		"/mods/packs/{modpack}/mod/update",
		ModPackModUpdateHandler,
		false,
	}, {
		"ModPackUploadMod",
		"POST",
		"/mods/packs/{modpack}/mod/upload",
		ModPackModUploadHandler,
		false,
	}, {
		"ModPackModPortalInstallMod",
		"POST",
		"/mods/packs/{modpack}/portal/install",
		ModPackModPortalInstallHandler,
		false,
	}, {
		"ModPackModPortalInstallMultiple",
		"POST",
		"/mods/packs/{modpack}/portal/install/multiple",
		ModPackModPortalInstallMultipleHandler,
		false,
	},
}
