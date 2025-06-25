module task25/proxy

go 1.19

replace (
	golang.org/x/crypto v0.39.0 => golang.org/x/crypto v0.12.0
	golang.org/x/net => golang.org/x/net v0.7.0
	golang.org/x/sys v0.33.0 => golang.org/x/sys v0.6.0
	golang.org/x/tools => golang.org/x/tools v0.4.0
)

require (
	github.com/ekomobile/dadata/v2 v2.16.0
	github.com/go-chi/chi v1.5.5
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/prometheus/client_golang v1.20.5
	github.com/swaggo/http-swagger v1.3.4
	github.com/swaggo/swag v1.16.4
	golang.org/x/crypto v0.39.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.37.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
