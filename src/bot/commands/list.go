package commands

import (
	"bot"
	"path/filepath"
)

type List struct{}

func (List) Help() string {
	return "list - list of executable (run) commands"
}

func (List) Run(b *bot.Bot, cmd ...string) string {
	files, err := filepath.Glob(b.Executables + "/*")
	if err != nil {
		return err.Error()
	}
	var res string
	for _, file := range files {
		if file == "." || file == ".." {
			continue
		}
		res += filepath.Base(file) + "\n"
	}
	return res
}
