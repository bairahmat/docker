package controllers

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	// "log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/securecookie"
	"github.com/kanciogo/kancio-chat/models"
	"github.com/kanciogo/kancio-chat/session"
	"golang.org/x/crypto/bcrypt"

)

//Lama
var dbUsers = map[string]models.Users{} // user ID, user
var dbSessions = map[string]string{}

var cookiehandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setsesi(nama_user string, res http.ResponseWriter) {
	value := map[string]string{
		"name": nama_user,
	}
	if encoded, err := cookiehandler.Encode("session", value); err == nil {
		cookie_ku := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(res, cookie_ku)
	}
}

func connect() *mgo.Session {
	var session, err = mgo.Dial("172.17.0.2")
	if err != nil {
		os.Exit(0)
	}
	return session
}

type Controllers struct {
	tpl *template.Template
}

func NewControlers(t *template.Template) *Controllers {
	return &Controllers{t}
}

//handlesrs

func (c Controllers) Namauser(r *http.Request) (name_usernya string) {
	if cookie_ini, err := r.Cookie("session"); err == nil {
		nilai_cookie := make(map[string]string)
		if err = cookiehandler.Decode("session", cookie_ini.Value, &nilai_cookie); err == nil {
			name_usernya = nilai_cookie["name"]
		}

	}
	return name_usernya
}


func (c Controllers) Daftar(w http.ResponseWriter, r *http.Request) {
	sessions := connect()
	defer sessions.Close()
	collection := sessions.DB("kancio").C("user")
	data := models.Users{}

	if r.Method == http.MethodPost {
		Username := r.FormValue("Username")
		Nama := r.FormValue("Nama")
		Email := r.FormValue("Email")
		Password := r.FormValue("Password")
		Jk := r.FormValue("Jk")
		mf, fh, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		//membua sha untuk nama file
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// mebuat file baru
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "image", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		//Copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		collection.Find(bson.M{"Email": Email}).One(&data)
		if Email == data.Email {
			http.Error(w, "Email ini sudah terdaftar", http.StatusForbidden)
			return
		}
		collection.Find(bson.M{"Username": Username}).One(&data)
		if Username == data.Username {
			http.Error(w, "Username ini sudah terdaftar", http.StatusForbidden)
			return
		}
		bs, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Gak bisa generate password", http.StatusInternalServerError)
			return
		}
		index := mgo.Index{
			Key:        []string{"Username"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = collection.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		data = models.Users{bson.NewObjectId(), Username, Nama, Email, bs, Jk, fname}
		fmt.Println(data)
		simpan := collection.Insert(data)
		if simpan != nil {
			http.Error(w, "Gak bisa nyimpan data", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	c.tpl.ExecuteTemplate(w, "daftar.html", data)
}

func (c Controllers) Login(w http.ResponseWriter, r *http.Request) {
	if session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	sessions := connect()
	defer sessions.Close()
	collection := sessions.DB("kancio").C("user")
	data := models.Users{}
	if r.Method == http.MethodPost {
		Username := r.FormValue("Username")
		Password := r.FormValue("Password")
		collection.Find(bson.M{"Username": Username}).One(&data)
		err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(Password))
		if err != nil {
			http.Error(w, "Username atau password salah", http.StatusForbidden)
			return
		}
		setsesi(data.Username, w)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "login.html", data)
}

func (c Controllers) Chat(w http.ResponseWriter, r *http.Request) {
	e := r
	fmt.Println(e)
	c.tpl.ExecuteTemplate(w, "chat.html", nil)
}

func (c Controllers) Home(w http.ResponseWriter, r *http.Request) {
	akses := c.Namauser(r)
	if akses == "" {
		fmt.Println("Nilai Akses", akses)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	sessions := connect()
	defer sessions.Close()
	collection := sessions.DB("kancio").C("user")
	var data_akun []models.Users
	err := collection.Find(bson.M{"Username": bson.M{"$eq": akses}}).All(&data_akun)
	if err != nil {
		fmt.Println("gagal mengambil data")
	}
	// idUser := collection.Find(bson.M{"_id"})
	data := models.Home{akses, data_akun}
	c.tpl.ExecuteTemplate(w, "home.html", data)
}
func (c Controllers) Logout(w http.ResponseWriter, req *http.Request) {
	ck, _ := req.Cookie("session")
	fmt.Println(ck)
	// delete the session
	delete(session.Sessions, ck.Value)
	// remove the cookie
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, ck)

	// clean up session.Sessions
	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		go session.Clean()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
