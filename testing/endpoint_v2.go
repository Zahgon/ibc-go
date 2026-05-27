package ibctesting

import (
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

// RegisterCounterparty will construct and execute a MsgRegisterCounterparty on the associated ep.
func (ep *Endpoint) RegisterCounterparty() error { _ = "STUB: not implemented"; return nil }

// setup counterparty

// MsgSendPacket sends a packet on the associated endpoint using a predefined sender. The constructed packet is returned.
func (ep *Endpoint) MsgSendPacket(timeoutTimestamp uint64, payloads ...channeltypesv2.Payload) (channeltypesv2.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Packet), nil
}

// MsgSendPacketWithSender sends a packet on the associated endpoint using the provided sender. The constructed packet is returned.
func (ep *Endpoint) MsgSendPacketWithSender(timeoutTimestamp uint64, payloads []channeltypesv2.Payload, sender SenderAccount) (channeltypesv2.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Packet), nil
}

// TODO: parse the packet from events instead of from the response. https://github.com/cosmos/ibc-go/issues/7459
// get sequence from msg response

// MsgRecvPacket sends a MsgRecvPacket on the associated endpoint with the provided packet.
func (ep *Endpoint) MsgRecvPacket(packet channeltypesv2.Packet) error {
	_ = "STUB: not implemented"
	// get proof of packet commitment from chainA
	return nil
}

// MsgRecvPacketWithAck returns the acknowledgement for the given packet by sending a MsgRecvPacket on the associated endpoint.
func (ep *Endpoint) MsgRecvPacketWithAck(packet channeltypesv2.Packet) (channeltypesv2.Acknowledgement, error) {
	_ = "STUB: not implemented"
	// get proof of packet commitment from chainA
	return *new(channeltypesv2.Acknowledgement), nil
}

// MsgAcknowledgePacket sends a MsgAcknowledgement on the associated endpoint with the provided packet and ack.
func (ep *Endpoint) MsgAcknowledgePacket(packet channeltypesv2.Packet, ack channeltypesv2.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

// MsgTimeoutPacket sends a MsgTimeout on the associated endpoint with the provided packet.
func (ep *Endpoint) MsgTimeoutPacket(packet channeltypesv2.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// RelayPacket relayes packet that was previously sent on the given endpoint.
func (ep *Endpoint) RelayPacket(packet channeltypesv2.Packet) error {
	_ = "STUB: not implemented"
	// receive packet on counterparty
	return nil
}

// acknowledge packet on endpoint
