package models

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	Name     string    `gorm:"column:name;not null"`
	Posts    []Post    `gorm:"foreignkey:PersonID"`
	CreateAt time.Time `gorm:"column:create_at;not null"`
	UpdateAt time.Time `gorm:"column:update_at;not null"`
}

type Post struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	PersonID uuid.UUID `gorm:"column:person_id;type:uuid"`
	Text     string    `gorm:"column:text;not null"`
	CreateAt time.Time `gorm:"column:create_at;not null"`
	UpdateAt time.Time `gorm:"column:update_at;not null"`
}

func NewPerson(name string) Person {
	id, err := uuid.NewUUID()
	if err != nil {
		panic("failed create uuid")
	}

	now := time.Now().UTC()

	return Person{
		ID:       id,
		Name:     name,
		CreateAt: now,
		UpdateAt: now,
	}
}

func NewPost(pid uuid.UUID, text string) Post {
	id, err := uuid.NewUUID()
	if err != nil {
		panic("failed create uuid")
	}

	now := time.Now().UTC()

	return Post{
		ID:       id,
		PersonID: pid,
		Text:     text,
		CreateAt: now,
		UpdateAt: now,
	}
}

func (p *Person) TableName() string {
	return "person"
}

func (p *Post) TableName() string {
	return "post"
}
