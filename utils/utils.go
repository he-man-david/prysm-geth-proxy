package utils

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"hash/crc32"
	log2 "log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

var (
	datadirJWTKey          = "jwtsecret" 
)

// obtainJWTSecret loads the jwt-secret, taken from GETH
func ObtainJWTSecret(cliParam string) ([]byte, error) {
	fileName := cliParam
	if len(fileName) == 0 {
		// no path provided, use default
		log2.Fatal("ERROR: filename to JWT secret must be provided!! \n")
	}
	// try reading from file
	if data, err := os.ReadFile(fileName); err == nil {
		jwtSecret := common.FromHex(strings.TrimSpace(string(data)))
		if len(jwtSecret) == 32 {
			log.Info("Loaded JWT secret file", "path", fileName, "crc32", fmt.Sprintf("%#x", crc32.ChecksumIEEE(jwtSecret)))
			return jwtSecret, nil
		}
		log.Error("Invalid JWT secret", "path", fileName, "length", len(jwtSecret))
		return nil, errors.New("invalid JWT secret")
	}
	// Need to generate one
	jwtSecret := make([]byte, 32)
	crand.Read(jwtSecret)
	// if we're in --dev mode, don't bother saving, just show it
	if fileName == "" {
		log.Info("Generated ephemeral JWT secret", "secret", hexutil.Encode(jwtSecret))
		return jwtSecret, nil
	}
	if err := os.WriteFile(fileName, []byte(hexutil.Encode(jwtSecret)), 0600); err != nil {
		return nil, err
	}
	log.Info("Generated JWT secret", "path", fileName)
	return jwtSecret, nil
}