package types

var (
	// DefaultIBCVersion represents the latest supported version of IBC used
	// in connection version negotiation. The current version supports the list
	// of orderings defined in SupportedOrderings and requires at least one channel type
	// to be agreed upon.
	DefaultIBCVersion = NewVersion(DefaultIBCVersionIdentifier, SupportedOrderings)

	// DefaultIBCVersionIdentifier is the IBC v1.0.0 protocol version identifier
	DefaultIBCVersionIdentifier = "1"

	// SupportedOrderings is the list of orderings supported by IBC. The current
	// version supports only ORDERED and UNORDERED channels.
	SupportedOrderings = []string{"ORDER_ORDERED", "ORDER_UNORDERED"}

	// AllowNilFeatureSet is a helper map to indicate if a specified version
	// identifier is allowed to have a nil feature set. Any versions supported,
	// but not included in the map default to not supporting nil feature sets.
	allowNilFeatureSet = map[string]bool{
		DefaultIBCVersionIdentifier: false,
	}

	// MaxVersionsLength is the maximum number of versions that can be supported
	MaxCounterpartyVersionsLength = 100
	// MaxFeaturesLength is the maximum number of features that can be supported
	MaxFeaturesLength = 100
)

// NewVersion returns a new instance of Version.
func NewVersion(identifier string, features []string) *Version {
	_ = "STUB: not implemented"
	return nil
}

// GetIdentifier implements the VersionI interface
func (v Version) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

// GetFeatures implements the VersionI interface
func (v Version) GetFeatures() []string {
	_ = "STUB: not implemented"

	// ValidateVersion does basic validation of the version identifier and
	// features. It unmarshals the version string into a Version object.
	return nil
}

func ValidateVersion(version *Version) error { _ = "STUB: not implemented"; return nil }

// VerifyProposedVersion verifies that the entire feature set in the
// proposed version is supported by this chain. If the feature set is
// empty it verifies that this is allowed for the specified version
// identifier.
func (v Version) VerifyProposedVersion(proposedVersion *Version) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifySupportedFeature takes in a version and feature string and returns
// true if the feature is supported by the version and false otherwise.
func VerifySupportedFeature(version *Version, feature string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetCompatibleVersions returns a descending ordered set of compatible IBC
// versions for the caller chain's connection end. The latest supported
// version should be first element and the set should descend to the oldest
// supported version.
func GetCompatibleVersions() []*Version { _ = "STUB: not implemented"; return nil }

// IsSupportedVersion returns true if the proposed version has a matching version
// identifier and its entire feature set is supported or the version identifier
// supports an empty feature set.
func IsSupportedVersion(supportedVersions []*Version, proposedVersion *Version) bool {
	_ = "STUB: not implemented"
	return false
}

// FindSupportedVersion returns the version with a matching version identifier
// if it exists. The returned boolean is true if the version is found and
// false otherwise.
func FindSupportedVersion(version *Version, supportedVersions []*Version) (*Version, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// PickVersion iterates over the descending ordered set of compatible IBC
// versions and selects the first version with a version identifier that is
// supported by the counterparty. The returned version contains a feature
// set with the intersection of the features supported by the source and
// counterparty chains. If the feature set intersection is nil and this is
// not allowed for the chosen version identifier then the search for a
// compatible version continues. This function is called in the ConnOpenTry
// handshake procedure.
//
// CONTRACT: PickVersion must only provide a version that is in the
// intersection of the supported versions and the counterparty versions.
func PickVersion(supportedVersions, counterpartyVersions []*Version) (*Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if the source version is supported by the counterparty

// GetFeatureSetIntersection returns the intersections of source feature set
// and the counterparty feature set. This is done by iterating over all the
// features in the source version and seeing if they exist in the feature
// set for the counterparty version.
func GetFeatureSetIntersection(sourceFeatureSet, counterpartyFeatureSet []string) []string {
	_ = "STUB: not implemented"
	return nil
}
