package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	// "strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	upb "github.com/tasnuvatina/grpc-blog/proto/user"
)

type Blog struct {
	ID            int64
	AuthorID      int64
	AuthorName    string
	CreatedAt     string
	UpdateAt      string
	PictureString string
	Title         string
	Description   string
	UpvoteCount   int64
	DownvoteCount int64
	CommentsCount int64
}
type BlogFormData struct {
	Blog     Blog
	UserData User
	Error    map[string]string
	Pic      string
}

type BlogList struct {
	Blogs      []*bpb.Blog
	UserData   User
	SearchTerm string
}

type SingleBlogData struct {
	Blog         Blog
	UserData     User
	IsAuthor     bool
	HasUpvoted   bool
	HasDownVoted bool
	Upvotes      []*bpb.Upvote
	Downvotes    []*bpb.Downvote
	Comments     []*bpb.Comment
}

func (b *Blog) Validate() error {
	return validation.ValidateStruct(b,
		validation.Field(&b.Title,
			validation.Required.Error("The Title can not be empty"),
		),
		validation.Field(&b.Description,
			validation.Required.Error("The Description can not be empty"),
		),
	)
}

// Blog home page handler
func (h *Handler) BlogHome(rw http.ResponseWriter, r *http.Request) {
	//getting sear value from url
	searchedValue := r.URL.Query().Get("search")

	userdata := User{}

	// get userdata from session
	userId := h.GetUserIdFromSession(r)

	// get userData
	userdata = h.GetUserStruct(rw, r, userId)
	
	var blogs []*bpb.Blog

	if len(searchedValue) > 0 {
		allBlogs, err := h.bc.ReadAllSearchedBlog(r.Context(), &bpb.ReadAllBlogSearchedRequest{
			SearchValue: searchedValue,
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		blogs=allBlogs.Blogs
	} else {
		//get all blogs
		allBlogs, err := h.bc.ReadAllBlog(r.Context(), &bpb.ReadAllBlogRequest{})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		blogs=allBlogs.Blogs
	}

	
	BlogList := BlogList{
		Blogs:      blogs,
		UserData:   userdata,
		SearchTerm: "",
	}

	if err := h.templates.ExecuteTemplate(rw, "blog-home.html", BlogList); err != nil {
		http.Error(rw, "Unable to load blog home template", http.StatusInternalServerError)
		return
	}

}

// single blog page handler
func (h *Handler) ReadBlog(rw http.ResponseWriter, r *http.Request) {

	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	userId := h.GetUserIdFromSession(r)
	res, err := h.bc.ReadBlog(r.Context(), &bpb.ReadBlogRequest{
		BlogID:   blogId,
		AuthorID: userId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// store responce data in blogInfo struct
	blogInfo := Blog{
		ID:            res.GetBlog().ID,
		AuthorID:      res.GetBlog().AuthorID,
		AuthorName:    res.GetBlog().AuthorName,
		CreatedAt:     res.GetBlog().CreatedAt,
		UpdateAt:      res.GetBlog().UpdateAt,
		PictureString: res.GetBlog().PictureString,
		Title:         res.GetBlog().Title,
		Description:   res.GetBlog().Description,
		UpvoteCount:   res.GetBlog().UpvoteCount,
		DownvoteCount: res.GetBlog().DownvoteCount,
		CommentsCount: res.GetBlog().CommentsCount,
	}
	// get userData
	user := h.GetUserStruct(rw, r, userId)
	// is the user authoer of the blog
	isUserAuthor := false
	if user.ID == blogInfo.AuthorID {
		isUserAuthor = true
	}

	// check is the user upvoted or down voted
	upvoteId := h.CheckHasUpvoted(rw, r, blogId, userId)
	hasUpVoted := false
	if upvoteId != 0 {
		hasUpVoted = true
	}
	downvoteId := h.CheckHasDownvoted(rw, r, blogId, userId)
	hasDownVoted := false
	if downvoteId != 0 {
		hasDownVoted = true
	}

	// get allupvotes
	upvoteRes, err := h.GetAllUpvotes(rw, r, blogId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//get all downvotes
	downvotesres, err := h.GetAllDownvotes(rw, r, blogId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//get all comments
	commentsres, err := h.GetAllComments(rw, r, blogId)
	if err != nil {
		// http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//get all upvote count
	upvoteCount, err := h.GetAllUpvoteCount(rw, r, blogId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	blogInfo.UpvoteCount = upvoteCount

	//get all comments count
	commentsCount, err := h.GetAllCommentCount(rw, r, blogId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	blogInfo.CommentsCount = commentsCount

	// initializ data
	SingleBlogData := SingleBlogData{
		Blog:         blogInfo,
		UserData:     user,
		IsAuthor:     isUserAuthor,
		HasUpvoted:   hasUpVoted,
		HasDownVoted: hasDownVoted,
		Upvotes:      upvoteRes,
		Downvotes:    downvotesres,
		Comments:     commentsres,
	}
	// fmt.Printf("%#v", SingleBlogData)
	// execute template
	if err := h.templates.ExecuteTemplate(rw, "blog-page.html", SingleBlogData); err != nil {
		http.Error(rw, "Unable to load blog page template", http.StatusInternalServerError)
		return
	}

}

// write new blog template execute
func (h *Handler) CreateNewBlog(rw http.ResponseWriter, r *http.Request) {
	blog := Blog{}
	vErrs := map[string]string{}
	h.loadCreateBlogTemplate(rw, blog, vErrs)

}

// take input from write new blog
func (h *Handler) StoreNewBlog(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	blog := Blog{}

	if err := h.decoder.Decode(&blog, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//handling file type input
	file, fileHeader, err := r.FormFile("Picture")
	if err != nil {
		fmt.Println("error retrieving the file from input", err)
		vErrs := map[string]string{
			"file": "error retrieving the file from input",
		}
		h.loadCreateBlogTemplate(rw, blog, vErrs)
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)
	fmt.Println(fileExtension)
	defer file.Close()

	//store file in local file storage
	tempFile, err := ioutil.TempFile("cms/assets/images", "book-*"+fileExtension)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()
	// read file and store in in a variable
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	newfile := tempFile.Name()
	fileName := filepath.Base(newfile)
	if len(fileName) != 0 {
		blog.PictureString = fileName
	}

	// form validation

	if err := blog.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadCreateBlogTemplate(rw, blog, vErrs)
			return
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	//getting user id from session
	userId := h.GetUserIdFromSession(r)
	// get userData
	userdata := h.GetUserStruct(rw, r, userId)

	craetedTime := time.Now().Format("2006-01-02 15:04:05")

	blog.AuthorID = userdata.ID
	blog.AuthorName = userdata.UserName
	blog.CreatedAt = craetedTime

	// fmt.Println("**************************")
	fmt.Printf("%#v", blog)
	writeRes, err := h.bc.WriteBlog(r.Context(), &bpb.WriteBlogRequest{
		Blog: &bpb.Blog{
			AuthorID:      blog.AuthorID,
			AuthorName:    blog.AuthorName,
			CreatedAt:     blog.CreatedAt,
			UpdateAt:      blog.UpdateAt,
			PictureString: blog.PictureString,
			Title:         blog.Title,
			Description:   blog.Description,
			UpvoteCount:   blog.UpvoteCount,
			DownvoteCount: blog.DownvoteCount,
			CommentsCount: blog.CommentsCount,
		},
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if writeRes.ID == 0 {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)

}

// delete book from table
func (h *Handler) DeleteBlog(rw http.ResponseWriter, r *http.Request) {
	//getting blogId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//getting user id from session
	userId := h.GetUserIdFromSession(r)
	// getting blog data from database by id
	blogRes, err := h.bc.ReadBlog(r.Context(), &bpb.ReadBlogRequest{
		BlogID:   blogId,
		AuthorID: userId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// delete blog from db calling grpc service
	_, err = h.bc.DeleteBlog(r.Context(), &bpb.DeleteBlogRequest{
		ID:       blogId,
		AuthorID: userId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//deleting image from local storage
	if err := deleteImage(blogRes.GetBlog().PictureString); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

// edit blog
func (h *Handler) EditBlog(rw http.ResponseWriter, r *http.Request) {
	//getting blogId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//getting user id from session
	userId := h.GetUserIdFromSession(r)
	// get userData
	userdata := h.GetUserStruct(rw, r, userId)
	// getting blog data from database by id
	blogRes, err := h.bc.ReadBlog(r.Context(), &bpb.ReadBlogRequest{
		BlogID:   blogId,
		AuthorID: userId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	blogData := Blog{
		ID:            blogRes.GetBlog().ID,
		AuthorID:      blogRes.GetBlog().AuthorID,
		AuthorName:    blogRes.GetBlog().AuthorName,
		CreatedAt:     blogRes.GetBlog().CreatedAt,
		UpdateAt:      blogRes.GetBlog().UpdateAt,
		PictureString: blogRes.GetBlog().PictureString,
		Title:         blogRes.GetBlog().Title,
		Description:   blogRes.GetBlog().Description,
		UpvoteCount:   blogRes.GetBlog().UpvoteCount,
		DownvoteCount: blogRes.GetBlog().DownvoteCount,
		CommentsCount: blogRes.GetBlog().CommentsCount,
	}
	if blogData.ID == 0 {
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		return
	}
	h.loadEditBlogTemplate(rw, blogData, userdata, map[string]string{})
}

func (h *Handler) Updateblog(rw http.ResponseWriter, r *http.Request) {
	//getting blogId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// parsing form
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	blog := Blog{}

	if err := h.decoder.Decode(&blog, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	blog.ID = blogId
	blog.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	//getting user id from session
	userId := h.GetUserIdFromSession(r)
	// get userData
	userdata := h.GetUserStruct(rw, r, userId)

	// form validation

	if err := blog.Validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[key] = value.Error()
			}
			h.loadEditBlogTemplate(rw, blog, userdata, vErrs)
			return
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	editRes, err := h.bc.EditBlog(r.Context(), &bpb.EditBlogRequest{
		Blog: &bpb.Blog{
			ID:            blog.ID,
			AuthorID:      blog.AuthorID,
			AuthorName:    blog.AuthorName,
			CreatedAt:     blog.CreatedAt,
			UpdateAt:      blog.UpdateAt,
			PictureString: blog.PictureString,
			Title:         blog.Title,
			Description:   blog.Description,
			UpvoteCount:   blog.UpvoteCount,
			DownvoteCount: blog.DownvoteCount,
			CommentsCount: blog.CommentsCount,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if editRes.GetBlog().ID == 0 {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)

}

// load create blog template
func (h *Handler) loadCreateBlogTemplate(rw http.ResponseWriter, blog Blog, vErrs map[string]string) {

	form := BlogFormData{
		Blog:  blog,
		Error: vErrs,
		Pic:   "",
	}

	if err := h.templates.ExecuteTemplate(rw, "write-blog.html", form); err != nil {
		http.Error(rw, "Unable to load write-blog template", http.StatusInternalServerError)
		return
	}
}

// load edit blog
func (h *Handler) loadEditBlogTemplate(rw http.ResponseWriter, blog Blog, user User, vErrs map[string]string) {

	form := BlogFormData{
		Blog:     blog,
		UserData: user,
		Error:    vErrs,
	}

	if err := h.templates.ExecuteTemplate(rw, "update-blog.html", form); err != nil {
		http.Error(rw, "Unable to load update-blog template", http.StatusInternalServerError)
		return
	}
}

// get userdata from session
func (h *Handler) GetUserIdFromSession(r *http.Request) int64 {
	session, err := h.sess.Get(r, sessionName)
	if err != nil {
		log.Fatal(err)
	}

	userId := session.Values["authUserId"]
	if userId != "" {
		userIdInt, _ := userId.(int64)
		return userIdInt
	} else {
		return 0
	}

}

// get userdata from database by id

func (h *Handler) GetUserIdByDataFromDb(rw http.ResponseWriter, r *http.Request, id int64) *upb.GetUserByIdResponce {
	res, err := h.uc.GetUserById(r.Context(), &upb.GetUserByIdRequest{
		ID: int64(id),
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return res
}

// get blogId from url
func (h *Handler) GetBlogIdFromUrl(rw http.ResponseWriter, r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	blogId := vars["blog"]
	if blogId == "" {
		http.Error(rw, "Can not get blog with empty id", http.StatusInternalServerError)
		return 0, errors.New("blog id is empty")
	}
	i, err := strconv.ParseInt(blogId, 10, 64)
	if err != nil {
		return 0, errors.New("blog id is invalid")
	}
	return i, nil
}

// get blogId from url
func (h *Handler) GetUserIdFromUrl(rw http.ResponseWriter, r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	userId := vars["user"]
	if userId == "" {
		http.Error(rw, "Can not get user with empty id", http.StatusInternalServerError)
		return 0, errors.New("blog id is empty")
	}
	i, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, errors.New("user id is invalid")
	}
	return i, nil
}

// get user struct from user data

func (h *Handler) GetUserStruct(rw http.ResponseWriter, r *http.Request, userId int64) User {
	if userId != 0 {
		newUser := h.GetUserIdByDataFromDb(rw, r, userId)
		return User{
			ID:       newUser.GetUser().ID,
			UserName: newUser.GetUser().UserName,
			Email:    newUser.GetUser().Email,
		}
	}
	return User{}
}

// delete image after deleting book from table

func deleteImage(imgName string) error {
	path := fmt.Sprintf("cms/assets/images/%s", imgName)

	if err := os.Remove(path); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
