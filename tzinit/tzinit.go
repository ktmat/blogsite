// This file is for setting the timemzone for the Gin framework.

package tzinit

import (
	"os"
)

func init() {
	os.Setenv("TZ", "Australia/Brisbane")
}