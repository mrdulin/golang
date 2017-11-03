package ad

type AdState string

const (
	ENABLED  AdState = "enabled"
	PAUSED   AdState = "paused"
	DISABLED AdState = "disabled"
)

type ApprovalStatus string

const (
	APPROVED         ApprovalStatus = "approved"
	APPROVED_LIMITED ApprovalStatus = "approved (limited)"
	ELIGIBLE         ApprovalStatus = "eligible"
	UNDER_REVIEW     ApprovalStatus = "under review"
	DISAPPROVED      ApprovalStatus = "disapproved"
	SITE_SUSPENDED   ApprovalStatus = "site suspended"
)
