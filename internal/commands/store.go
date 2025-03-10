package commands

import (
	"fmt"
	osuser "os/user"
	"time"

	"github.com/garagator3000/gopass/internal/cipher"
	"github.com/garagator3000/gopass/internal/entities"
	"github.com/garagator3000/gopass/internal/storage"
	"github.com/urfave/cli/v2"
)

func Store() *cli.Command {
	return &cli.Command{
		Name:   "store",
		Usage:  "Store the secret.",
		Action: store,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "data",
				Aliases:  []string{"d"},
				Usage:    "The data to be stored.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "The name for the data to be saved.",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Usage:       "Username on behalf of which the data is stored.",
				Required:    false,
				DefaultText: "",
			},
			&cli.StringFlag{
				Name:        "group",
				Aliases:     []string{"g"},
				Usage:       "Group of secrets wich the data will be stored.",
				Required:    false,
				DefaultText: "",
			},
		},
	}
}

func store(ctx *cli.Context) error {
	secret := formSecret(ctx)

	storageType := ctx.String("storage-type")
	storagePath := ctx.String("storage-path")

	storage := storage.Init(storageType, storagePath)
	defer storage.Close()

	if err := storage.CreateSecret(ctx.Context, secret); err != nil {
		return err
	}

	return nil
}

func formSecret(ctx *cli.Context) entities.Secret {
	key := ctx.String("key")

	name := ctx.String("name")
	data := ctx.String("data")
	group := ctx.String("group")

	encodedSecret := cipher.Encrypt(key, data)

	user := ctx.String("user")
	if user == "" {
		if currentUser, err := osuser.Current(); err == nil {
			user = currentUser.Name
		} else {
			fmt.Println("Can't get current user. Use \"\"")
		}
	}

	return entities.Secret{
		Name:      name,
		Data:      encodedSecret,
		User:      user,
		Group:     group,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
