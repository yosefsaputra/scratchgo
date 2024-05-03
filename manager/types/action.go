package types

type Action interface {
	Run() error
}
