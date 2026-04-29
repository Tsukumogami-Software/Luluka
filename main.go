package main

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Luluka",
	Short: "View and tweak shaders made with Kage for Ebiten",
	Long: `Luluka helps you display Ebiten shaders in Kage.
	Simply load a shader by indicating its path:
		luluka shader.kage
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Run(args[0])
	},
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
