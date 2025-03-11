package commands

import (
	"fmt"

	"github.com/garagator3000/gopass/internal/cipher"

	"github.com/garagator3000/gopass/internal/storage"
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:      "list",
		Usage:     "List secrets of group.",
		Action:    list,
		Args:      true,
		ArgsUsage: "<groupname>",
	}
}

func list(ctx *cli.Context) error {
	key := ctx.String("key")
	groupname := ctx.Args().Get(0)

	storageType := ctx.String("storage-type")
	storagePath := ctx.String("storage-path")

	storage := storage.Init(storageType, storagePath)
	defer storage.Close()

	encryptedData, err := storage.ListSecret(ctx.Context, groupname)
	if err != nil {
		return err
	}

	fmt.Printf("%s:\n", groupname)
	for _, encrypted := range encryptedData {
		data := cipher.Decrypt(key, encrypted.Data)
		if data == "" {
			continue
		}
		fmt.Printf("  %s: %s\n", encrypted.Name, data)
	}

	return nil
}
