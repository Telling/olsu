package main

// ReleaseClient is the interface needed to create a release
type ReleaseClient interface {
	createRelease() (*Release, error)
	deleteRelease() (*Release, error)
	uploadAssets() (*Release, error)
	doesTagExist() (bool, *Release, error)
}

// Release represents a release in general
type Release struct {
	ReleaseName    string
	ReleaseVersion string
	ReleaseText    string
	Draft          bool
	Prerelease     bool
	Assets         []string
	ID             int
}
