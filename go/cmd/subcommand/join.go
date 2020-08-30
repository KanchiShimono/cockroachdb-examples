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

const selectStatement = `
person.id,
person.name,
post.id,
person.create_at,
person.update_at,
post.person_id,
post.text,
post.create_at,
post.update_at
`

const joinStatement = `
inner join post on person.id = post.person_id
`

func Join(db *gorm.DB) error {
	results := []*PersonWithPost{}
	if err := db.Table("person").Select(selectStatement).Joins(joinStatement).Scan(&results).Error; err != nil {
		return err
	}

	for _, r := range results {
		fmt.Println(r)
	}

	return nil
}
