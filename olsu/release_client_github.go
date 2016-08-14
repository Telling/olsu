package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"path"
)

// GithubRelease implements the ReleaseClient interface
type GithubRelease struct {
	ReleaseClient
	client  github.Client
	release *Release
}

// NewGithubReleaseClient creates a new release client for github
func NewGithubReleaseClient(rel *Release) *GithubRelease {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: args.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	var client github.Client
	client = *github.NewClient(tc)

	return &GithubRelease{client: client, release: rel}
}

// createRelease creates a new release on github, if theres any extra assets
// it will upload these to the release on github.
func (g *GithubRelease) createRelease() (*Release, error) {
	_githubRelease := github.RepositoryRelease{
		Name:       &g.release.ReleaseName,
		TagName:    &g.release.ReleaseVersion,
		Body:       &g.release.ReleaseText,
		Draft:      &g.release.Draft,
		Prerelease: &g.release.Prerelease,
	}

	githubRelease, _, err := g.client.Repositories.CreateRelease(
		args.Owner,
		args.Repository,
		&_githubRelease,
	)
	if err != nil {
		return g.release, err
	}

	// Set release id
	g.release.ID = *githubRelease.ID

	if len(args.Assets) > 0 {
		_, err := g.uploadAssets()
		if err != nil {
			return g.release, err
		}
		return g.release, nil
	}

	return g.release, nil
}

// deleteRelease deletes a release on github.
func (g *GithubRelease) deleteRelease() (*Release, error) {
	githubRelease, _, err := g.client.Repositories.GetReleaseByTag(
		args.Owner,
		args.Repository,
		g.release.ReleaseVersion,
	)

	if err != nil {
		return g.release, err
	}

	_, err = g.client.Repositories.DeleteRelease(
		args.Owner,
		args.Repository,
		*githubRelease.ID,
	)
	if err != nil {
		return g.release, err
	}

	return g.release, nil
}

// uploadAssets uploads any extra files to the release on github.
func (g *GithubRelease) uploadAssets() (*Release, error) {
	var assets []string

	for _, asset := range g.release.Assets {
		if _, err := os.Stat(asset); os.IsNotExist(err) {
			return g.release, err
		}
		file, err := os.Open(asset)
		if err != nil {
			return g.release, err
		}
		uploadOptions := &github.UploadOptions{
			Name: path.Base(file.Name()),
		}
		releaseAsset, _, err := g.client.Repositories.UploadReleaseAsset(
			args.Owner,
			args.Repository,
			g.release.ID,
			uploadOptions,
			file,
		)
		if err != nil {
			g.release.Assets = assets
			return g.release, err
		}
		assets = append(assets, *releaseAsset.Name)
	}

	g.release.Assets = assets

	return g.release, nil
}

// doesTagExist checks if a given release (tag) exists.
func (g *GithubRelease) doesTagExist() (bool, *Release, error) {
	checkRelease, response, err := g.client.Repositories.GetReleaseByTag(
		args.Owner,
		args.Repository,
		g.release.ReleaseVersion,
	)
	if response == nil {
		return false, g.release, fmt.Errorf("No response from github")
	}
	if err != nil {
		return false, g.release, err
	}
	if checkRelease != nil {
		return true, g.release, nil
	}

	return false, g.release, nil
}
