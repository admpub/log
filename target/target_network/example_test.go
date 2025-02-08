package target_network

import "github.com/admpub/log"

func ExampleNewNetworkTarget() {
	logger := log.NewLogger()

	// creates a NetworkTarget which uses tcp network and address :10234
	target := NewNetworkTarget()
	target.Network = "tcp"
	target.Address = ":10234"

	logger.Targets = append(logger.Targets, target)

	logger.Open()

	// ... logger is ready to use ...
}
