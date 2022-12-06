package backend

import "errors"

func GetPostById(object_id int) (Post, error) {
	var post = Post{ID: object_id}
	result := db.First(&post)
	return post, result.Error
}

func DeletePost(object_id int) {
	db.Exec("DELETE FROM posts WHERE id = ?", object_id)
}

// Returns Post Errors as a map
// returns an empty map if none
func (post *Post) Errors() map[string]string {
	errors := map[string]string{}

	if post.Title == "" {
		errors["title"] = "This field is required"
	}

	if post.Body == "" {
		errors["body"] = "This field is required"
	}

	if post.Author == "" {
		errors["author"] = "This field is required"
	}

	return errors
}

// Creates Post and returns errors
func CreatePost(post Post) (map[string]string, error) {
	postErrors := post.Errors()
	if len(postErrors) == 0 {
		if result := db.Create(&post); result.Error != nil {
			return postErrors, errors.New("error creating post")
		} else {
			return postErrors, nil
		}
	} else {
		return postErrors, errors.New("post data invalid")
	}
}
