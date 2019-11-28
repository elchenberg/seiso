package commands

import (
	"fmt"

	"github.com/heroku/docker-registry-client/registry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var image string

func init() {
	rootCmd.AddCommand(tagCmd)
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Print the available tags",
	Long:  `tbd`,
	Args:  cobra.ExactValidArgs(1),
	Run:   printImageStreamTags,
}

func printImageStreamTags(cmd *cobra.Command, args []string) {
	image := args[0]
	url := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	if err != nil {
		log.WithError(err).WithField("url", url).Fatal("Registry is currently unavailable.")
	}
	tags, err := hub.Tags(image)

	if err != nil {
		log.WithError(err).WithField("url", url).Fatal("Could not list image tags.")
	}

	for _, tag := range tags {
		fmt.Println(tag)
	}
}
