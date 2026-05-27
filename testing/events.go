package ibctesting

import (
	testifysuite "github.com/stretchr/testify/suite"

	abci "github.com/cometbft/cometbft/abci/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

// ParseClientIDFromEvents parses events emitted from a MsgCreateClient and returns the
// client identifier.
func ParseClientIDFromEvents(events []abci.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseConnectionIDFromEvents parses events emitted from a MsgConnectionOpenInit or
// MsgConnectionOpenTry and returns the connection identifier.
func ParseConnectionIDFromEvents(events []abci.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseChannelIDFromEvents parses events emitted from a MsgChannelOpenInit or
// MsgChannelOpenTry or a MsgCreateChannel and returns the channel identifier.
func ParseChannelIDFromEvents(events []abci.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseV1PacketFromEvents parses events emitted from a send packet and returns
// the first EventTypeSendPacket packet found.
// Returns an error if no packet is found.
func ParseV1PacketFromEvents(events []abci.Event) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}

// ParseRecvV1PacketFromEvents parses events emitted from a MsgRecvPacket and returns
// the first EventTypeRecvPacket packet found.
// Returns an error if no packet is found.
func ParseRecvV1PacketFromEvents(events []abci.Event) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}

// ParseIBCV1Packets parses events and returns all the v1 packets found.
// Returns an error if no v1 packet is found.
func ParseIBCV1Packets(eventType string, events []abci.Event) ([]channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseV1PacketFromEvents parses events emitted from a v2 send packet
// and returns the first EventTypeSendPacket packet found.
// Returns an error if no packet is found.
func ParseV2PacketFromEvents(events []abci.Event) (channeltypesv2.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.Packet), nil
}

// ParseIBCV2Packets parses events and returns all the v2 packets found.
// Returns an error if no v2 packet is found.
func ParseIBCV2Packets(eventType string, events []abci.Event) ([]channeltypesv2.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we find a complete packet, we unmarshall it. We don't need to check for any other
// attributes from this event.

// Ignore unknown attributes

// ParseAckFromEvents parses events emitted from a MsgRecvPacket and returns the
// acknowledgement.
func ParseAckFromEvents(events []abci.Event) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseAckV2FromEvents parses events emitted from a MsgRecvPacket and returns the
// acknowledgement for v2 packets.
func ParseAckV2FromEvents(events []abci.Event) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseProposalIDFromEvents parses events emitted from MsgSubmitProposal and returns proposalID
func ParseProposalIDFromEvents(events []abci.Event) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ParsePacketSequenceFromEvents parses events emitted from MsgRecvPacket and returns the packet sequence
func ParsePacketSequenceFromEvents(events []abci.Event) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AssertEvents asserts that expected events are present in the actual events.
func AssertEvents(
	suite *testifysuite.Suite,
	expected []abci.Event,
	actual []abci.Event,
) {
	_ = "STUB: not implemented"
	return
}

// any expected attributes that are not contained in the actual events will cause this event
// not to match

// shouldProcessEvent returns true if the given expected event should be processed based on event type.
func shouldProcessEvent(expectedEvent abci.Event, actualEvent abci.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// the actual event will have an extra attribute added automatically
// by Cosmos SDK since v0.50, that's why we subtract 1 when comparing
// with the number of attributes in the expected event.

// containsAttribute returns true if the given key/value pair is contained in the given attributes.
// NOTE: this ignores the indexed field, which can be set or unset depending on how the events are retrieved.
func containsAttribute(attrs []abci.EventAttribute, key, value string) bool {
	_ = "STUB: not implemented"
	return false
}

// containsAttributeKey returns true if the given key is contained in the given attributes.
func containsAttributeKey(attrs []abci.EventAttribute, key string) bool {
	_ = "STUB: not implemented"
	return false
}

// attributeByKey returns the event attribute's value keyed by the given key and a boolean indicating its presence in the given attributes.
func attributeByKey(attributes []abci.EventAttribute, key string) (abci.EventAttribute, bool) {
	_ = "STUB: not implemented"
	return *new(abci.EventAttribute), false
}

// ParsePacketFromEvents parses events emitted from a send packet and returns
// the first EventTypeSendPacket packet found.
// Returns an error if no packet is found.
//
// Deprecated: This function will be removed in the next major release. Use
// ParseV1PacketFromEvents instead
func ParsePacketFromEvents(events []abci.Event) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}

// ParseRecvPacketFromEvents parses events emitted from a MsgRecvPacket and returns
// the first EventTypeRecvPacket packet found.
// Returns an error if no packet is found.
//
// Deprecated: This function will be removed in the next major release. Use
// ParseRecvV1PacketFromEvents instead
func ParseRecvPacketFromEvents(events []abci.Event) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}

// ParsePacketsFromEvents parses events emitted from a MsgRecvPacket and returns
// all the packets found.
// Returns an error if no packet is found.
//
// Deprecated: This function will be removed in the next major release. Use ParseIBCV1Packets instead.
func ParsePacketsFromEvents(eventType string, events []abci.Event) ([]channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
