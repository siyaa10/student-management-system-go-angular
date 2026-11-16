package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var (
	students = []Student{}
	nextID   = 1
	mu       sync.Mutex
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/students", studentsHandler)
	mux.HandleFunc("/api/students/", deleteStudentHandler)

	log.Println("Student API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		result := append([]Student{}, students...)
		mu.Unlock()
		json.NewEncoder(w).Encode(result)

	case http.MethodPost:
		var student Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
			return
		}
		student.Name = strings.TrimSpace(student.Name)
		student.Email = strings.TrimSpace(student.Email)
		if student.Name == "" || student.Email == "" || student.Age < 1 || student.Age > 120 {
			http.Error(w, `{"error":"Name, email and a valid age are required"}`, http.StatusBadRequest)
			return
		}

		mu.Lock()
		student.ID = nextID
		nextID++
		students = append(students, student)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(student)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", "DELETE")
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/api/students/")
	id, err := strconv.Atoi(idText)
	if err != nil || id < 1 {
		http.Error(w, `{"error":"Invalid student ID"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted"})
			return
		}
	}
	http.Error(w, `{"error":"Student not found"}`, http.StatusNotFound)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
