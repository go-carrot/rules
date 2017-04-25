package rules

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func IsSet(name, input string) error {
	if input == "" {
		return errors.New(fmt.Sprintf("Missing parameter '%v'", name))
	}
	return nil
}

func MaxLen(maxLength int) func(name, input string) error {
	return func(name, input string) error {
		if len(input) > maxLength {
			return errors.New(fmt.Sprintf("Maximum allowed length of parameter '%v' is %v", name, maxLength))
		}
		return nil
	}
}

func MinLen(minLength int) func(name, input string) error {
	return func(name, input string) error {
		if len(input) < minLength {
			return errors.New(fmt.Sprintf("Minimum allowed length of parameter '%v' is %v", name, minLength))
		}

		return nil
	}
}

func IsEmail(name, input string) error {
	regex, _ := regexp.Compile("^.+@.+\\..+$")
	if !regex.MatchString(input) {
		return errors.New(fmt.Sprintf("Parameter '%v' must be a valid email address", name))
	}
	return nil
}

func MaxValue(maxValue float64) func(name, input string) error {
	return func(name, input string) error {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}
		if value > maxValue {
			return errors.New(fmt.Sprintf("Maximum allowed value of parameter '%v' is %v", name, maxValue))
		}
		return nil
	}
}

func MinValue(minValue float64) func(name, input string) error {
	return func(name, input string) error {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}
		if value < minValue {
			return errors.New(fmt.Sprintf("Minimum allowed value of parameter '%v' is %v", name, minValue))
		}
		return nil
	}
}
