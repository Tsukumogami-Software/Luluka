package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "luluka sample/shader.kage",
	Short: "View and tweak shaders made with Kage for Ebiten",
	Long: `Luluka helps you display Ebiten shaders written in Kage.
Simply load a shader by indicating its path.
Optionally specify uniforms using -u and images using -i.
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uniformFlags, err := cmd.Flags().GetStringSlice("uniform")
		if err != nil {
			log.Panicf("Failed to read uniform flags: %v", err)
		}

		images, err := cmd.Flags().GetStringSlice("images")
		if err != nil {
			log.Panicf("Failed to read image flags: %v", err)
		}
		if len(images) > 4 {
			log.Panicf("Cannot use more than 4 images")
		}

		Run(args[0], uniformFlags, images)
	},
}

func init() {
	rootCmd.Flags().StringSliceP("uniform", "u", []string{}, "specifies a uniform value (use name.0 name.1 etc for vectors)")
	rootCmd.Flags().StringSliceP("images", "i", []string{}, "passes a png image/texture to the shader (max 4)")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
