package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bugamarian/nanta/common"
	"github.com/bugamarian/nanta/edit"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the last created document",
	Run: func(cmd *cobra.Command, args []string) {
		notesDir := viper.GetString("notes_dir")
		modifier := viper.GetString("modifier")
		path, err := edit.FindLastNote(notesDir)
		if err != nil {
			log.Fatalf("Error while trying to edit lates document: %s", err)
		}
		common.OpenFile(modifier, path)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
