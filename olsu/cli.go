package main

import (
	"flag"
	"fmt"
	"os"
)

// OlsuArgs keeps all the info
// TODO: use an OlsuRelease to keep release information instead
type OlsuArgs struct {
	Owner          string
	Repository     string
	ReleaseText    string
	ReleaseName    string
	ReleaseVersion string
	Assets         []string
	Draft          bool
	Prerelease     bool
	Token          string
	Quiet          bool
	DeleteRelease  bool
	Backend        string
}

// parseArgsAndEnvs parses commandline arguments, flags and environment variables
// It returns an instance of OlsuArgs
func parseArgsAndEnvs() (OlsuArgs, error) {
	// TODO: add docstring
	flag.StringVar(&args.Owner, "o", "", "specify the repository owner.")
	flag.StringVar(&args.Repository, "r", "", "specify a repository.")
	flag.BoolVar(&args.Draft, "d", false, "if it's a draft.")
	flag.BoolVar(&args.Prerelease, "p", false, "if it's a prerelease.")
	flag.StringVar(&args.Token, "t", "", "github token")
	flag.BoolVar(&args.Quiet, "q", false, "Only output release id.")
	flag.StringVar(&args.Backend, "b", "github", "backend")

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
		return args, fmt.Errorf("name, version and description are required arguments")
	}

	args.ReleaseName = inputArgs[0]
	args.ReleaseVersion = inputArgs[1]
	args.ReleaseText = inputArgs[2]

	if len(inputArgs) > numRequiredArgs {
		args.Assets = inputArgs[3:]
	}

	// TODO: make function for this
	// If env vars are set, they overrule
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
		return args, fmt.Errorf("-t or env. variable OLSU_TOKEN is required")
	}
	if args.Owner == "" && os.Getenv("OLSU_OWNER") == "" {
		return args, fmt.Errorf("-o or env. variable OLSU_OWNER is required")
	}
	if args.Repository == "" && os.Getenv("OLSU_REPOSITORY") == "" {
		return args, fmt.Errorf("-o or env. variable OLSU_REPOSITORY is required")
	}

	return args, nil
}
