package core

import (
	"context"

	bifrosteg "github.com/aperturerobotics/bifrost/entitygraph"
	"github.com/aperturerobotics/bifrost/link/establish"
	"github.com/aperturerobotics/bifrost/link/hold-open"
	nctr "github.com/aperturerobotics/bifrost/peer/controller"
	"github.com/aperturerobotics/bifrost/pubsub/floodsub/controller"
	"github.com/aperturerobotics/bifrost/pubsub/relay"
	"github.com/aperturerobotics/bifrost/stream/echo"
	"github.com/aperturerobotics/bifrost/stream/forwarding"
	"github.com/aperturerobotics/bifrost/stream/listening"
	iproctpt "github.com/aperturerobotics/bifrost/transport/inproc"
	udptpt "github.com/aperturerobotics/bifrost/transport/udp"
	wtpt "github.com/aperturerobotics/bifrost/transport/websocket"
	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/controller"
	"github.com/aperturerobotics/controllerbus/controller/resolver/static"
	cbc "github.com/aperturerobotics/controllerbus/core"
	egc "github.com/aperturerobotics/entitygraph/controller"
	"github.com/sirupsen/logrus"
)

// NewCoreBus constructs a standard in-memory bus stack with Bifrost controllers.
func NewCoreBus(
	ctx context.Context,
	le *logrus.Entry,
	builtInFactories ...controller.Factory,
) (bus.Bus, *static.Resolver, error) {
	b, sr, err := cbc.NewCoreBus(ctx, le, builtInFactories...)
	if err != nil {
		return nil, nil, err
	}

	AddFactories(b, sr)
	return b, sr, nil
}

// AddFactories adds factories to an existing static resolver.
func AddFactories(b bus.Bus, sr *static.Resolver) {
	// node controller
	sr.AddFactory(nctr.NewFactory())

	// link management controllers
	sr.AddFactory(link_holdopen_controller.NewFactory(b))
	sr.AddFactory(link_establish_controller.NewFactory(b))

	// stream controllers
	sr.AddFactory(stream_forwarding.NewFactory(b))
	sr.AddFactory(stream_echo.NewFactory(b))
	sr.AddFactory(stream_listening.NewFactory(b))

	// in-proc transport
	sr.AddFactory(iproctpt.NewFactory(b))
	// udp transport
	sr.AddFactory(udptpt.NewFactory(b))
	// websocket transport
	sr.AddFactory(wtpt.NewFactory(b))

	// pubsub
	sr.AddFactory(pubsub_relay.NewFactory(b))
	sr.AddFactory(floodsub_controller.NewFactory(b))

	// entity graph
	sr.AddFactory(egc.NewFactory(b))
	sr.AddFactory(bifrosteg.NewFactory(b))
}
