package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type PersonID uuid.UUID

func NewPersonID() (PersonID, error) {
	uuid, err := uuid.NewUUID()
	return PersonID(uuid), err
}

func (pid PersonID) String() string {
	return uuid.UUID(pid).String()
}

func (pid *PersonID) Scan(value interface{}) error {
	uuid := uuid.UUID{}
	if err := uuid.Scan(value); err != nil {
		return err
	}

	*pid = PersonID(uuid)

	return nil
}

func (pid PersonID) Value() (driver.Value, error) {
	return pid.String(), nil
}

type Person struct {
	ID       PersonID  `gorm:"column:id;type:uuid;primary_key"`
	Name     string    `gorm:"column:name;not null"`
	Posts    []Post    `gorm:"foreignkey:PersonID"`
	CreateAt time.Time `gorm:"column:create_at;not null"`
	UpdateAt time.Time `gorm:"column:update_at;not null"`
}

type PostID uuid.UUID

func NewPostID() (PostID, error) {
	uuid, err := uuid.NewUUID()
	return PostID(uuid), err
}

func (pid PostID) String() string {
	return uuid.UUID(pid).String()
}

func (pid *PostID) Scan(value interface{}) error {
	uuid := uuid.UUID{}
	if err := uuid.Scan(value); err != nil {
		return err
	}

	*pid = PostID(uuid)

	return nil
}

func (pid PostID) Value() (driver.Value, error) {
	return pid.String(), nil
}

type Post struct {
	ID       PostID    `gorm:"column:id;type:uuid;primary_key"`
	PersonID PersonID  `gorm:"column:person_id;type:uuid"`
	Text     string    `gorm:"column:text;not null"`
	CreateAt time.Time `gorm:"column:create_at;not null"`
	UpdateAt time.Time `gorm:"column:update_at;not null"`
}

func NewPerson(name string) Person {
	id, err := NewPersonID()
	if err != nil {
		panic("failed create PersonID")
	}

	now := time.Now().UTC()

	return Person{
		ID:       id,
		Name:     name,
		CreateAt: now,
		UpdateAt: now,
	}
}

func NewPost(pid PersonID, text string) Post {
	id, err := NewPostID()
	if err != nil {
		panic("failed create PostID")
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
