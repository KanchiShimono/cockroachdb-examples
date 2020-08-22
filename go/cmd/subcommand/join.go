package subcommand

import (
	"fmt"

	"github.com/KanchiShimono/cockroachdb-examples/go/models"
	"github.com/jinzhu/gorm"
)

type PersonWithPost struct {
	models.Person
	models.Post
}

func Join(db *gorm.DB) error {
	results := []*PersonWithPost{}
	if err := db.Table("person").Select("person.id, person.name, post.id, post.person_id, post.text").Joins("inner join post on person.id = post.person_id").Scan(&results).Error; err != nil {
		return err
	}

	for _, r := range results {
		fmt.Println(r)
	}

	return nil
}
