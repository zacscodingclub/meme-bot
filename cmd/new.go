package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Copies a new URL for a custom meme to clipboard",
	Long: `How to use it:
$ meme-bot new kermit "some top words" "some bottom words"

Now the string "https://memegen.link/kermit/some_top_words/some_bottom_words.jpg" will be 
copied to your clipboard.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		template := url.PathEscape(args[0])
		top := url.PathEscape(underlinize(args[1]))
		bottom := url.PathEscape(underlinize(args[2]))

		u := fmt.Sprintf("https://memegen.link/%s/%s/%s.jpg", template, top, bottom)
		fmt.Printf("URL (%s) copied to your clipboard.\n", u)
		clipboard.WriteAll(u)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func underlinize(str string) string {
	return strings.Replace(str, " ", "_", -1)
}
