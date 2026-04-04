package main

import (
	"net/http"
	"log"
	"html/template"
)

func topPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/top.html")
	if err != nil {
		http.Error(w, "parse Error", http.StatusInternalServerError)
		log.Println("parse error:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "execute Error", http.StatusInternalServerError)
		log.Println("execute error:", err)
		return
	}
}

func enterRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.FormValue("room_id")
	if roomID == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/"+roomID, http.StatusSeeOther)
}