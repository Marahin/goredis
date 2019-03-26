package protocol

import "github.com/marahin/goredis/rdb"

const SetTrigger = Trigger("set")

type SetCommand struct {
	Command
}

func (c SetCommand) Evaluate(args []string) Result {
	rdb.GeneralDataStore.Set(args[0], args[1:])

	return Result{[]byte("OK")}
}
