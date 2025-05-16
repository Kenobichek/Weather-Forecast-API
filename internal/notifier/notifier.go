package notifier

type Notifier interface {
	Send(destination, message string) error
}
