package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of available templates",
	Long: `There are only certain memes that can be used and this command will print 
out an alphabetized list of the options.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		l := getList()

		printList(l)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

type Template struct {
	Name      string
	Shortname string
	URL       string
}

func getList() []Template {
	var list map[string]string
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://memegen.link/api/templates/", nil)
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	b, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(b, &list); err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}

	l := make([]Template, len(list), len(list))

	for name, url := range list {
		splitPath := strings.Split(url, "/")
		path := splitPath[len(splitPath)-1]
		l = append(l, Template{Name: name, Shortname: path, URL: url})
	}

	return l
}

func printList(l []Template) {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Shortname < l[j].Shortname
	})
	for i, item := range l {
		fmt.Printf("%d. %s: %s\n", i+1, item.Shortname, item.Name)
	}
}
