package golife

import (
	"strings"
	"errors"
	"strconv"
)

func StringToSeed(s string, w, h int) ([]bool, error) {
	seed := strings.Split(s, " ")
	if len(seed) != w*h {
		return nil, errors.New("bad seed langth ")
	}
	seedArray := make([]bool, 0, h*w)
	for _, e := range seed {
		num, err := strconv.Atoi(e)
		if err != nil {
			return nil, errors.New("bad seed content ")
		}
		seedArray = append(seedArray, num == 1)
	}
	return seedArray, nil
}
