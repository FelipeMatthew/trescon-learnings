package services

import (
	"fmt"
	"strconv"
	"time"
)

func TimestampToDateTime(currentTimestamp string) (string, error) {

	tsFloat, err := strconv.ParseFloat(currentTimestamp, 64)
	if err != nil {
		return "", fmt.Errorf("erro ao converter timestamp: %v", err)
	}

	seconds := int64(tsFloat)
	nanoseconds := int64((tsFloat - float64(seconds)) * 1e9)

	t := time.Unix(seconds, nanoseconds)

	// Caso queria mudar formato
	return t.Format("02/01 15:04"), nil
}
