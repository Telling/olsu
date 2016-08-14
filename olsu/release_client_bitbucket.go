package main

import (
	"fmt"
	"net/http"
)

// BitbucketRelease implements the ReleaseClient interface
type BitbucketRelease struct {
	ReleaseClient
	client  http.Client
	release *Release
}

// NewBitbucketReleaseClient creates a new release client for bitbucket
func NewBitbucketReleaseClient(rel *Release) *BitbucketRelease {
	return &BitbucketRelease{}
}

// createRelease creates a new release on bitbucket, if theres any extra assets
// it will upload these to the release on bitbucket.
func (b *BitbucketRelease) createRelease() (*Release, error) {
	return b.release, fmt.Errorf("Not implemented.")
}

// deleteRelease deletes a release on github.
func (b *BitbucketRelease) deleteRelease() (*Release, error) {
	return b.release, fmt.Errorf("Not implemented.")
}

// uploadAssets uploads any extra files to the release on bitbucket.
func (b *BitbucketRelease) uploadAssets() (*Release, error) {
	return b.release, fmt.Errorf("Not implemented.")
}

// doesTagExist checks if a given release (tag) exists.
func (b *BitbucketRelease) doesTagExist() (bool, *Release, error) {
	return false, b.release, fmt.Errorf("Not implemented.")
}
