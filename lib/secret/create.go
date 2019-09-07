package secret

import "fmt"

func Create(name, value string) error {

	m, err := readSecrets()
	if err != nil {
		return err
	}

	_, ok := m[name]
	if ok {
		fmt.Println("Warning: overwriting existing value")
	}

	m[name] = value

	err = writeSecrets(m)
	if err != nil {
		return err
	}

	return nil
}
