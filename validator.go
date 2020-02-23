package keyval

import "fmt"

type IntValidator func(int) error

func ValidateIntBetween(lower, upper int) IntValidator {
	return IntValidator(func(v int) error {
		switch {
		case v < lower:
			return fmt.Errorf("value %d is below lower limit %d", v, lower)
		case v > upper:
			return fmt.Errorf("value %d is above upper limit %d", v, upper)
		default:
			return nil
		}
	})
}

func ValidateIntInSet(values ...int) IntValidator {
	set := make(map[int]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return IntValidator(func(v int) error {
		if _, ok := set[v]; !ok {
			return fmt.Errorf("value %d is not in set %v", v, set)
		}

		return nil
	})
}

type StringValidator func(string) error

func ValidateStringInSet(values ...string) StringValidator {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return StringValidator(func(v string) error {
		if _, ok := set[v]; !ok {
			return fmt.Errorf("value '%s' not in set %v", v, set)
		}

		return nil
	})
}
