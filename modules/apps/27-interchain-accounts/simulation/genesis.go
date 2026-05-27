package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
)

// RandomEnabled randomized controller or host enabled param with 75% prob of being true.
func RandomEnabled(r *rand.Rand) bool { _ = "STUB: not implemented"; return false }

// RandomizedGenState generates a random GenesisState for ics27.
// Only the params are non nil
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }

// allow all messages
