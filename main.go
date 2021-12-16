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
		Action:          appfunc,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func appfunc(c *cli.Context) error {
	args := c.Args().Slice()

	if len(args) < 1 {
		return errors.New("at least one path is required")
	}

	num := 0

	for i, v := range args {
		_, err := os.Stat(v)
		if err != nil {
			return err
		}

		if i == 0 {
			p := filepath.Join(
				filepath.Dir(v),
				"*"+filepath.Ext(v))
			fl, err := filepath.Glob(p)

			if err != nil {
				return err
			}

			num = len(fl)
		}

		err = os.Rename(
			v,
			filepath.Join(
				filepath.Dir(v),
				fmt.Sprintf("%04d", num)+"_"+filepath.Base(v)))
		if err != nil {
			return err
		}
	}

	return nil
}
