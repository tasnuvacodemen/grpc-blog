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



