package mock

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

var _ api.IBCModule = (*IBCModule)(nil)

const (
	// ModuleNameA is a name that can be used for the first mock application.
	ModuleNameA = ModuleName + "A"
	// ModuleNameB is a name that can be used for the second mock application.
	ModuleNameB = ModuleName + "B"
	// PortIDA is a port ID that can be used for the first mock application.
	PortIDA = ModuleNameA
	// PortIDB is a port ID that can be used for the second mock application.
	PortIDB = ModuleNameB
)

// IBCModule is a mock implementation of the IBCModule interface.
// which delegates calls to the underlying IBCApp.
type IBCModule struct {
	IBCApp *IBCApp
}

// NewIBCModule creates a new IBCModule with an underlying mock IBC application.
func NewIBCModule() IBCModule { _ = "STUB: not implemented"; return *new(IBCModule) }

func (im IBCModule) OnSendPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, data channeltypesv2.Payload, signer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (im IBCModule) OnRecvPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) channeltypesv2.RecvPacketResult {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.RecvPacketResult)
}

func (im IBCModule) OnAcknowledgementPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, acknowledgement []byte, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (im IBCModule) OnTimeoutPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (IBCModule) UnmarshalPacketData(payload channeltypesv2.Payload) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
