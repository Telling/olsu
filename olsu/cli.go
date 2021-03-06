package main

import (
	"flag"
	"fmt"
	"os"
)

// OlsuArgs keeps all information regarding arguments and release
type OlsuArgs struct {
	Release       Release
	Owner         string
	Repository    string
	Token         string
	Quiet         bool
	DeleteRelease bool
	Backend       string
}

// parseArgsAndEnvs parses commandline arguments, flags and environment variables
// It returns an instance of OlsuArgs
func parseArgsAndEnvs() error {
	flag.StringVar(&args.Owner, "o", "", "owner of the repository.")
	flag.StringVar(&args.Repository, "r", "", "name of the repository.")
	flag.BoolVar(&args.Release.Draft, "d", false, "if it's a draft.")
	flag.BoolVar(&args.Release.Prerelease, "p", false, "if it's a prerelease.")
	flag.StringVar(&args.Token, "t", "", "access token for github.")
	flag.BoolVar(&args.Quiet, "q", false, "quiet, only output release id.")
	flag.StringVar(&args.Backend, "b", "github", "backend")
	flag.BoolVar(&args.DeleteRelease, "dr", false, "delete release if it exists before creating new.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("   name version description [attachment ... attachment]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	inputArgs := flag.Args()

	// number of required arguments
	numRequiredArgs := 3

	if len(inputArgs) < numRequiredArgs {
		return fmt.Errorf("name, version and description are required arguments")
	}

	args.Release.ReleaseName = inputArgs[0]
	args.Release.ReleaseVersion = inputArgs[1]
	args.Release.ReleaseText = inputArgs[2]

	if len(inputArgs) > numRequiredArgs {
		args.Release.Assets = inputArgs[3:]
	}

	if os.Getenv("OLSU_OWNER") != "" {
		args.Owner = os.Getenv("OLSU_OWNER")
	}
	if os.Getenv("OLSU_REPOSITORY") != "" {
		args.Repository = os.Getenv("OLSU_REPOSITORY")
	}
	if os.Getenv("OLSU_TOKEN") != "" {
		args.Token = os.Getenv("OLSU_TOKEN")
	}
	if os.Getenv("OLSU_DELETE_RELEASE") != "" {
		args.DeleteRelease = true
	}

	if args.Token == "" && os.Getenv("OLSU_TOKEN") == "" {
		return fmt.Errorf("-t or env. variable OLSU_TOKEN is required")
	}
	if args.Owner == "" && os.Getenv("OLSU_OWNER") == "" {
		return fmt.Errorf("-o or env. variable OLSU_OWNER is required")
	}
	if args.Repository == "" && os.Getenv("OLSU_REPOSITORY") == "" {
		return fmt.Errorf("-o or env. variable OLSU_REPOSITORY is required")
	}

	return nil
}
