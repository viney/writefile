package main

import (
	"html/template"
	"net/http"
)

func init() {
	RegisterHandle("/", indexHandle)
	RegisterHandle("/add", addHandle)
	RegisterHandle("/show", showHandle)
}

func RegisterHandle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleCommon(w, r, handler)
	})
}

func handleCommon(w http.ResponseWriter, r *http.Request, handler func(w http.ResponseWriter, r *http.Request)) {
	handler(w, r)
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

var (
	temp = template.Must(template.ParseFiles("index.html", "show.html"))
	ms   = New()
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "index.html", nil); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	return
}

func addHandle(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	if len(message) == 0 {
		w.Write([]byte("message is null"))
		return
	}

	if err := ms.Add([]byte(message)); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, "/show", http.StatusFound)
	return
}

func showHandle(w http.ResponseWriter, r *http.Request) {
	data, err := ms.Show()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := temp.ExecuteTemplate(w, "show.html", template.FuncMap{"Message": string(data)}); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	return
}

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
