package extern

func ImportFetchAll() (string, error) {
	// need a deps file

	return "TBD", nil
}

func ImportAddBundle(bundle, version string) (string, error) {
	err := CloneAndWrite(bundle, version)
	if err != nil {
		return "", err
	}

	// TODO update some deps file

	return "TBD", nil
}

func ImportUpdateBundle(bundle, version string) (string, error) {
	err := CloneAndWrite(bundle, version)
	if err != nil {
		return "", err
	}

	// TODO update some deps file

	return "TBD", nil
}

func ImportRemoveBundle(bundle string) (string, error) {
	// TODO update some deps file for version

	// Clone, walk, delete

	return "TBD", nil
}
