package errors

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ConvertToErrorEvents converts all events to error events by appending the
// error attribute prefix to each event's attribute key.
func ConvertToErrorEvents(events sdk.Events) sdk.Events {
	_ = "STUB: not implemented"
	return *new(sdk.Events)
}
