package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db,err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/todo_app?parseTime=true")
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	fmt.Println("MySQLに接続成功")

	todoTableSQL := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL
    );`
    _, err = db.Exec(todoTableSQL)
    if err != nil {
        log.Fatal("error:", err)
    }
    fmt.Println("ready")

	http.HandleFunc("/add", add)
	http.HandleFunc("/list", list)
	fmt.Println("waiting for requests...")
	http.ListenAndServe(":8080", nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	_, err := db.Exec("INSERT INTO tasks (title) VALUES (?)", title)
	if err != nil {
		fmt.Printf("Added failed: %v\n", err)
		fmt.Fprintf(w, "Added failed: %v\n", err)
		return
	}
	fmt.Printf("Added: %s\n", title)
	fmt.Fprintf(w, "Added: %s\n", title)
}

func list(w http.ResponseWriter, r *http.Request) {
	rows,err := db.Query("SELECT id, title FROM tasks")
	if err != nil {
		fmt.Fprintf(w,"Loading error: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Fprintln(w, "--タスク一覧--:")

	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			fmt.Fprintf(w, "Loading error: %v\n", err)
			return
		}

		fmt.Fprintf(w, "%d: %s\n", id, title)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(w, "Loading error: %v\n", err)
	}
}