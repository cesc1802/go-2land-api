package comment

import "time"

type Comment struct {
	ID        int    `gorm:"primary_key"`
	Content   string `gorm:"type:varchar(1000)"`
	PostId    int    `gorm:"type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewComment() *Comment {
	return &Comment{}
}
