package mock

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	MockEventType              = "mock-event-type"
	MockEventTypeRecvPacket    = "mock-recv-packet"
	MockEventTypeAckPacket     = "mock-ack-packet"
	MockEventTypeTimeoutPacket = "mock-timeout"

	MockAttributeKey1 = "mock-attribute-key-1"
	MockAttributeKey2 = "mock-attribute-key-2"

	MockAttributeValue1 = "mock-attribute-value-1"
	MockAttributeValue2 = "mock-attribute-value-2"
)

// NewMockRecvPacketEvent returns a mock receive packet event
func NewMockRecvPacketEvent() sdk.Event { _ = "STUB: not implemented"; return *new(sdk.Event) }

// NewMockAckPacketEvent returns a mock acknowledgement packet event
func NewMockAckPacketEvent() sdk.Event { _ = "STUB: not implemented"; return *new(sdk.Event) }

// NewMockTimeoutPacketEvent emits a mock timeout packet event
func NewMockTimeoutPacketEvent() sdk.Event { _ = "STUB: not implemented"; return *new(sdk.Event) }

// emitMockEvent returns a mock event with the given event type
func newMockEvent(eventType string) sdk.Event { _ = "STUB: not implemented"; return *new(sdk.Event) }
