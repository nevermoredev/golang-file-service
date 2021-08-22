package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// App export
type App struct {
    Router *mux.Router
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    var response map[string]interface{}
      r.ParseForm()                     // Parses the request body
      x := r.Form.Get("text")
      log.Fatal(x)
    json.Unmarshal([]byte(x), &response)
    respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
    app.Router = mux.NewRouter()
    app.Router.HandleFunc("/", helloWorldHandler)
}

func (app *App) run() {
    log.Fatal(http.ListenAndServe(":26002", app.Router))
}