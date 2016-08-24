package model

import "time"

type User struct {
	ID       string   `json:"id"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Sns      *Sns     `json:"sns"`
	Tags     []string `json:"tags"`
}

type Sns struct {
	Email    string `json:"email,omitempty"`
	Twitter  string `json:"twitter,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	Line     string `json:"line,omitempty"`
}

type SpareTime struct {
	ID      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Comment string    `json:"comment"`
	Tags    []string  `json:"tags"`
}
