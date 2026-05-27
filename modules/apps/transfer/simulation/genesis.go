package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
)

// Simulation parameter constants
const port = "port_id"

// RandomEnabled randomized send or receive enabled param with 75% prob of being true.
func RandomEnabled(r *rand.Rand) bool { _ = "STUB: not implemented"; return false }

// RandomizedGenState generates a random GenesisState for transfer.
func RandomizedGenState(simState *module.SimulationState) { _ = "STUB: not implemented"; return }
