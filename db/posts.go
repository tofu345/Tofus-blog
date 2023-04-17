package db

type Post struct {
	BaseModel
	ID       int       `gorm:"primarykey" json:"id"`
	Title    string    `json:"title" gorm:"unique"`
	Slug     string    `json:"slug"  gorm:"unique"`
	Body     string    `json:"body"`
	Author   string    `json:"author"`
	Views    int       `json:"views"`
	Likes    uint64    `json:"likes"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ID"`
}

func GetPostBySlug(slug string) (Post, error) {
	var post Post
	result := DB.First(&post, "slug = ?", slug)
	return post, result.Error
}

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
