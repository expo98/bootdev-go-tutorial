package main

import "fmt"

func validateStatus(status string) error {

	if status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	if len(status) > 140 {
		return fmt.Errorf("status exceeds 140 characters")
	}

	return nil
}
