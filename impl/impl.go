// Package impl defines the ucc implentation.
package impl // import "hkjn.me/ucc/impl"

import "fmt"

type (
	// Cache holds in-memory data for Inputs Uuid to Output.
	Cache map[Uuid]Output
	// Impl implements api.Computation.
	Impl struct {
		Cache cache
	}
)

// Get retrieves data.
func (impl Impl) Get(in Inputs) (*Output, error) {
	// Retrieve from cache if available.
	out, cached := impl.Cache[in.Id()]
	if cached {
		return out, nil
	}

	out, err := peers.Fetch(in)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch %v from peers: %v", in, err)
	}

	return fmt.Errorf("TODO: Perform the f(in)->out computation locally here, abort it if peers return result before we're done")
	// TODO: Also store timing info on how expensive the computation was? That way, other nodes could query for cost before starting computation.
}
