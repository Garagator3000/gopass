package commands

import (
	"fmt"

	"github.com/garagator3000/gopass/internal/cipher"
	"github.com/garagator3000/gopass/internal/storage"
	"github.com/urfave/cli/v2"
)

var Read = &cli.Command{
	Name:   "read",
	Usage:  "Прочитать секрет",
	Action: read,

	Flags: []cli.Flag{
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

func read(ctx *cli.Context) error {
	key := ctx.String("key")
	name := ctx.String("name")

	storageType := ctx.String("storage-type")
	storagePath := ctx.String("storage-path")

	storage := storage.StorageInit(storageType, storagePath)
	defer storage.Close()

	encryptedData, err := storage.ReadSecret(ctx.Context, name)
	if err != nil {
		return err
	}

	data := cipher.Decrypt(key, encryptedData)

	fmt.Println(data)

	return nil
}
