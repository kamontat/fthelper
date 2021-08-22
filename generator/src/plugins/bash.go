package plugins

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kamontat/fthelper/generator/v4/src/clusters"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Bash(data maps.Mapper, config maps.Mapper) runners.Runner {
	return clusters.NewRunner(data, config, func(p *clusters.ExecutorParameter) error {
		var withSudo = p.Data.Bo("withSudo", false)

		var sudoPath = ""
		if withSudo {
			p, err := exec.LookPath("sudo")
			if err != nil {
				return err
			}

			sudoPath = p
		}

		var cmd *exec.Cmd = nil
		if p.Data.Has("commands") {
			var raw = p.Data.Ai("commands")
			var commands []string = make([]string, 0)
			if withSudo && sudoPath != "" {
				commands = append(commands, sudoPath)
			}

			for i, c := range raw {
				if i == 0 {
					var cmd, err = exec.LookPath(c.(string))
					if err != nil {
						return err
					}

					commands = append(commands, cmd)
				} else {
					commands = append(commands, c.(string))
				}
			}

			cmd = &exec.Cmd{
				Path: commands[0],
				Args: commands,

				Stdin:  os.Stdin,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
		} else {
			var args []string = make([]string, 0)
			if withSudo && sudoPath != "" {
				args = append(args, sudoPath)
			}

			var bashExecutePath = ""
			if bash, err := exec.LookPath("bash"); err == nil {
				bashExecutePath = bash
			}
			if sh, err := exec.LookPath("sh"); bashExecutePath == "" && err == nil {
				bashExecutePath = sh
			}
			if bashExecutePath == "" {
				return fmt.Errorf("cannot found bash execute (search for 'bash' and 'sh')")
			}
			args = append(args, bashExecutePath)

			file, err := fs.Build(fs.ToObject(p.Data.Zi("file"), p.Config), p.VarConfig)
			if err != nil {
				return err
			}

			args = append(args, file.Single().Abs())
			cmd = &exec.Cmd{
				Path: args[0],
				Args: args,

				Stdin:  os.Stdin,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
		}

		p.Logger.Debug("execute: %v", cmd)
		return cmd.Run()
	}, &clusters.Settings{
		DefaultWithCluster: false,
	})
}
