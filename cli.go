package GoCLI

import (
	"os"
	"fmt"

	"github.com/urfave/cli"
)

type Command struct {
	Name   string
	Usage  string
	Action func(flags map[string]string, args ...string) error
	StringFlags   []*StringFlag
}

type StringFlag struct {
	Name    string
	Usage   string
	Default string
}

type BoolFlag struct {
	*cli.BoolFlag
}

type Options struct {
	AppName       string
	AppUsage      string
	Commands      []*Command
	StringFlags   []*StringFlag
	BoolFlags     []*BoolFlag
	DefaultAction func(args ...string) error
}

func StartCLI(options *Options) error {
	return createCLIApp(options).Run(os.Args)
}

func createCLIApp(options *Options) *cli.App {
	app := cli.NewApp()
	app.Name = options.AppName
	app.Usage = options.AppUsage
	app.Commands = createAppCommands(options.Commands)
	app.Flags = createAppFlags(options.StringFlags)
	if options.DefaultAction != nil {
		app.Action = func(c *cli.Context) error {
			return options.DefaultAction()
		}
	}
	return app
}

func createAppCommands(commands []*Command) cli.Commands {
	co := cli.Commands{}
	for _, command := range commands {
		co = append(co, createCommand(command))
	}
	return co
}

func createCommand(command *Command) cli.Command {
	return cli.Command{
		Name:  command.Name,
		Usage: command.Usage,
		Action: func(c *cli.Context) error {
			if err := command.Action(getFlags(c), getArgs(c)...); err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
		Flags: createAppFlags(command.StringFlags),
		HideHelp: true,
	}
}

func getArgs(context *cli.Context) (args []string) {
	cliArgs := context.Args()
	for i := 0; i < context.NArg(); i++ {
		args = append(args, cliArgs.Get(i))
	}
	return args
}

func getFlags(context *cli.Context) map[string]string {
	flags := map[string]string{}
	for _, name := range context.FlagNames() {
		flags[name] = context.String(name)
	}
	return flags
}

func createAppFlags(flags []*StringFlag) []cli.Flag{
	fl := []cli.Flag{}
	for _, flag := range flags {
		fl = append(fl, createFlag(flag))
	}
	return fl
}

func createFlag(flag *StringFlag) cli.Flag {
	return cli.StringFlag{
		Name:  flag.Name,
		Usage: flag.Usage,
		Value: flag.Default,
	}
}
