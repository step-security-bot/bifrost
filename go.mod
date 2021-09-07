module github.com/aperturerobotics/bifrost

go 1.16

// aperture: use compatibility forks
replace (
	github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x
	github.com/lucas-clemente/quic-go => github.com/aperturerobotics/quic-go v0.23.1-0.20210906210823-4c73aa0bb4f2 // aperture
	github.com/nats-io/nats-server/v2 => github.com/aperturerobotics/bifrost-nats-server/v2 v2.1.8-0.20200831101324-59acc8fe7f74 // aperture-2.0
	github.com/nats-io/nats.go => github.com/aperturerobotics/bifrost-nats-client v1.10.1-0.20200831103200-24c3d0464e58 // aperture-2.0
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => github.com/paralin/grpc-go v1.30.1-0.20210804030014-1587a7c16b66 // aperture
)

require (
	github.com/aperturerobotics/controllerbus v0.8.6-0.20210902104809-9f0fc115965e
	github.com/aperturerobotics/entitygraph v0.1.4-0.20210530040557-f19da9c2be6d
	github.com/aperturerobotics/timestamp v0.2.4-0.20210530040952-1422410fbd4a
	github.com/blang/semver v3.5.1+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/djherbis/buffer v1.1.0
	github.com/frankban/quicktest v1.10.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.3
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/hashicorp/yamux v0.0.0-20210826001029-26ff87cf9493
	github.com/klauspost/compress v1.13.6-0.20210906113219-38d4ba985ac1
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-tls v0.2.0
	github.com/lucas-clemente/quic-go v0.23.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.3.0
	github.com/nats-io/nats-server/v2 v2.3.1
	github.com/nats-io/nats.go v1.11.0
	github.com/nats-io/nkeys v0.3.0
	github.com/paralin/kcp-go-lite v1.0.2-0.20210907042029-a4cf7baab73a // aperture
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pauleyj/gobee v0.0.0-20190212035730-6270c53072a4
	github.com/pierrec/lz4 v2.6.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b
	github.com/tjfoc/gmsm v1.4.0
	github.com/urfave/cli v1.22.5
	github.com/xtaci/smux v1.5.15
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	google.golang.org/grpc v1.39.0
	gortc.io/stun v1.23.0
	nhooyr.io/websocket v1.8.7
)
