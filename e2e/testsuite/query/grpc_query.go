package query

import (
	"context"

	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/interchaintest/v11/ibc"
	"google.golang.org/grpc"

	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
)

var queryReqToPath = make(map[string]string)

func PopulateQueryReqToPath(ctx context.Context, chain ibc.Chain) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip services that are annotated with the "cosmos.msg.v1.service" option.

// trim the first character from input which is a dot

// GRPCQuery queries the chain with a query request and deserializes the response to T
func GRPCQuery[T any](ctx context.Context, chain ibc.Chain, req proto.Message, opts ...grpc.CallOption) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// grpcQueryWithMethod queries the chain with a query request with a specific method (grpc path) and deserializes the response to T
func grpcQueryWithMethod[T any](ctx context.Context, chain ibc.Chain, req proto.Message, method string, opts ...grpc.CallOption) (*T, error) {
	_ = "STUB: not implemented"
	// Create a connection to the gRPC server.
	return nil, nil
}

func queryFileDescriptors(ctx context.Context, chain ibc.Chain) (*reflectionv1.FileDescriptorsResponse, error) {
	_ = "STUB: not implemented"
	// Create a connection to the gRPC server.
	return nil, nil
}

// TODO: Remove all of the below when v6 -> v7 upgrade is not supported anymore:
func getProtoPath(req proto.Message) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getQueryProtoPath(queryTypeURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Add to the index to account for the length of "Query"

// Add a slash before the query

func getCmtProtoPath(cmtTypeURL string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Add a slash before the commitment
