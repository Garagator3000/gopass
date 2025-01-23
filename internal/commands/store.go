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

var Store = &cli.Command{
	Name:   "store",
	Usage:  "Сохранить секрет",
	Action: store,

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "data",
			Aliases:  []string{"d"},
			Usage:    "Данные, которые нужно сохранить",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Название для сохраняемых данных",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "user",
			Aliases:     []string{"u"},
			Usage:       "Имя пользователя от имени, которого хранятся данные.",
			Required:    false,
			DefaultText: "",
		},
	},
}

func store(ctx *cli.Context) error {
	secret := formSecret(ctx)

	storageType := ctx.String("storage-type")
	storagePath := ctx.String("storage-path")

	storage := storage.StorageInit(storageType, storagePath)
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
