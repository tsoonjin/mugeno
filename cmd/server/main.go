package main

import (
    "fmt"
    "time"
    "net/http"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/tsoonjin/mugeno/pkg/rss"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

type Config struct {
    MiddlewareTimeout int
    ServerPort string
}

var config Config;

func init() {
    log.Println("Loading environment variables ...")
     if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    serverPort, _ := os.LookupEnv("SERVER_PORT")
    config = Config{
        ServerPort: serverPort,
    }
}

func main() {
    r := chi.NewRouter()
    // A good base middleware stack
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Health check")
        w.Write([]byte("OK"))
    })

    // Handler RSS request
    r.Route("/rss", func(r chi.Router) {
        r.Get("/", ListRSS)
    })

    // Run Server
    http.ListenAndServe(fmt.Sprintf(":%v", config.ServerPort), r)
}


func ListRSS(w http.ResponseWriter, r *http.Request) {
    urls := rss.List("test")
    fmt.Println(urls)
}
