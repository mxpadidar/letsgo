package specs

type Command interface {
	Validate() error
}
