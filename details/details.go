package details

import (
	"os"
)

func GetHostName() (string, error) {
	return os.Hostname()
}
