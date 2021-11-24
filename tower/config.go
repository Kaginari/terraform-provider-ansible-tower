package tower

import (
	"strconv"
)

func getStateID(id int) string {
	return strconv.Itoa(id)
}

func decodeStateId(id string) (int, error) {
	result, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return result, nil
}
