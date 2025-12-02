package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Application interface {
    Start() error
    Shutdown(ctx context.Context) error
}

type App struct {
    Port   string
    Server *http.Server
	
}

func NewApp(port string) (*App, error) {
    if port == "" {
        return nil, fmt.Errorf("missing port")
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })

    srv := &http.Server{
        Addr:    ":" + port,
        Handler: mux,
    }

    return &App{
        Port:   port,
        Server: srv,
    }, nil
}

func (a *App) Start() error {
    log.Printf("start server in %s\n", a.Port)
    return a.Server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
    log.Println("Stopping server...")
    return a.Server.Shutdown(ctx)
}
