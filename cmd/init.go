package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/knbr13/mpwd/internal/hash"
)

var storePath string

var initCmd = &cobra.Command{
	Use:   "init <master-password>",
	Short: "initialize a new password manager",
	Long:  "initialize a new password manager, providing the master password and optionally the path for the store file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		masterPassword := args[0]
		var err error
		if storePath == "" {
			storePath, err = os.UserHomeDir()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		}

		fName := filepath.Join(storePath, ".mpwd.yaml")

		f, err := os.Create(fName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		hashedPassword, err := hash.HashPassword(masterPassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error hashing password: %v\n", err)
			os.Exit(1)
		}

		_, err = f.Write(hashedPassword)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing to file: %v\n", err)
			os.Exit(1)
		}

		absPath, err := filepath.Abs(fName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("password manager initialized, store file created %q\n", absPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&storePath, "store", "s", "", "Path to the store file (optional)")
}
