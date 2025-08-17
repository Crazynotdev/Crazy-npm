go
package cmd
import {
  "fmt"
  "os"
  "github.com/spf13/cobra"
  }
var rootCmd = &cobra.command{
  Use: "crazy",
  Short: "crazy is fast and crazy package manager",
  Long: "crazy CLI tools to manage custom package from central repository",
  }
func Execute() {
  if err := rootCmd.Execute();
  err != nil {
    fmt.Println(err)
    os.Exit(1)
    }
  }
