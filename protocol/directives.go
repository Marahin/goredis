package protocol

import (
	"regexp"
	"strings"

	logger "github.com/sirupsen/logrus"
)

//Has to match:
//	PING
//	SET a raz dwa trzy
//	GET a
const Matcher = `^(\w+) ?(\w+)? ?(.*)?$`

type Result struct {
	Response []byte
}

type Directive struct {
	command   Command
	arguments []string
}

func Determine(payload string) *Directive {
	logger.WithFields(logger.Fields{"payload": payload}).Debug("Determining by payload")

	re := regexp.MustCompile(Matcher)
	result := re.FindStringSubmatch(payload)

	logger.WithFields(logger.Fields{"result": result}).Debug("Regexp payload matching")

	stringifiedCommand, arguments := result[1], result[2:] // set, [a, raz dwa trzy]
	trigger := Trigger(strings.ToLower(stringifiedCommand))

	if val, ok := Triggers[trigger]; ok {
		return &Directive{val, arguments}
	}

	return &Directive{UnknownCommand{}, arguments}
}

func (d *Directive) Evaluate() Result {
	return d.command.Evaluate(d.arguments)
}
