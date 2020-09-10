package server

import (
	"DZ6/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func (srv *Server) defineRoutes() {
	srv.router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r) })
	})
	srv.router.Route("/api/v1/posts/", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { srv.PostCreate(w, r) })
		r.Put("/{messegeID}", func(w http.ResponseWriter, r *http.Request) { srv.PostUpdate(w, r) })
		r.Delete("/{messegeID}", func(w http.ResponseWriter, r *http.Request) { srv.PostDelete(w, r) })
		r.Get("/{messegeID}", func(w http.ResponseWriter, r *http.Request) { srv.PostGet(w, r) })
	})
}
func (srv *Server) GetPage(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./templates/messeges.html")
	data, _ := ioutil.ReadAll(file)

	tmp, err := template.New("messeges").Parse(string(data))
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	NewPosts, err := models.Get(srv.Ctx, srv.DBMongo)
	if err != nil {
		srv.lg.WithError(err).Fatal("NewPost err")
		return
	}
	srv.Posts = *NewPosts
	tmp.ExecuteTemplate(w, "messeges", srv)

}
func (srv *Server) PostCreate(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	messege := models.Messege{}
	_ = json.Unmarshal(data, &messege)

	id, err := srv.Messeges.NewPost(srv.Ctx, srv.DBMongo, messege.Header, messege.Text)
	if err != nil {
		srv.lg.WithError(err).Fatal("NewPost err")
		return
	}
	data, _ = json.Marshal(srv.Messeges[id])

	w.Write(data)
}

func (srv *Server) PostGet(w http.ResponseWriter, r *http.Request) {
	messegeID := chi.URLParam(r, "messegeID")
	messege, ok := srv.Messeges[messegeID]
	if !ok {
		srv.lg.Warningln("no such ID=", messegeID)
		return
	}

	data, err := json.Marshal(messege)
	if err != nil {
		srv.lg.WithError(err).Fatal("PostGet Marshal err")
		return
	}

	w.Write(data)
}

//PostDelete -
func (srv *Server) PostDelete(w http.ResponseWriter, r *http.Request) {
	messegeID := chi.URLParam(r, "messegeID")
	srv.lg.Debug("Del", messegeID)
	messege, ok := srv.Messeges[messegeID]
	if !ok {
		srv.lg.Warningln("no such ID=", messegeID)
		return
	}

	data, _ := json.Marshal(messege)

	if err := srv.Messeges.DeletePost(srv.Ctx, srv.DBMongo, messegeID); err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)
}

func (srv *Server) PostUpdate(w http.ResponseWriter, r *http.Request) {
	messegeID := chi.URLParam(r, "messegeID")
	srv.lg.Debug("upd ID=", messegeID)
	_, ok := srv.Messeges[messegeID]
	if !ok {
		srv.lg.Warningln("no such ID=", messegeID)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	messege := models.Messege{}
	_ = json.Unmarshal(data, &messege)

	err := srv.Messeges.UpdatePost(srv.Ctx, srv.DBMongo, messegeID, messege.Name, messege.Text)
	if err != nil {
		srv.lg.WithError(err).Fatal("UpdatePost err")
		return
	}

	data, _ = json.Marshal(messege)

	w.Write(data)
}
