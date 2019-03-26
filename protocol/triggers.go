package protocol

type Trigger string

var Triggers = map[Trigger]Command{
	PingTrigger: PingCommand{},
	SetTrigger:  SetCommand{},
	GetTrigger:  GetCommand{},
}
