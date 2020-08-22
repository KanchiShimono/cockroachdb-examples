package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/KanchiShimono/cockroachdb-examples/go/cmd/subcommand"
	"github.com/KanchiShimono/cockroachdb-examples/go/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/openpgp/errors"
)

func main() {
	env, err := env.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", env.DBURL)
	if err != nil {
		log.Fatal(fmt.Errorf("gorm.Open error %s", err))
	}
	defer db.Close()

	c := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "select",
				Subcommands: []*cli.Command{
					{
						Name:        "person",
						Description: "select all records of person",
						Action: func(c *cli.Context) error {
							return subcommand.SelectPerson(db)
						},
					},
					{
						Name:        "post",
						Description: "select all records of post",
						Action: func(c *cli.Context) error {
							return subcommand.SelectPost(db)
						},
					},
					{
						Name:        "person-with-post",
						Description: "select person with related post",
						ArgsUsage:   "num",
						Action: func(c *cli.Context) error {
							num, err := strconv.Atoi(c.Args().First())
							if err != nil {
								return errors.InvalidArgumentError("parameter num is required as int")
							}
							return subcommand.SelectPersonWithPost(db, num)
						},
					},
				},
			},
			{
				Name:        "join",
				Description: "select all person records joined with post",
				Action: func(c *cli.Context) error {
					return subcommand.Join(db)
				},
			},
			{
				Name:        "delete",
				Description: "delete all records of person and post",
				Action: func(c *cli.Context) error {
					return subcommand.Delete(db)
				},
			},
			{
				Name:      "insert",
				ArgsUsage: "num-person num-post-per-person",
				Action: func(c *cli.Context) error {
					numPerson, err := strconv.Atoi(c.Args().Get(0))
					if err != nil {
						return errors.InvalidArgumentError("parameter num-person is required as int")
					}
					numPostPerPerson, err := strconv.Atoi(c.Args().Get(1))
					if err != nil {
						return errors.InvalidArgumentError("parameter num-post-per-person is required as int")
					}
					return subcommand.Insert(db, numPerson, numPostPerPerson)
				},
			},
		},
	}

	if err := c.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
