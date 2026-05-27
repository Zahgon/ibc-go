package blsverifier

/*
 * This custom query handler is used to aggregate public keys and verify a signature using BLS.
 * It is used by the 08-wasm union light client, which we we use in the solidity IBC v2 e2e tests.
 * The code here is taken from here: https://github.com/unionlabs/union/tree/main/uniond/app/custom_query
 */
import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MessageSize = 32
)

type CustomQuery struct {
	AggregateVerify *QueryAggregateVerify `json:"aggregate_verify,omitempty"`
	Aggregate       *QueryAggregate       `json:"aggregate,omitempty"`
}
type QueryAggregate struct {
	PublicKeys [][]byte `json:"public_keys"`
}
type QueryAggregateVerify struct {
	PublicKeys [][]byte `json:"public_keys"`
	Signature  []byte   `json:"signature"`
	Message    []byte   `json:"message"`
}

func CustomQuerier() func(sdk.Context, json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G602 -- The index is controlled
