module github.com/omniboost/go-fortnox

require (
	github.com/go-redis/redis/v9 v9.0.0-beta.2
	github.com/gorilla/schema v1.2.0
	golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

go 1.19

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
