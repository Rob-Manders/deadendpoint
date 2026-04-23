package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	server  *http.Server
	running bool
	mutex   sync.Mutex
}

type CapturedRequest struct {
	Method  string
	Path    string
	Query   map[string][]string
	Headers map[string][]string
	Body    string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartServer(port int, responseCode int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.running {
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := io.ReadAll(r.Body)

		request := CapturedRequest{
			Method:  r.Method,
			Path:    r.URL.Path,
			Query:   r.URL.Query(),
			Headers: r.Header,
			Body:    string(bodyBytes),
		}

		runtime.EventsEmit(a.ctx, "request_received", request)

		w.WriteHeader(responseCode)
	})

	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	a.running = true

	go func() {
		runtime.EventsEmit(a.ctx, "server_running", true)

		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Failed to start server:", err)
		}
	}()
}

func (a *App) StopServer() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.running {
		return
	}

	a.server.Shutdown(a.ctx)
	a.running = false

	runtime.EventsEmit(a.ctx, "server_running", false)
}
