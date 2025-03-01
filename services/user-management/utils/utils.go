package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
)

// GenerateAccountNumber creates a random 8-digit account number (UK format)
func GenerateAccountNumber() string {
	min := int64(10000000) // Minimum 8-digit number (to avoid leading zero)
	max := int64(99999999) // Maximum 8-digit number

	num, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return "00000000" // Fallback in case of error
	}

	return fmt.Sprintf("%08d", num.Int64()+min)
}

// GenerateSortCode creates a random 6-digit UK sort code (e.g., 12-34-56)
func GenerateSortCode() string {
	min := int64(100000) // Minimum 6-digit number
	max := int64(999999) // Maximum 6-digit number

	num, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return "000000" // Fallback in case of error
	}

	return fmt.Sprintf("%06d", num.Int64()+min)
}

// AES encryption key (should be stored securely, e.g., in environment variables)
var encryptionKey = []byte("my32characterrandomkeyforAES!") // 32 bytes key for AES-256

// EncryptData encrypts a given string using AES-GCM
func EncryptData(plainText string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptData decrypts a given AES-GCM encrypted string
func DecryptData(encryptedText string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	decryptedText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedText), nil
}

// GenerateIBAN creates a UK IBAN using a sort code and account number
func GenerateIBAN(sortCode, accountNumber string) string {
	// Define the Bank Identifier Code (BIC) – Example: `NWBK` for NatWest
	bankCode := "NWBK" // Modify as per your bank's code

	// The UK IBAN format is as follows:
	// IBAN = CountryCode (GB) + CheckDigits + BankCode + SortCode + AccountNumber
	iban := fmt.Sprintf("GB00%s%s%s%s", bankCode, sortCode, accountNumber)

	// Calculate Check Digits (GB00) for IBAN validation
	checkDigits := calculateCheckDigits(iban)
	iban = fmt.Sprintf("GB%s%s%s%s", checkDigits, bankCode, sortCode, accountNumber)

	// Return the full IBAN
	return strings.ToUpper(iban)
}

func calculateCheckDigits(iban string) string {
	// Move the first four characters to the end of the string
	iban = iban[4:] + iban[:4]

	// Replace letters with their numeric equivalent (A=10, B=11, ..., Z=35)
	ibanNumeric := ""
	for _, char := range iban {
		if char >= 'A' && char <= 'Z' {
			ibanNumeric += fmt.Sprintf("%d", int(char-'A')+10)
		} else {
			ibanNumeric += string(char)
		}
	}

	// Perform modulo 97 operation
	modulus, _ := new(big.Int).SetString(ibanNumeric, 10)
	checkDigits := modulus.Mod(modulus, big.NewInt(97)).Int64()

	// Ensure check digits are always 2 digits
	return fmt.Sprintf("%02d", 98-checkDigits)
}
