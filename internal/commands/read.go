package commands

import (
	"fmt"

	"github.com/garagator3000/gopass/internal/cipher"
	"github.com/garagator3000/gopass/internal/storage"
	"github.com/urfave/cli/v2"
)

func Read() *cli.Command {
	return &cli.Command{
		Name:   "read",
		Usage:  "Read the secret.",
		Action: read,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "The name of the secret to be requested.",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Usage:       "Username on behalf of which the data is stored.",
				Required:    false,
				DefaultText: "",
			},
		},
	}
}

func read(ctx *cli.Context) error {
	key := ctx.String("key")
	name := ctx.String("name")

	storageType := ctx.String("storage-type")
	storagePath := ctx.String("storage-path")

	storage := storage.Init(storageType, storagePath)
	defer storage.Close()

	encryptedData, err := storage.ReadSecret(ctx.Context, name)
	if err != nil {
		return err
	}

	data := cipher.Decrypt(key, encryptedData)

	fmt.Println(data)

	return nil
}
