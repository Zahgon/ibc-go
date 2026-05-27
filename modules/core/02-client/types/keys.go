package types

import (
	"regexp"
)

const (
	// SubModuleName defines the IBC client name
	SubModuleName string = "client"

	// RouterKey is the message route for IBC client
	RouterKey string = SubModuleName

	// QuerierRoute is the querier route for IBC client
	QuerierRoute string = SubModuleName

	// KeyNextClientSequence is the key used to store the next client sequence in
	// the keeper.
	KeyNextClientSequence = "nextClientSequence"

	// ParamsKey is the store key for the IBC client parameters
	ParamsKey = "clientParams"

	// KeyCreator is the key for the creator in the client-specific store
	KeyCreator = "creator"

	// AllowAllClients is the value that if set in AllowedClients param
	// would allow any wired up light client modules to be allowed
	AllowAllClients = "*"
)

var KeyAllowedClients = []byte("AllowedClients")

// FormatClientIdentifier returns the client identifier with the sequence appended.
// This is an SDK specific format not enforced by IBC protocol.
func FormatClientIdentifier(clientType string, sequence uint64) string {
	_ = "STUB: not implemented"
	return ""
}

// IsClientIDFormat checks if a clientID is in the format required on the SDK for
// parsing client identifiers. The client identifier must be in the form: `{client-type}-{N}
// which per the specification only permits ASCII for the {client-type} segment and
// 1 to 20 digits for the {N} segment.
// `([\w-]+\w)?` allows for a letter or hyphen, with the {client-type} starting with a letter
// and ending with a letter, i.e. `letter+(letter|hyphen+letter)?`.
var IsClientIDFormat = regexp.MustCompile(`^\w+([\w-]+\w)?-[0-9]{1,20}$`).MatchString

// IsValidClientID checks if the clientID is valid and can be parsed into the client
// identifier format.
func IsValidClientID(clientID string) bool { _ = "STUB: not implemented"; return false }

// ParseClientIdentifier parses the client type and sequence from the client identifier.
func ParseClientIdentifier(clientID string) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// MustParseClientIdentifier parses the client type from the provided client identifier.
// If an invalid client identifier is provided this function will panic.
func MustParseClientIdentifier(clientID string) string { _ = "STUB: not implemented"; return "" }

// CreatorKey returns the key under which the client creator is stored in the client store
func CreatorKey() []byte { _ = "STUB: not implemented"; return nil }
