package handler

import (
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	tpb "github.com/tasnuvatina/grpc-blog/proto/todo"
)

type Todo struct {
	ID          int64
	Title       string
	Description string
	Iscompleted bool
}
type FormData struct {
	Todo  Todo
	Error map[string]string
}

func (t *Todo) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Title,
			validation.Required.Error("The Title can not be empty"),
		),
		validation.Field(&t.Description,
			validation.Required.Error("The Description can not be empty"),
		),
	)
}

//show create todo form
func (h *Handler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	vErrs := map[string]string{}
	h.LoadCreatedTodoForm(rw, todo, vErrs)
}

//take input from the create todo form
func (h *Handler) StoreTodo(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var todo Todo

	if err := h.decoder.Decode(&todo, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := todo.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()
			}
			fmt.Printf("%#v", vErrs)
			h.LoadCreatedTodoForm(rw, todo, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// file,handler,err:=r.FormFile("Image")
	// if err != nil {
	//     fmt.Println("Error Retrieving the File")
	//     fmt.Println(err)
	//     return
	// }
	// defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// tempFile,err:=ioutil.TempFile("assets/images","upload-*.png")
	// if err != nil {
	//     fmt.Println(err)
	// }
	// defer tempFile.Close()
	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	//     fmt.Println(err)
	// }
	// write this byte array to our temporary file
	// tempFile.Write(fileBytes)

	_, err := h.tc.Create(r.Context(), &tpb.CreateTodoRequest{
		Todo: &tpb.Todo{
			Title:       todo.Title,
			Description: todo.Description,
		},
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)

}

func (h *Handler) LoadCreatedTodoForm(rw http.ResponseWriter, todo Todo, errs map[string]string) {
	form := FormData{
		Todo:  todo,
		Error: errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "create-todo.html", form); err != nil {
		http.Error(rw, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoadEditTodoForm(rw http.ResponseWriter, todo Todo, errs map[string]string) {
	form := FormData{
		Todo:  todo,
		Error: errs,
	}

	if err := h.templates.ExecuteTemplate(rw, "edit-todo.html", form); err != nil {
		http.Error(rw, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}
