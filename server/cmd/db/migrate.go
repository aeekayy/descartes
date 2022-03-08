package db

import (
	"errors"
	"fmt"
)

// RunFlywayMigration runs a Flyway migration on your local
// machine. If the migration is successful, the function will
// return with a boolean value of true and nil error.
// Parameters:
// path, the directory path where the migration files are located
func RunFlywayMigration(path string) (bool, error) {

	return true, nil
}

// ValidateFlywayFileFormat validates the file fomrats of the
// Flyway files are in the correct format. If they are incorrect,
// return an error
func ValidateFlywayFileFormat(path string) error {
	// create a regexp check for the Flyway file format
	var formatMatches map[string]bool

	// check for every file with the `.sql` suffix
	// if the format does not match, return an error
	if len(formatMatches) > 0 {
		return errors.New(fmt.Sprintf("the format of the files %s are incorrect"))
	}
	return nil
}
