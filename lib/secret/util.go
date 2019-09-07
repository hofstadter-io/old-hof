package secret

import (
	"io/ioutil"
	"os"
	"strings"
)

const filename = "secrets.env"

func ensureFile() {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		ioutil.WriteFile(filename, []byte{}, 0644)
	}
}

func readSecrets() (map[string]string, error) {
	ensureFile()

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	for _, line := range strings.Split(string(contents), "\n") {
		if line != "" {
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			m[key] = parts[1]
		}
	}

	return m, nil
}

func writeSecrets(m map[string]string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for key, val := range m {
		line := key + "=" + val
		file.Write([]byte(line + "\n"))
	}

	return nil
}
