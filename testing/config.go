package ibctesting

import (
	"time"

	connectiontypes "github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	ibctm "github.com/cosmos/ibc-go/v11/modules/light-clients/07-tendermint"
)

type ClientConfig interface {
	GetClientType() string
}

type TendermintConfig struct {
	TrustLevel      ibctm.Fraction
	TrustingPeriod  time.Duration
	UnbondingPeriod time.Duration
	MaxClockDrift   time.Duration
}

func NewTendermintConfig() *TendermintConfig { _ = "STUB: not implemented"; return nil }

func (*TendermintConfig) GetClientType() string { _ = "STUB: not implemented"; return "" }

type ConnectionConfig struct {
	DelayPeriod uint64
	Version     *connectiontypes.Version
}

func NewConnectionConfig() *ConnectionConfig { _ = "STUB: not implemented"; return nil }

type ChannelConfig struct {
	PortID  string
	Version string
	Order   channeltypes.Order
}

func NewChannelConfig() *ChannelConfig { _ = "STUB: not implemented"; return nil }
