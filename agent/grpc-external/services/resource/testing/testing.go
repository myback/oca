package testing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/myback/oca/acl/resolver"
	svc "github.com/myback/oca/agent/grpc-external/services/resource"
	internal "github.com/myback/oca/agent/grpc-internal"
	"github.com/myback/oca/internal/resource"
	"github.com/myback/oca/internal/storage/inmem"
	"github.com/myback/oca/proto-public/pbresource"
	"github.com/myback/oca/sdk/testutil"
)

// RunResourceService runs a Resource Service for the duration of the test and
// returns a client to interact with it. ACLs will be disabled.
func RunResourceService(t *testing.T, registerFns ...func(resource.Registry)) pbresource.ResourceServiceClient {
	t.Helper()

	backend, err := inmem.NewBackend()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go backend.Run(ctx)

	registry := resource.NewRegistry()
	for _, fn := range registerFns {
		fn(registry)
	}

	server := grpc.NewServer()

	svc.NewServer(svc.Config{
		Backend:     backend,
		Registry:    registry,
		Logger:      testutil.Logger(t),
		ACLResolver: resolver.DANGER_NO_AUTH{},
	}).Register(server)

	pipe := internal.NewPipeListener()
	go server.Serve(pipe)
	t.Cleanup(server.Stop)

	conn, err := grpc.Dial("",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(pipe.DialContext),
		grpc.WithBlock(),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = conn.Close() })

	return pbresource.NewResourceServiceClient(conn)
}
