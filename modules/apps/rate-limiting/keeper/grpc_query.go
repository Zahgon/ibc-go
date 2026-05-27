package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	k *Keeper
}

func NewQuerier(keeper *Keeper) Querier {
	_ = "STUB: not implemented"
	return *

	// Query all rate limits
	new(Querier)
}

func (k Querier) AllRateLimits(c context.Context, req *types.QueryAllRateLimitsRequest) (*types.QueryAllRateLimitsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query a rate limit by denom and channelId
func (k Querier) RateLimit(c context.Context, req *types.QueryRateLimitRequest) (*types.QueryRateLimitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query all rate limits for a given chain
func (k Querier) RateLimitsByChainID(c context.Context, req *types.QueryRateLimitsByChainIDRequest) (*types.QueryRateLimitsByChainIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine the client state from the channel Id

// Check if the client state is a tendermint client

// Type assert to tendermint client state

// This should never happen if ClientType() == Tendermint, but check anyway

// If the chain ID matches, add the rate limit to the returned list

// Query all rate limits for a given channel
func (k Querier) RateLimitsByChannelOrClientID(c context.Context, req *types.QueryRateLimitsByChannelOrClientIDRequest) (*types.QueryRateLimitsByChannelOrClientIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query all blacklisted denoms
func (k Querier) AllBlacklistedDenoms(c context.Context, req *types.QueryAllBlacklistedDenomsRequest) (*types.QueryAllBlacklistedDenomsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query all whitelisted addresses
func (k Querier) AllWhitelistedAddresses(c context.Context, req *types.QueryAllWhitelistedAddressesRequest) (*types.QueryAllWhitelistedAddressesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
