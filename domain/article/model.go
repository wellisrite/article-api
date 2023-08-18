package article

import "time"

type Article struct {
	ID        int        `json:"id"`
	Author    string     `json:"author"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ArticleRead struct {
	Query  int    `json:"query"`
	Author string `json:"author"`
}

type ArticleWrite struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
