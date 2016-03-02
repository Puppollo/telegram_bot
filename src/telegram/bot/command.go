package bot

import (
	"fmt"
	"path/filepath"
)

const (
	RUN   = "run"
	GETME = "getme"
	HELP  = "help"
	LIST  = "list"
)

var (
	commands = map[string]command{
		GETME: getMe,
		HELP:  help,
		RUN:   run,
		LIST:  list,
	}
)

func getMe(b *Bot, cmd ...string) string {
	raw, err := b.request(METHOD_GETME, "")
	if err != nil {
		return err.Error()
	}
	return string(raw[0:])
}

// help output
func help(b *Bot, cmd ...string) string {
	return fmt.Sprintf("supported commands:\n%s - <shell> <arguments>\n%s - help\n%s - bot info", RUN, HELP, GETME)
}

// run shell script located in commands dir
func run(b *Bot, cmd ...string) string {
	output, err := b.run(cmd...)
	if err != nil {
		return err.Error()
	}
	return output
}

// list of commands in commands dir
func list(b *Bot, cmd ...string) string {
	files, err := filepath.Glob(b.Commands + "/*")
	if err != nil {
		return err.Error()
	}
	var res string
	for _, file := range files {
		if file == "." || file == ".." {
			continue
		}
		res += file + "\n"
	}
	return res
}
