package subcommand

import (
	"fmt"

	"github.com/KanchiShimono/cockroachdb-examples/go/models"
	"github.com/jinzhu/gorm"
)

func Insert(db *gorm.DB, numPerson int, numPostPerPerson int) error {
	for i := range make([]int, numPerson) {
		err := db.Transaction(func(tx *gorm.DB) error {
			person := models.NewPerson(fmt.Sprintf("hoge_%d", i))

			if err := db.Create(&person).Error; err != nil {
				return err
			}

			for j := range make([]int, numPostPerPerson) {
				post := models.NewPost(person.ID, fmt.Sprintf("fuga_%d_%d", i, j))
				if err := db.Create(&post).Error; err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}
