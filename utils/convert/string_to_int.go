package convert

import "strconv"

func StrToInt(str string) (int, error) {
	int, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int, nil
}