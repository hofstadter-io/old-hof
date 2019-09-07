package secret

import (
	"fmt"
)

func Update(name, value string) error {

	m, err := readSecrets()
	if err != nil {
		return err
	}

	_, ok := m[name]
	if !ok {
		fmt.Println("Warning: no value to update")
	}

	m[name] = value

	err = writeSecrets(m)
	if err != nil {
		return err
	}

	return nil
}
