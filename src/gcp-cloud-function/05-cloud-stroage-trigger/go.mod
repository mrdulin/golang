module gcp-cloud-function/05-cloud-stroage-trigger

require (
	cloud.google.com/go v0.41.0
	gcp-cloud-function/05-cloud-stroage-trigger/utils v0.0.0
	google.golang.org/api v0.7.0
)

replace gcp-cloud-function/05-cloud-stroage-trigger/utils => ./utils
