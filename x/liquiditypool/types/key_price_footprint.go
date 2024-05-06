package types

import (
	"encoding/binary"
	"time"
)

var _ binary.ByteOrder

const (
	// PriceFootprintKeyPrefix is the prefix to retrieve all PriceFootprint
	PriceFootprintKeyPrefix = "PriceFootprint/value/"
)

// PriceFootprintKey returns the store key to retrieve a PriceFootprint from the index fields
func PriceFootprintKey(
	baseDenom string,
	quoteDenom string,
	timestamp time.Time,
) []byte {
	var key []byte

	baseDenomBytes := []byte(baseDenom)
	quoteDenomBytes := []byte(quoteDenom)
	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(timestamp.Unix()))

	key = append(key, baseDenomBytes...)
	key = append(key, []byte("/")...)
	key = append(key, quoteDenomBytes...)
	key = append(key, []byte("/")...)
	key = append(key, timestampBytes...)

	return key
}

func PriceFootprintIterationPrefix(
	baseDenom string,
	quoteDenom string,
) []byte {
	var key []byte

	baseDenomBytes := []byte(baseDenom)
	quoteDenomBytes := []byte(quoteDenom)

	key = append(key, baseDenomBytes...)
	key = append(key, []byte("/")...)
	key = append(key, quoteDenomBytes...)

	return key
}