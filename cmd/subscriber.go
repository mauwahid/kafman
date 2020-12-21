package cmd

import "github.com/mauwahid/kafman/internal/presenter/process"

func RunSubscriber() {
	subs := process.NewSubscriber()
	go subs.Run()
}
