package adGroup

type AdGroupState string

const (
	ENABLED AdGroupState = "enabled"
	PAUSED  AdGroupState = "paused"
	REMOVED AdGroupState = "removed"
)
