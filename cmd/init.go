package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var storeFile string

var initCmd = &cobra.Command{
	Use:   "init <master-password>",
	Short: "initialize a new password manager",
	Long:  "initialize a new password manager, providing the master password and optionally the path for the store file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		masterPassword := args[0]
		var err error
		if storeFile == "" {
			storeFile, err = os.UserHomeDir()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		}
		if !strings.HasPrefix(storeFile, ".") {
			storeFile = "." + storeFile // make it a hidden file
		}

		f, err := os.Create(storeFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		_, err = fmt.Fprintln(f, fmt.Sprintf("master-password: %s", masterPassword))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("store file created")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&storeFile, "store", "s", "", "Path to the store file (optional)")
}
