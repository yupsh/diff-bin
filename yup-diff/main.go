package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/diff"
)

const (
	flagUnified          = "unified"
	flagContext          = "context"
	flagBrief            = "brief"
	flagIgnoreCase       = "ignore-case"
	flagIgnoreWhitespace = "ignore-all-space"
	flagSideBySide       = "side-by-side"
	flagRecursive        = "recursive"
)

func main() {
	app := &cli.App{
		Name:  "diff",
		Usage: "compare files line by line",
		UsageText: `diff [OPTIONS] FILE1 FILE2

   Compare FILES line by line.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagUnified,
				Aliases: []string{"u"},
				Usage:   "output NUM (default 3) lines of unified context",
			},
			&cli.IntFlag{
				Name:    flagContext,
				Aliases: []string{"c", "C"},
				Usage:   "output NUM (default 3) lines of copied context",
			},
			&cli.BoolFlag{
				Name:    flagBrief,
				Aliases: []string{"q"},
				Usage:   "report only when files differ",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreCase,
				Aliases: []string{"i"},
				Usage:   "ignore case differences in file contents",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreWhitespace,
				Aliases: []string{"w"},
				Usage:   "ignore all white space",
			},
			&cli.BoolFlag{
				Name:    flagSideBySide,
				Aliases: []string{"y"},
				Usage:   "output in two columns",
			},
			&cli.BoolFlag{
				Name:    flagRecursive,
				Aliases: []string{"r"},
				Usage:   "recursively compare any subdirectories found",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "diff: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (requires exactly 2 files)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagUnified) {
		params = append(params, Unified, UnifiedContext(c.Int(flagUnified)))
	}
	if c.IsSet(flagContext) {
		params = append(params, ContextDiff, ContextLines(c.Int(flagContext)))
	}
	if c.Bool(flagBrief) {
		params = append(params, Brief)
	}
	if c.Bool(flagIgnoreCase) {
		params = append(params, IgnoreCase)
	}
	if c.Bool(flagIgnoreWhitespace) {
		params = append(params, IgnoreWhitespace)
	}
	if c.Bool(flagSideBySide) {
		params = append(params, SideBySide)
	}
	if c.Bool(flagRecursive) {
		params = append(params, Recursive)
	}

	// Create and execute the diff command
	cmd := Diff(params...)
	return gloo.Run(cmd)
}
