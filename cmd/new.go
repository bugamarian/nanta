package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bugamarian/nanta/new"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Provide a title to create a new file",
	Run: func(cmd *cobra.Command, args []string) {
		new.CreateNote()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
