# goka-tester-race

```bash
go test -count=10 -race .
```

```
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
==================
WARNING: DATA RACE
Write at 0x00c000122010 by goroutine 14:
  runtime.recvDirect()
      /opt/homebrew/Cellar/go/1.24.2/libexec/src/runtime/chan.go:405 +0x7c
  github.com/lovoo/goka/tester.(*cgClaim).close()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:380 +0xdc
  github.com/lovoo/goka/tester.(*consumerGroup).Consume.func2()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:85 +0xcc
  github.com/lovoo/goka/tester.(*consumerGroup).Consume.(*ErrGroup).Go.func4()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/multierr/errgroup.go:48 +0x40
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/jakub.gruszecki/go/pkg/mod/golang.org/x/sync@v0.11.0/errgroup/errgroup.go:78 +0x7c

Previous read at 0x00c000122010 by goroutine 16:
  runtime.chansend1()
      /opt/homebrew/Cellar/go/1.24.2/libexec/src/runtime/chan.go:162 +0x2c
  github.com/lovoo/goka/tester.(*cgSession).pushMessageToClaim()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:325 +0x49c
  github.com/lovoo/goka/tester.(*cgSession).catchupAndWait()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:286 +0x394
  github.com/lovoo/goka/tester.(*consumerGroup).catchupAndWait()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:53 +0xa8
  github.com/lovoo/goka/tester.(*client).catchup()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/client.go:24 +0x5c
  github.com/lovoo/goka/tester.(*Tester).waitForClients()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/tester.go:369 +0x174
  github.com/lovoo/goka/tester.(*Tester).Consume()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/tester.go:419 +0x1ec
  github.com/ditrytus/goka-tester-race.TestGokaRace.gowrap2()
      /Users/jakub.gruszecki/Documents/goka-tester-race/goka_tester_race_test.go:44 +0x70

Goroutine 14 (running) created at:
  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/jakub.gruszecki/go/pkg/mod/golang.org/x/sync@v0.11.0/errgroup/errgroup.go:75 +0x10c
  github.com/lovoo/goka/multierr.(*ErrGroup).Go()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/multierr/errgroup.go:47 +0x6d0
  github.com/lovoo/goka/tester.(*consumerGroup).Consume()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/tester/consumergroup.go:80 +0x548
  github.com/lovoo/goka.(*Processor).rebalanceLoop()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/processor.go:382 +0x7d8
  github.com/lovoo/goka.(*Processor).Run.func6()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/processor.go:333 +0x54
  github.com/lovoo/goka.(*Processor).Run.(*ErrGroup).Go.func8()
      /Users/jakub.gruszecki/go/pkg/mod/github.com/lovoo/goka@v1.1.14/multierr/errgroup.go:48 +0x40
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/jakub.gruszecki/go/pkg/mod/golang.org/x/sync@v0.11.0/errgroup/errgroup.go:78 +0x7c

Goroutine 16 (finished) created at:
  github.com/ditrytus/goka-tester-race.TestGokaRace()
      /Users/jakub.gruszecki/Documents/goka-tester-race/goka_tester_race_test.go:44 +0x6a4
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.24.2/libexec/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.24.2/libexec/src/testing/testing.go:1851 +0x40
==================
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
--- FAIL: TestGokaRace (0.11s)
    testing.go:1490: race detected during execution of test
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:15 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:15 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:16 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:16 [Processor test-group] Consumer group cancelled. Stopping
2025/08/04 12:47:16 [Processor test-group] setup generation 0, claims=map[string][]int32{"test-stream":[]int32{0}}
2025/08/04 12:47:16 [Processor test-group] Consumer group cancelled. Stopping
FAIL
FAIL	github.com/ditrytus/goka-tester-race	1.416s
FAIL
```
