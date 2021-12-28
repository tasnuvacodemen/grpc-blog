package handler

import (
	"fmt"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
)

type Upvote struct {
	ID     int64
	BlogID int64
	UserID int64
}
type Downvote struct {
	ID     int64
	BlogID int64
	UserID int64
}
type Comment struct {
	ID          int64
	BlogID      int64
	UserID      int64
	UserName    string
	Content     string
	CommentedAt string
}

func (c *Comment) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Content,
			validation.Required.Error("The comment can not be empty"),
		),
	)
}


func (h *Handler) PostComment(rw http.ResponseWriter, r *http.Request)  {
	// getting blogId and userId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := h.GetUserIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// getting user data from user id
	user := h.GetUserStruct(rw,r,userId)

	// getting comment time
	commentTime := time.Now().Format("2006-01-02 15:04:05")


	// parsing form
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	comment :=Comment{}

	if err := h.decoder.Decode(&comment, r.PostForm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// insert comment data
	comment.BlogID=blogId
	comment.UserID=user.ID
	comment.UserName=user.UserName
	comment.CommentedAt=commentTime

		// form validation

		if err := comment.Validate(); err != nil {
			vErrors, ok := err.(validation.Errors)
			if ok {
				vErrs := make(map[string]string)
				for key, value := range vErrors {
					vErrs[key] = value.Error()
				}
				http.Redirect(rw,r,"/",http.StatusTemporaryRedirect)
				return
			} else {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	res, err := h.bc.CommentBlog(r.Context(),&bpb.CommentBlogRequest{
		Comment: &bpb.Comment{
			BlogID: comment.BlogID,
			UserID: comment.UserID,
			UserName: comment.UserName,
			Content: comment.Content,
			CommentedAt: comment.CommentedAt,
		},
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.CommentID ==0{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	url:=fmt.Sprintf("/blog/%v/read",blogId)
	http.Redirect(rw,r,url,http.StatusTemporaryRedirect)
	fmt.Printf("%#v", comment)

}

func (h *Handler)Upvote(rw http.ResponseWriter, r *http.Request)  {
	// getting blogId and userId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := h.GetUserIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if the user already upvoted the blog

	res,err := h.bc.GetUpvote(r.Context(),&bpb.GetUpvoteRequest{
		BlogID: blogId,
		UserID: userId,
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.IsUpvotedId == 0{
		upvoteres,err := h.bc.UpvoteBlog(r.Context(),&bpb.UpvoteBlogRequest{
			Upvote: &bpb.Upvote{
				BlogID: blogId,
				UserID: userId,
			},
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if upvoteres.UpvoteId!=0{
			url:=fmt.Sprintf("/blog/%v/read",blogId)
			http.Redirect(rw,r,url,http.StatusTemporaryRedirect)
			return
		}
		return
	}else {
		_,err := h.bc.RevertUpvoteBlog(r.Context(),&bpb.RevertUpvoteBlogRequest{
			UpvoteId: res.IsUpvotedId,
			UserId: userId,
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		url:=fmt.Sprintf("/blog/%v/read",blogId)
		http.Redirect(rw,r,url,http.StatusTemporaryRedirect)
		return
		
	}
}


func (h *Handler)Downvote(rw http.ResponseWriter, r *http.Request)  {
	// getting blogId and userId from url
	blogId, err := h.GetBlogIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := h.GetUserIdFromUrl(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if the user already upvoted the blog

	res,err := h.bc.GetDownvote(r.Context(),&bpb.GetDownvoteRequest{
		BlogID: blogId,
		UserID: userId,
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.IsDownvotedId == 0{
		downvoteres,err := h.bc.DownVoteBlog(r.Context(),&bpb.DownVoteRequest{
			Downvote: &bpb.Downvote{
				BlogID: blogId,
				UserID: userId,
			},
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if downvoteres.DownvoteId!=0{
			url:=fmt.Sprintf("/blog/%v/read",blogId)
			http.Redirect(rw,r,url,http.StatusTemporaryRedirect)
			return
		}
		return
	}else {
		_,err := h.bc.RevertDownVoteBlog(r.Context(),&bpb.RevertDownVoteBlogRequest{
			DownvoteId: res.IsDownvotedId,
			UserId: userId,
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		url:=fmt.Sprintf("/blog/%v/read",blogId)
		http.Redirect(rw,r,url,http.StatusTemporaryRedirect)
		return
		
	}
}


// chec if the user has upvoted 

func (h *Handler) CheckHasUpvoted(rw http.ResponseWriter, r *http.Request,blogId int64,userId int64) int64  {
	res,err := h.bc.GetUpvote(r.Context(),&bpb.GetUpvoteRequest{
		BlogID: blogId,
		UserID: userId,
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 0
	}
	return res.IsUpvotedId
}

// chec if the user has downvoted 

func (h *Handler) CheckHasDownvoted(rw http.ResponseWriter, r *http.Request,blogId int64,userId int64) int64  {
	res,err := h.bc.GetDownvote(r.Context(),&bpb.GetDownvoteRequest{
		BlogID: blogId,
		UserID: userId,
	})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 0
	}
	return res.IsDownvotedId
}

// get all upvotes,downvotes and comments

func (h *Handler) GetAllUpvotes(rw http.ResponseWriter, r *http.Request,blogId int64) ([]*bpb.Upvote,error)   {
	
	res,err :=h.bc.GetAllUpvote(r.Context(),&bpb.GetAllUpvoteRequest{
		BlogID: blogId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return []*bpb.Upvote{},err
	}
	return res.Upvotes,nil
}

func (h *Handler) GetAllDownvotes(rw http.ResponseWriter, r *http.Request,blogId int64) ([]*bpb.Downvote,error)   {
	
	res,err :=h.bc.GetAllDownvote(r.Context(),&bpb.GetAllDownvoteRequest{
		BlogID: blogId,
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return []*bpb.Downvote{},err
	}
	return res.Downvotes,nil

}

func (h *Handler) GetAllComments(rw http.ResponseWriter, r *http.Request,blogId int64) ([]*bpb.Comment,error)   {
	
	res,err :=h.bc.GetAllComments(r.Context(),&bpb.GetAllCommentsRequest{
		BlogID: blogId,
	})
	if err != nil {
		// fmt.Print(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return []*bpb.Comment{},err
	}
	fmt.Print(res.Comments,res,nil)
	return res.Comments,nil


}

// get all counts

func (h *Handler)GetAllUpvoteCount(rw http.ResponseWriter,r *http.Request,blogId int64)(int64,error)  {
	res,err :=h.bc.GetAllUpvoteCount(r.Context(),&bpb.GetAllUpvoteCountRequest{
		BlogID: blogId,
	})

	if err !=nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 0,err
	}
	return res.UpvoteCount,nil
}

func (h *Handler)GetAllDownvoteCount(rw http.ResponseWriter,r *http.Request,blogId int64)(int64,error)  {
	res,err :=h.bc.GetAllDownvoteCount(r.Context(),&bpb.GetAllDownvoteCountRequest{
		BlogID: blogId,
	})

	if err !=nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 0,err
	}
	return res.DownvoteCount,nil
}
func (h *Handler)GetAllCommentCount(rw http.ResponseWriter,r *http.Request,blogId int64)(int64,error)  {
	res,err :=h.bc.GetAllCommentCount(r.Context(),&bpb.GetAllCommentCountRequest{
		BlogID: blogId,
	})

	if err !=nil{
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 0,err
	}
	return res.CommentCount,nil
}

