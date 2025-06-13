package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type LeaveRequest struct {
	ID         int
	Name       string
	EmpID      string
	Position   string
	Department string
	Type       string
	StartDate  string
	EndDate    string
	Reason     string
}

var tmpl = template.Must(template.ParseFiles("static/index.html"))

func main() {
	db, err := sql.Open("sqlite3", "leaves.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS requests (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT,
                emp_id TEXT,
                position TEXT,
                department TEXT,
                type TEXT,
                start_date TEXT,
                end_date TEXT,
                reason TEXT
        )`); err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			_, err := db.Exec("INSERT INTO requests(name, emp_id, position, department, type, start_date, end_date, reason) VALUES(?,?,?,?,?,?,?,?)",
				r.FormValue("name"), r.FormValue("empid"), r.FormValue("position"), r.FormValue("department"), r.FormValue("type"), r.FormValue("start"), r.FormValue("end"), r.FormValue("reason"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, "OK")
		case http.MethodGet:
			rows, err := db.Query("SELECT id, name, emp_id, position, department, type, start_date, end_date, reason FROM requests")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			var list []LeaveRequest
			for rows.Next() {
				var lr LeaveRequest
				if err := rows.Scan(&lr.ID, &lr.Name, &lr.EmpID, &lr.Position, &lr.Department, &lr.Type, &lr.StartDate, &lr.EndDate, &lr.Reason); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				list = append(list, lr)
			}
			tmpl.Execute(w, list)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
