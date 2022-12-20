package helpers

import (
	"github.com/Thomascogez/dump-and-dumper/notifier"
	"github.com/spf13/pflag"
)

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func GetNotifiersFromFlags(flags *pflag.FlagSet) []notifier.Notifier {
	notifiers := make([]notifier.Notifier, 0)
	if flags.Changed("ntfy-endpoint") {
		notifiers = append(notifiers, notifier.NtfyNotifier{
			NotifierOptions: notifier.NotifierOptions{
				Endpoint: flags.Lookup("ntfy-endpoint").Value.String(),
			},
		})
	}

	return notifiers
}
