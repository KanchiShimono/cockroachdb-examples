package subcommand

import (
	"log"

	"github.com/KanchiShimono/cockroachdb-examples/go/models"
	"github.com/jinzhu/gorm"
)

func SelectPerson(db *gorm.DB) error {
	persons := []*models.Person{}
	if err := db.Find(&persons).Error; err != nil {
		return err
	}

	for _, p := range persons {
		log.Println(p)
	}

	return nil
}

func SelectPost(db *gorm.DB) error {
	posts := []*models.Post{}
	if err := db.Find(&posts).Error; err != nil {
		return err
	}

	for _, p := range posts {
		log.Println(p)
	}

	return nil
}

func SelectPersonWithPost(db *gorm.DB, num int) error {
	persons := []*models.Person{}
	if err := db.Limit(num).Preload("Posts").Find(&persons).Error; err != nil {
		return err
	}

	for _, p := range persons {
		log.Println(p)
	}

	return nil
}
