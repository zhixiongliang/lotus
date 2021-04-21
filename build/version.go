package build

var CurrentCommit string
var BuildType int

const (
	BuildDefault = iota
	BuildMainnet
	Build2k
	BuildDebug
	BuildCalibnet
	BuildNerpanet
	BuildButterflynet
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	case BuildNerpanet:
		return "+nerpanet"
	case BuildButterflynet:
		return "+butterflynet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.7.1-dev"

func UserVersion() string {
	return BuildVersion + buildType() + CurrentCommit
}
