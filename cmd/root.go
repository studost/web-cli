package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
    // "strings"
)

var rootCmd = &cobra.Command{
    Use: "web-cli",
    Short: "Read Atom feeds",
    Long: `Just a small CLI application.
            Read Atom feeds`,
}

var cmdLs = &cobra.Command{
    Use: "ls",
    Short: "List articles",
    Long: `Show the first five articles from heise.de`,
    Run: func(cmd *cobra.Command, args []string) {
            list()
    },
}


var cmdDescribe = &cobra.Command{
    Use: "describe [id]",
    Short: "Show details for an article",
    Long: `Details`,
    Args: cobra.MinimumArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
            describe(args[0])
    },
}
func Exec() {
    rootCmd.AddCommand(cmdLs)
    rootCmd.AddCommand(cmdDescribe)
    err := rootCmd.Execute()
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}

