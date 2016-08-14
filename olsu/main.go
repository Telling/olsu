package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	args OlsuArgs
)

// printReleaseInfo pretty prints release info
func printReleaseInfo(state string, rel *Release) {
	fmt.Println(fmt.Sprintf("%v release of %v (%v/%v)", state, args.Repository, args.Owner, args.Repository))
	fmt.Println(fmt.Sprintf(" - name: %v (ID: %v)", rel.ReleaseName, rel.ID))
	fmt.Println(" - tag:", rel.ReleaseVersion)
	fmt.Println(" - description:", rel.ReleaseText)
}

func main() {
	args, err := parseArgsAndEnvs()
	// TODO: fix these err != nill && !args.Quiet, is not correct
	if err != nil && !args.Quiet {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	var release *Release
	release = &Release{
		ReleaseName:    args.ReleaseName,
		ReleaseVersion: args.ReleaseVersion,
		ReleaseText:    args.ReleaseText,
		Draft:          args.Draft,
		Prerelease:     args.Prerelease,
		Assets:         args.Assets,
	}

	// how do I make one type be both?????
	var client ReleaseClient

	if args.Backend == "github" {
		client = NewGithubReleaseClient(release)
	}

	if args.Backend == "bitbucket" {
		client = NewBitbucketReleaseClient(release)
	}

	_, rel, err := client.doesTagExist()
	if err != nil {
		printReleaseInfo("Existing", rel)
		os.Exit(1)
	}

	var deleteErr error
	var deletedRelease *Release

	release, err = client.createRelease()
	if err != nil {
		if os.Getenv("OLSU_DELETE_RELEASE") != "" {
			deletedRelease, deleteErr = client.deleteRelease()
		}
	}

	if !args.Quiet && err != nil {
		fmt.Println("Error creating release:")
		fmt.Println(err)
		if deleteErr != nil {
			fmt.Println("Error deleting release:")
			fmt.Println(deleteErr)
		} else {
			printReleaseInfo("Deleted", deletedRelease)
		}
		os.Exit(1)
	}

	if args.Quiet {
		fmt.Printf("%v", release.ID)
	}

	if !args.Quiet {
		printReleaseInfo("Created", release)
		if len(release.Assets) > 0 {
			fmt.Println("Uploaded assets:")
			for _, asset := range release.Assets {
				fmt.Println(" -", asset)
			}
		}
	}

	os.Exit(0)
}
