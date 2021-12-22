package handler

import (
	"flag"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	tpb "github.com/tasnuvatina/grpc-blog/proto/todo"
	upb "github.com/tasnuvatina/grpc-blog/proto/user"
)

const sessionName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	sess      *sessions.CookieStore
	tc        tpb.TodoServiceClient
	uc        upb.TodoServiceClient
	bc        bpb.BlogServiceClient
}

func New(decoder *schema.Decoder, sess *sessions.CookieStore, tc tpb.TodoServiceClient, uc upb.TodoServiceClient,bc bpb.BlogServiceClient) *mux.Router {
	h := &Handler{
		decoder: decoder,
		sess:    sess,
		tc:      tc,
		uc:      uc,
		bc:bc,
	}
	h.parseTemplates()

	// for serving static files
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/", h.BlogHome)
	r.HandleFunc("/signup", h.Signup)
	r.HandleFunc("/register", h.Register)
	r.HandleFunc("/signin", h.Signin)
	r.HandleFunc("/login", h.Login)
	s := r.NewRoute().Subrouter()
	// blogs
	r.HandleFunc("/create", h.CreateNewBlog)
	r.HandleFunc("/blog/store", h.StoreNewBlog)
	r.HandleFunc("/blog/{blog:[0-9]+}/read", h.ReadBlog)
	r.HandleFunc("/blog/{blog:[0-9]+}/edit", h.EditBlog)
	r.HandleFunc("/blog/{blog:[0-9]+}/update", h.Updateblog)
	r.HandleFunc("/blog/{blog:[0-9]+}/delete", h.DeleteBlog)

	// s.HandleFunc("/todos/create", h.CreateTodo)
	// s.HandleFunc("/todos/store", h.StoreTodo)
	s.HandleFunc("/logout", h.Logout)

	s.Use(h.authMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := h.templates.ExecuteTemplate(rw, "404.html", nil); err != nil {
			http.Error(rw, "page not found", http.StatusInternalServerError)
			return
		}
	})
	return r

}

func (h *Handler) parseTemplates() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/create-todo.html",
		"cms/assets/templates/404.html",
		"cms/assets/templates/sign-up.html",
		"cms/assets/templates/login.html",
		"cms/assets/templates/write-blog.html",
		"cms/assets/templates/update-blog.html",
		"cms/assets/templates/blog-home.html",
		"cms/assets/templates/blog-page.html",
	))
}

//middleware function to check if the header token exists

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
		return
		// session,err := h.sess.Get(r,sessionName)
		// if err!=nil{
		// 	log.Fatal(err)
		// }

		// authUserID := session.Values["authUserId"]
		// if authUserID !=""{
		// 	next.ServeHTTP(rw,r)
		// }else{
		// 	http.Error(rw,"unauthorized access",http.StatusUnauthorized)
		// }
	})
}

//middelware function to set token

func setToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if strings.Contains(url, "/todos/") {
			r.Header.Set("token", "sdhfghgfh4354hg")
		}
		next.ServeHTTP(rw, r)

	})
}
