package main

import (
	"github.com/garagator3000/gopass/internal/commands"
	"github.com/urfave/cli/v2"
)

const (
	copyright = `
The MIT License (MIT)

© 2024 Sukhodubenko Aleksandr. Novosibirsk.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
)

var (
	version = "dev"
)

func initApp() *cli.App {
	return &cli.App{
		Name:    "gopass",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Sukhodubenko Aleksandr. Novosibirsk",
				Email: "sanya.suhoy99@gmail.com",
			},
		},
		Copyright:              copyright,
		Usage:                  "Утилита для управления паролями или любыми другими текстовыми данными.",
		UsageText:              "gopass [global options] [command] [options]",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "key",
				Aliases:  []string{"k"},
				EnvVars:  []string{"GOPASS_KEY"},
				Usage:    "Ключ, при помощи которого (де)шифруются данные.",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "storage-type",
				Aliases:     []string{"st"},
				EnvVars:     []string{"GOPASS_STORAGE_TYPE"},
				Usage:       "Тип хранилища секретов.",
				Required:    false,
				DefaultText: "sqlite",
			},
			&cli.StringFlag{
				Name:        "storage-path",
				Aliases:     []string{"sp"},
				EnvVars:     []string{"GOPASS_STORAGE_PATH"},
				Usage:       "Путь к хранилищу секретов",
				Required:    false,
				DefaultText: "",
			},
		},

		Commands: []*cli.Command{
			commands.Store,
			commands.Read,
		},
	}
}
