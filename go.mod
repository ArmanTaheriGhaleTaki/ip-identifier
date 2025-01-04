module IpIdentifier

go 1.23.2

require (
	github.com/stretchr/testify v1.8.1
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
	go.opentelemetry.io/otel
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog
	go.opentelemetry.io/otel/sdk/log
	go.opentelemetry.io/otel/log/global
	go.opentelemetry.io/otel/propagation
	go.opentelemetry.io/otel/sdk/metric
	go.opentelemetry.io/otel/sdk/resource
	go.opentelemetry.io/otel/sdk/trace
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
	go.opentelemetry.io/contrib/bridges/otelslog
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
