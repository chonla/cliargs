# CLI-Args

Simple Command Line Interface Argument for Go.

## Examples

```go
func main() {
    args, err := cliargs.Parse()

    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args.Command {
	case "task":
		switch args.Subcommand {
		case "ls":
			h.List()
		case "add":
		    h.Add(args.Arguments[0])
		}
	case "version":
		printVersion()
	default:
		printVersion()
		fmt.Println("")
		printUsage()
	}
}
```

## License

[MIT](LICENSE)