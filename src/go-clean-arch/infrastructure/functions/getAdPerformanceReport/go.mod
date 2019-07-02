module go-clean-arch/infrastructure/functions/getAdPerformanceReport

require (
	cloud.google.com/go v0.41.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	go-clean-arch/application v0.0.0
	go-clean-arch/domain v0.0.0
	go-clean-arch/infrastructure v0.0.0 // indirect
	go-clean-arch/interfaces v0.0.0 // indirect
)

replace (
	go-clean-arch/application => ../../../application
	go-clean-arch/domain => ../../../domain
	go-clean-arch/infrastructure => ../../
	go-clean-arch/interfaces => ../../../interfaces
)
