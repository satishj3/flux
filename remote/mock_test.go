package remote

import (
	"testing"

	"github.com/weaveworks/flux/api"
)

// Just test that the mock does its job.
func TestMock(t *testing.T) {
	ServerTestBattery(t, func(mock api.UpstreamServer) api.UpstreamServer { return mock })
}
