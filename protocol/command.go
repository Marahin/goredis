package protocol

type Command interface {
	Evaluate([]string) Result
}
