package cmd

import "github.com/mauwahid/kafman/internal/interfaces/process"

func RunSubscriber() {
	subs := process.NewSubscriber()
	go subs.Run()
}
