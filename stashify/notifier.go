package stashify

type Notifier interface {
	Notify() error
}
