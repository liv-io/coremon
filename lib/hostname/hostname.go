package hostname

import (
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func GetHostname(version string) (name string) {
	switch version {
	case "long":
		switch os := runtime.GOOS; os {
		case "openbsd":
			cmd := exec.Command("/bin/hostname")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			name = out.String()
			name = name[:len(name)-1] // Remove newline
			name = strings.ToLower(name)
			return
		default:
			cmd := exec.Command("/bin/hostname", "-f")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			name = out.String()
			name = name[:len(name)-1] // Remove newline
			name = strings.ToLower(name)
			return
		}
	default:
		cmd := exec.Command("/bin/hostname", "-s")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		name = out.String()
		name = name[:len(name)-1] // Remove newline
		name = strings.ToLower(name)
		return
	}
}
