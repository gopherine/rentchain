package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RentalAgreementKeyPrefix is the prefix to retrieve all RentalAgreement
	RentalAgreementKeyPrefix = "RentalAgreement/value/"
)

// RentalAgreementKey returns the store key to retrieve a RentalAgreement from the index fields
func RentalAgreementKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
