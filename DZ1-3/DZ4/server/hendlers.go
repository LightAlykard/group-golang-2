package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"server/models"
	"strconv"

	"github.com/go-chi/chi"
)

func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")
	MessegeID := r.URL.Query().Get("ID")
	if templateName == "" {
		templateName = serv.indexTemplate
	}

	file, err := os.Open(path.Join(serv.rootDir, serv.templatesDir, templateName))
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		serv.SendInternalErr(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	templ, err := template.New("Page").Parse(string(data))
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	if MessegeID != "" {
		MessegeID, _ := strconv.Atoi(MessegeID)
		Messeges, err := models.GetPostItem(serv.db, MessegeID)
		if err != nil {
			serv.SendInternalErr(w, err)
			return
		}
		serv.Page.Messeges = Messeges
		if err := templ.Execute(w, serv.Page); err != nil {
			serv.SendInternalErr(w, err)
			return
		}
	} else {
		Messeges, err := models.GetAllPostItems(serv.db)
		if err != nil {
			serv.SendInternalErr(w, err)
			return
		}
		serv.Page.Messeges = Messeges
		if err := templ.Execute(w, serv.Page); err != nil {
			serv.SendInternalErr(w, err)
			return
		}
	}

}

//NewPostHandler creates new Post in DB
func (serv *Server) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	Messege := models.MessegeItem{}
	_ = json.Unmarshal(data, &Messege)

	if err := Messege.Insert(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(Messege)
	w.Write(data)
}

//Редактирование задачи
func (serv *Server) putPostHandler(w http.ResponseWriter, r *http.Request) {
	MessegeID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, _ := ioutil.ReadAll(r.Body)

	Messege := models.MessegeItem{}
	_ = json.Unmarshal(data, &Messege)
	Messege.ID = MessegeID
	fmt.Println(Messege)
	if err := Messege.Update(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(Messege)
	w.Write(data)
}

//Удаление задачи
func (serv *Server) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	MessegeID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	//fmt.Println(taskID, "deleteTaskHandler")

	Messege := models.MessegeItem{ID: MessegeID}

	if err := Messege.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	w.Write(nil)
}
