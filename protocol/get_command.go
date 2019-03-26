package protocol

import (
	"bytes"
	"encoding/gob"

	logger "github.com/sirupsen/logrus"

	"github.com/marahin/goredis/rdb"
)

const GetTrigger = Trigger("get")

type GetCommand struct {
	Command
}

func (c GetCommand) Evaluate(args []string) Result {
	logger.WithFields(logger.Fields{"args": args}).Debug("GetCommand evaluate")

	var buf bytes.Buffer

	resp := rdb.GeneralDataStore.Get(args[0])
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(resp)

	if err != nil {
		panic(err)
	}

	return Result{buf.Bytes()}
}
