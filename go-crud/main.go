package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

// ─── 모델 ───────────────────────────────────────────

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// ─── 저장소 (in-memory) ─────────────────────────────

var users = []User{}
var nextID = 1

// ─── 라우터 ─────────────────────────────────────────

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.URL.Path == "/users" {
        switch r.Method {
        case "GET":
            getUsers(w, r)
        case "POST":
            createUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
        return
    }

    if strings.HasPrefix(r.URL.Path, "/users/") {
        switch r.Method {
        case "PUT":
            updateUser(w, r)
        case "DELETE":
            deleteUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
        return
    }

    http.NotFound(w, r)
}

// ─── 진입점 ─────────────────────────────────────────

func main() {
    http.HandleFunc("/users", handler)
    http.HandleFunc("/users/", handler)

    fmt.Println("서버 시작: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}