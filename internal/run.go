package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/directory"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/signature"
)

// Run is the definition of a single "docker pull load test".
type Run struct {
	Registry       string   `json:"registry"`
	Images         []string `json:"images"`
	TotalPulls     int      `json:"totalPulls"`
	PullsPerSecond int      `json:"pullsPerSecond"`
}

func (r Run) Exec() error {
	// Fake policy just to make this work
	p := &signature.Policy{
		Default: []signature.PolicyRequirement{
			signature.NewPRInsecureAcceptAnything(),
		},
	}

	pc, err := signature.NewPolicyContext(p)
	if err != nil {
		return err
	}

	// Pull each image sequentially
	for _, t := range r.Images {
		td, err := ioutil.TempDir("", "puller-")
		if err != nil {
			return err
		}

		destRef, err := directory.NewReference(td)
		if err != nil {
			return err
		}

		srcRef, err := docker.ParseReference(fmt.Sprintf("//%s/%s", r.Registry, t))
		if err != nil {
			return err
		}

		manifest, err := copy.Image(context.Background(), pc, destRef, srcRef, nil)
		if err != nil {
			return err
		}

		fmt.Printf("%s", manifest)

		// Clean up stored image files
		if err := os.RemoveAll(td); err != nil {
			return err
		}
	}

	return nil
}
