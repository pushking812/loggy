package exemplar1

import ls "github.com/pushking812/loggy/logging/simple"

func Run(debug bool) {
	ls.Debug(debug)

	ls.Log("This is a debug statement... (Exemplar2)")
}
