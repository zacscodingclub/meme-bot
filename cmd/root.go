package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meme-bot",
	Short: "Generates link to your customized meme.",
	Long: `Creates a customized meme using one of many templates and copies the URL 
to your clipboard.  For example:
 $ ./meme-bot new icanhas "i can haz" "customized meme"
Copies this URL string to your clipboard so you can use it later: 
  https://memegen.link/icanhas/i_can_haz/customized_meme.jpg
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
