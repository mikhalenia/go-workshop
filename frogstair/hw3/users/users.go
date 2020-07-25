package users

import (
	"go-workshop/frogstair/hw3/posts"
	"go-workshop/frogstair/hw3/util"
)

// User struct represents a single user
type User struct {
	ID       string
	Username string
	Password string
	Posts    []*posts.Post
}

// New creates a new user
func New(username, password string) *User {
	user := new(User)
	user.ID = util.GenID(10)
	user.Username = username
	user.Password = password
	return user
}

// Post creates a new post
func (u *User) Post(text string) {
	post := new(posts.Post)
	post.Owner = u.ID
	post.ID = util.GenID(20)
	post.Text = text
	post.Likes = 0
	u.Posts = append(u.Posts, post)
}
