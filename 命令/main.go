package main

import command "p14/command"

func main() {
	tv := &command.Tv{}

	onCommand := &command.OnCommand{
		Device: tv,
	}
	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton:=&command.Button{
		Command: offCommand,
	}
	offButton.Press()
}