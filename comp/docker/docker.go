// Package docker is a ucc computation for building Docker images.
package docker // import "hkjn.me/ucc/comp/docker"

import (
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"

	"hkjn.me/ucc/api"
)

type (
	// Computation builds Docker images.
	Computation struct {
	}
	// Inputs for the build.
	Inputs struct {
		// Ancestor of the image we're building (if any).
		Ancestor string
		// Dockerfile to use for the build.
		Dockerfile string
		// Input files.
		Files []string
		// Version of docker used.
		DockerVersion string
		// CPU architecture, e.g. "amd64"
		Architecture string
		// Os used for build, e.g. "linux".
		Os string
		// Graphdriver, .e.g "devicemapper".
		GraphDriver string
	}
	// The output uses the SHA256 signature of the image as id.
	Output struct {
	}
)

// Id returns a unique id for the Inputs.
func (in Inputs) Id() api.Uuid {
	// Hash together all the input fields.
	h := sha512.New512_256()
	h.Write([]byte(in.Ancestor))
	h.Write([]byte(in.Dockerfile))
	for _, p := range in.Files {
		f, err := os.Open(p)
		if err != nil {
			log.Fatalf("FATAL: Couldn't read input file: %v", err)
		}
		if _, err := io.Copy(h, f); err != nil {
			log.Fatalf("FATAL: Couldn't copy from file to hash: %v", err)
		}
	}
	h.Write([]byte(in.DockerVersion))
	h.Write([]byte(in.Architecture))
	h.Write([]byte(in.Os))
	h.Write([]byte(in.GraphDriver))
	return api.Uuid(h.Sum(nil))
}

// Get retrieves the result of the computation.
//
// Get retrieves the result of the input from cache if it exists,
// from peers otherwise, computing it if needed.
func (c Computation) Get(in Inputs) (*Output, error) {
	return nil, fmt.Errorf("Get() is not implemented yet")
}
