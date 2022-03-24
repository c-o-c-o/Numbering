package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "Numbering",
		Usage:           Version,
		Description:     "---",
		Version:         Version,
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "directory",
				Aliases: []string{"d", "dir"},
				Usage:   "基準になるディレクトリを指定します。",
			},
		},
		Action: appfunc,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func appfunc(c *cli.Context) error {
	paths := c.Args().Slice()
	if len(paths) < 1 {
		return errors.New("at least one path is required")
	}

	num, err := getNumber(c.Path("directory"), paths[0])
	if err != nil {
		return err
	}

	for _, path := range paths {
		_, err := os.Stat(path)
		if err != nil {
			return err
		}

		err = os.Rename(
			path,
			filepath.Join(
				filepath.Dir(path),
				fmt.Sprintf("%04d", num)+"_"+filepath.Base(path)))
		if err != nil {
			return err
		}
	}

	return nil
}

func getNumber(basedir string, path string) (int, error) {
	if basedir != "" {
		_, err := os.Stat(basedir)
		if err != nil {
			return 0, err
		}
	}

	topnum := 0

	if basedir == "" {
		basedir = filepath.Dir(path)
	} else {
		topnum = 1
	}

	files, err := filepath.Glob(filepath.Join(basedir, "*"+filepath.Ext(path)))
	if err != nil {
		return 0, err
	}

	return topnum + len(files), err
}
