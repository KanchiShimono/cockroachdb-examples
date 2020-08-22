package subcommand

import (
	"github.com/KanchiShimono/cockroachdb-examples/go/models"
	"github.com/jinzhu/gorm"
)

func Delete(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete([]*models.Post{}).Error; err != nil {
			return err
		}

		if err := tx.Delete([]*models.Person{}).Error; err != nil {
			return err
		}

		return nil
	})
}
