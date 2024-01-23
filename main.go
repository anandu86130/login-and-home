package main

import (
	"html/template"
	"net/http"

	"github.com/icza/session"
)

var tmpl *template.Template
var err string
var data = "Anandu"

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func clearCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, no-transform, must-revalidate, private, max-age=0")
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	clearCache(w, r)
	sess := session.Get(r)
	if sess == nil {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else {
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	}
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	clearCache(w, r)
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "Anandu" && password == "Sonu" {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"username": username},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		err = "invalid username or password"
		tmpl.ExecuteTemplate(w, "index.html", err)
	}
}

func welcomeHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	clearCache(w, r)
	if sess != nil {
		tmpl.ExecuteTemplate(w, "welcome.html", data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	clearCache(w, r)
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/welcome", welcomeHandle)
	http.HandleFunc("/logout", logoutHandle)
	http.ListenAndServe(":9999", nil)
}
