//golangcitest:args -Einterfacer --internal-cmd-test
package testdata

import "io"

func InterfacerCheck(f io.ReadCloser) { // ERROR "`f` can be `io.Closer`"
	f.Close()
}
