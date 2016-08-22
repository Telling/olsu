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
	if err != nil {
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

	var backendClients = [...]string{
		"github",
		"bitbucket",
	}

	// Client used to create release
	var client ReleaseClient

	switch {
	// github
	case args.Backend == backendClients[0]:
		client = NewGithubReleaseClient(release)
	// bitbucket
	case args.Backend == backendClients[1]:
		client = NewBitbucketReleaseClient(release)
	default:
		fmt.Println(fmt.Sprintf("Unknown backend: %v", args.Backend))
		fmt.Println(fmt.Sprintf("Supported backends: %v", backendClients))
		os.Exit(1)
	}

	tag, release, err := client.doesTagExist()
	if err != nil {
		fmt.Println("Error checking tag:")
		fmt.Println(err)
		os.Exit(1)
	}
	if tag && !args.DeleteRelease {
		printReleaseInfo("Existing", release)
		os.Exit(1)
	} else if tag && args.DeleteRelease {
		// if there's an existing tag its deleted before creating a new one
		release, err = client.deleteRelease()
		if err != nil {
			fmt.Println("Error deleting release:")
			fmt.Println(err)
			os.Exit(1)
		}
		if !args.Quiet {
			printReleaseInfo("Deleted", release)
		}
	}

	release, err = client.createRelease()
	if release == nil && err != nil {
		fmt.Println("Error creating release:")
		fmt.Println(err)
		os.Exit(1)
	}
	// Release created, but error uploading assets.
	if release != nil && err != nil {
		fmt.Println("Error uploading assets:")
		fmt.Println(release.Assets)
		fmt.Println(err)
		printReleaseInfo("Created", release)
		os.Exit(1)
	}

	if args.Quiet {
		fmt.Printf("%v", release.ID)
		os.Exit(0)
	}

	printReleaseInfo("Created", release)
	if len(release.Assets) > 0 {
		fmt.Println("Uploaded assets:")
		for _, asset := range release.Assets {
			fmt.Println(" -", asset)
		}
	}

	os.Exit(0)
}
