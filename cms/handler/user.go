package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	upb "github.com/tasnuvatina/grpc-blog/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	UserName string
	Email    string
	Password string
}
type UserData struct {
	User  User
	Error map[string]string
}

type LoginForm struct {
	Username string
	Password string
	Errors   map[string]string
}

func (l LoginForm) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username, validation.Required.Error("Username is required")),
		validation.Field(&l.Password, validation.Required.Error("Password is required"), validation.Length(6, 30).Error("Password must be min 6 letters long")),
	)
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.UserName,
			validation.Required.Error("The UserName can not be empty"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("The Email can not be empty"),
			is.Email.Error("Please enter a valid email"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("The password can not be empty"),
			validation.Length(6, 15).Error("Password must be 6 characters long"),
		),
	)
}

func (h *Handler) Signup(rw http.ResponseWriter, r *http.Request) {
	form := UserData{}
	h.LoadSignupForm(rw, form)
}

func (h *Handler) Register(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	var form User

	if err := h.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatal(err)
	}

	if err := form.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}

			formData := UserData{
				User:  form,
				Error: vErrs,
			}
			h.LoadSignupForm(rw, formData)

		} 
			// http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		
	}

	// if form data is ok then generate password
	pass, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	res, err := h.uc.Register(r.Context(), &upb.RegisterUserRequest{
		User: &upb.User{
			UserName: form.UserName,
			Email:    form.Email,
			Password: string(pass),
		},
	})

	if err != nil {
		// http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(res)
	http.Redirect(rw, r, "/signin", http.StatusTemporaryRedirect)

}

func (h *Handler) LoadSignupForm(rw http.ResponseWriter, form UserData) {
	if err := h.templates.ExecuteTemplate(rw, "sign-up.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Signin(rw http.ResponseWriter, r *http.Request) {
	form := &LoginForm{}

	if err := h.templates.ExecuteTemplate(rw, "login.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Login(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	var form LoginForm

	if err := h.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatal(err)

	}

	if err := form.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			form.Errors = vErrs
			if err := h.templates.ExecuteTemplate(rw, "login.html", form); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

		} 

		// http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.uc.GetUser(r.Context(), &upb.GetUserRequest{
		UserName: form.Username,
	})

	if err != nil {
		form=LoginForm{}
		form.Errors = map[string]string{"Username": "Invalid Username or password"}
		h.LoadLoginForm(rw, form)
		return
	}


	if res==nil{
		form=LoginForm{}
		form.Errors = map[string]string{"Username": "Invalid Username"}
		h.LoadLoginForm(rw, form)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.GetUser().GetPassword()), []byte(form.Password)); err != nil {
		form.Errors = map[string]string{"Password": "Invalid Password"}
		h.LoadLoginForm(rw, form)
		return
	}

	session, err := h.sess.Get(r, sessionName)
	if err != nil {
		log.Fatal(err)
	}
	session.Values["authUserId"] = res.GetUser().GetID()
	fmt.Printf("sessions %#v", session)
	if err := session.Save(r, rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) LoadLoginForm(rw http.ResponseWriter, form LoginForm) {
	if err := h.templates.ExecuteTemplate(rw, "login.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
	session, err := h.sess.Get(r, sessionName)
	if err != nil {
		log.Fatal(err)
	}
	session.Values["authUserId"] = ""
	fmt.Printf("sessions %#v", session)
	fmt.Print(time.Now().GoString())
	fmt.Print(time.Now().Unix())
	if err := session.Save(r, rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}
