package postmodel

import (
	"go-rest-api/comment"
	"time"
)

type Post struct {
	ID        int               `gorm:"primary_key" json:"id"`
	Content   string            `gorm:"type:varchar(1000)" json:"content"`
	Comments  []comment.Comment `gorm:"ForeignKey:PostId;AssociationForeignKey:ID" json:"comments"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) PostID() int {
	return p.ID
}
