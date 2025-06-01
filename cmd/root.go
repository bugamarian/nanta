package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "nanta",
		Short: "Not a note taking app",
		Long: `Just a simple CLI tool that will create a timestamped file
	for taking notes. The app supports templates which let's you customize
	the initial state of a document.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Could not read configuration: %s", err)
	}

	configDir = filepath.Join(configDir, "nanta")
	configFile := filepath.Join(configDir, "config.yaml")
	viper.SetConfigFile(configFile)

	viper.SetDefault("note_dir", filepath.Join(configDir, "./notes"))
	viper.SetDefault(
		"template",
		filepath.Join(configDir, "./templates", "default.tmpl"),
	)
	viper.SetDefault("modifier", "nvim")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Could not read config, using defaults")
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// var cfg config.Config
	rootCmd.PersistentFlags().
		StringP("config", "c", "", "config file (default is $HOME/.config/nanta/config.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().
		StringP("title", "t", "<Untitled Note>", "Title of the note")
	viper.BindPFlag("title", rootCmd.PersistentFlags().Lookup("title"))
	cobra.OnInitialize(initConfig)
}
