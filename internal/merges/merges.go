package merges

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/github"
)

// Info struct
type Info struct {
	Title     string
	Number    int
	CreatedAt time.Time
	MergedAt  time.Time
}

// Merger is a module merge
type Merger struct {
	Client  *github.Client
	PerPage int
	Page    int
}

// Option function
type Option func(*Merger)

// WithClient sets a github client
func WithClient(c *github.Client) Option {
	return func(m *Merger) {
		m.Client = c
	}
}

// WithPage sets a page number
func WithPage(p int) Option {
	return func(m *Merger) {
		m.Page = p
	}
}

// WithPerPage sets a number of records per page
func WithPerPage(p int) Option {
	return func(m *Merger) {
		m.Page = p
	}
}

// New creates a new merger instance
func New(options ...Option) *Merger {
	m := &Merger{
		Client:  github.NewClient(nil),
		Page:    1,
		PerPage: 15,
	}

	for _, o := range options {
		o(m)
	}

	return m
}

// GetTimes of each merged pr
func (m *Merger) GetTimes(owner, repo string) ([]Info, error) {
	log.WithField("owner", owner).WithField("repo", repo).Debug("getting pull requests")
	pulls, _, err := m.Client.PullRequests.List(context.Background(), owner, repo, &github.PullRequestListOptions{
		State:     "closed",
		Direction: "desc",
		ListOptions: github.ListOptions{
			Page:    m.Page,
			PerPage: m.PerPage,
		},
	})
	if err != nil {
		return nil, err
	}

	var pp []Info
	for _, pull := range pulls {
		pp = append(pp, Info{
			Title:     pull.GetTitle(),
			Number:    pull.GetNumber(),
			CreatedAt: pull.GetCreatedAt(),
			MergedAt:  pull.GetMergedAt(),
		})
	}

	return pp, nil
}
