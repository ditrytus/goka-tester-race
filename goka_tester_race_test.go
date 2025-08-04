// This file demonstrates a race condition in the Goka tester library.
// The race occurs in github.com/lovoo/goka/tester between:
// - cgClaim.close() closing a channel
// - cgSession.pushMessageToClaim() sending to that channel
//
// Run with: go test -race goka_tester_race_test.go

package main

import (
	"context"
	"testing"
	"time"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/lovoo/goka/tester"
)

func TestGokaRace(t *testing.T) {
	gokaTester := tester.New(t)

	ctx, cancel := context.WithCancel(context.Background())

	// Create a processor with slow processing
	proc, _ := goka.NewProcessor(
		nil,
		goka.DefineGroup(
			goka.Group("test-group"),
			goka.Input(goka.Stream("test-stream"), new(codec.String),
				func(ctx goka.Context, msg interface{}) {
					// Simulate slow processing
					time.Sleep(50 * time.Millisecond)
				}),
		),
		goka.WithTester(gokaTester),
	)

	// Start processor
	go proc.Run(ctx)
	proc.WaitForReadyContext(ctx)

	// Send message
	go gokaTester.Consume("test-stream", "key", "message")

	// Cancel context while processing
	// This triggers the race as the processor shuts down
	// while the tester is still active
	time.Sleep(10 * time.Millisecond)
	cancel()

	// Cleanup time
	time.Sleep(100 * time.Millisecond)
}
