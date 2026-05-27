package mock

import (
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	mockv1 "github.com/cosmos/ibc-go/v11/testing/mock"
)

const (
	ModuleName = "mockv2"
)

var MockRecvPacketResult = channeltypesv2.RecvPacketResult{
	Status:          channeltypesv2.PacketStatus_Success,
	Acknowledgement: mockv1.MockAcknowledgement.Acknowledgement(),
}

func NewMockPayload(sourcePort, destPort string) channeltypesv2.Payload {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Payload)
}

func NewErrorMockPayload(sourcePort, destPort string) channeltypesv2.Payload {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Payload)
}

func NewAsyncMockPayload(sourcePort, destPort string) channeltypesv2.Payload {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Payload)
}
