package logging

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SdkEventsToLogArguments converts a given sdk.Events and returns a slice of strings that provide human
// readable values for the event attributes.
func SdkEventsToLogArguments(events sdk.Events) []string { _ = "STUB: not implemented"; return nil }
