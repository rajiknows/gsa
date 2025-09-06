package cli

import (
	"fmt"
	"os"

	"github.com/rajiknows/gsa/internal/engine"
	"github.com/rajiknows/gsa/internal/rules"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gsa",
	Short: "Static analysis for Go",
}

var analyzeCommand = &cobra.Command{
	Use:   "analyze [path]",
	Short: "Analyze Go code",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		files, err := engine.CollectGoFiles(path)
		if err != nil {
			return err
		}

		issues, err := engine.Run(files, []engine.Rule{rules.TodoRule{}, rules.SleepRule{}, rules.ConcurrencyRule{}, rules.UncheckedErrorRule{}})
		if err != nil {
			return err
		}

		for _, i := range issues {
			fmt.Printf("[%s] %s:%d %s\n", i.Rule, i.File, i.Line, i.Message)
		}
		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(analyzeCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
