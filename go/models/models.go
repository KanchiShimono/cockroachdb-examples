package models

import (
	"github.com/google/uuid"
)

type Person struct {
	ID    uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	Name  string    `gorm:"column:name;not null"`
	Posts []Post    `gorm:"foreignkey:PersonID"`
}

type Post struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	PersonID uuid.UUID `gorm:"column:person_id;type:uuid"`
	Text     string    `gorm:"column:text;not null"`
}

func NewPerson(name string) Person {
	id, err := uuid.NewUUID()
	if err != nil {
		panic("failed create uuid")
	}

	return Person{
		ID:   id,
		Name: name,
	}
}

func NewPost(pid uuid.UUID, text string) Post {
	id, err := uuid.NewUUID()
	if err != nil {
		panic("failed create uuid")
	}

	return Post{
		ID:       id,
		PersonID: pid,
		Text:     text,
	}
}

func (p *Person) TableName() string {
	return "person"
}

func (p *Post) TableName() string {
	return "post"
}
