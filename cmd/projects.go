package cmd

import (
	"fmt"
	"github.com/cheynewallace/tabby"
	"github.com/leanovate/mite-go/mite"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	projectsCommand.AddCommand(listProjectsCommand)
	rootCmd.AddCommand(projectsCommand)
}

var projectsCommand = &cobra.Command{
	Use:   "projects",
	Short: "list & adds projects",
	Run:   listProjectsCommand.Run,
}

var listProjectsCommand = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	Run: func(cmd *cobra.Command, args []string) {
		api := mite.NewMiteApi(deps.conf.GetApiUrl(), deps.conf.GetApiKey())
		projects, err := api.Projects()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		t := tabby.New()
		t.AddHeader("id", "name", "notes")
		for _, project := range projects {
			t.AddLine(project.Id, project.Name, project.Note)
		}
		t.Print()
	},
}