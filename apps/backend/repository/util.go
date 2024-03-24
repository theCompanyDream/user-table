package repository

import (
	"crypto/sha256"
    "encoding/hex"
    "encoding/json"
)

func hashObject(data interface{}) (*string, error) {
    // Serialize the struct to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    // Apply the SHA-256 hashing algorithm
    hash := sha256.New()
    hash.Write(jsonData)
	encodedHash := hex.EncodeToString(hash.Sum(nil))
    return &encodedHash, nil
}