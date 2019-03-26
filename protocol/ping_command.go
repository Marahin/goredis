package protocol

const PingTrigger = Trigger("ping")

type PingCommand struct {
	Command
}

func (c PingCommand) Evaluate(_ []string) Result {
	return Result{[]byte("PONG")}
}
