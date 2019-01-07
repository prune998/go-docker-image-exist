package main

import (
	"fmt"

	"github.com/prune998/docker-registry-client/registry"
	"github.com/sirupsen/logrus"

	"os"

	"github.com/namsral/flag"
)

var (
	// version is filled by -ldflags  at compile time
	version     = "no version set"
	registryURL = flag.String("registryURL", "https://us.gcr.io", "The Docker Registry URL")
	project     = flag.String("project", "", "The Docker project, if using gcloud registry")
	username    = flag.String("username", "", "The Docker Registry user name, use '_token' if using a a gcloud generated token")
	password    = flag.String("password", "", "The Docker Registry password. use 'gcloud auth print-access-token' if connecting to gcloud")
	logLevel    = flag.String("logLevel", "warn", "log level from debug, info, warning, error")
	image       = flag.String("image", "", "full image name to check (ex: prune998/go-docker-image-exist)")
	tag         = flag.String("tag", "latest", "tag of the image to check (default to 'latest')")
	log         = logrus.New()
)

func main() {
	flag.Parse()

	log.Out = os.Stdout
	log.Formatter = new(logrus.JSONFormatter)
	log.Level, _ = logrus.ParseLevel(*logLevel)

	imageName := *image
	if *project != "" {
		// we are using gcloud, prepend project in image name
		imageName = fmt.Sprintf("%s/%s", *project, *image)
	}

	log.Debugf("searching image %s", imageName)

	// connect to the Docker Registry
	hub, err := registry.New(*registryURL, *username, *password, log.Debugf)
	if err != nil {
		log.Fatalf("error connecting to hub, %v", err)
	}

	tags, err := hub.Tags(imageName)
	if err != nil {
		log.Errorf("Listing image tags error, %v", err)
		os.Exit(1)
	}
	for _, value := range tags {
		if value == *tag {
			log.Debugf("image %s found", imageName)
			os.Exit(0)
		}
	}
	log.Debugf("image %s not found", imageName)
	os.Exit(1)
}
