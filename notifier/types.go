package notifier

type NotifierOptions struct {
	Endpoint string
}

type Notifier interface {
	Notify(message string)
}
