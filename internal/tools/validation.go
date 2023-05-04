package tools

import (
	"errors"
	"strconv"
)

func ValidateInn(inn string) error {

	if len(inn) == 0 {
		return &ServiceError{Code: 3, Err: errors.New("inn is required")}
	}
	if len(inn) != 10 && len(inn) != 12 {
		return &ServiceError{Code: 3, Err: errors.New("inn length must be 10 or 12")}
	}

	if len(inn) == 10 {
		k := []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
		x := 0

		for i, j := range k {
			v, err := strconv.Atoi(string(inn[i]))
			if err != nil {
				return &ServiceError{Code: 3, Err: errors.New("inn must contains only digits ")}
			}

			x += j * v
		}

		xs := strconv.Itoa(x % 11)

		if xs[len(xs)-1] != inn[len(inn)-1] {
			return &ServiceError{Code: 3, Err: errors.New("inn checksum is invalid")}
		}

	}

	if len(inn) == 12 {
		k1 := []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
		x := 0

		for i, j := range k1 {
			v, err := strconv.Atoi(string(inn[i]))
			if err != nil {
				return &ServiceError{Code: 3, Err: errors.New("inn must contains only digits ")}
			}

			x += j * v
		}
		xs := strconv.Itoa(x % 11)

		k2 := []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
		y := 0

		for i, j := range k2 {
			v, err := strconv.Atoi(string(inn[i]))
			if err != nil {
				return &ServiceError{Code: 3, Err: errors.New("inn must contains only digits ")}
			}

			y += j * v
		}
		ys := strconv.Itoa(y % 11)

		if xs[len(xs)-1] != inn[10] && ys[len(ys)-1] != inn[11] {
			return &ServiceError{Code: 3, Err: errors.New("inn checksum is invalid")}
		}

	}

	return nil
}
