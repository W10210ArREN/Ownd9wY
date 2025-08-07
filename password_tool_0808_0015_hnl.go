// 代码生成时间: 2025-08-08 00:15:22
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
    "log"
    "os"
    
    "github.com/revel/revel"
)

// EncryptionKey is the key used for encryption/decryption.
// It should be a 32-byte key for AES-256.
var EncryptionKey = []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f")

// Encrypt encrypts the given plain text using AES-256.
func Encrypt(plainText string) (string, error) {
    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        log.Printf("Error creating cipher: %v", err)
        return "", err
    }
    
    plainTextBytes := []byte(plainText)
    cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        log.Printf("Error generating IV: %v", err)
        return "", err
    }
    
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText[aes.BlockSize:], plainTextBytes)
    
    return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts the given cipher text using AES-256.
func Decrypt(cipherText string) (string, error) {
    cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
    if err != nil {
        log.Printf("Error decoding cipher text: %v", err)
        return "", err
    }
    
    block, err := aes.NewCipher(EncryptionKey)
    if err != nil {
        log.Printf("Error creating cipher: %v", err)
        return "", err
    }
    
    if len(cipherTextBytes) < aes.BlockSize {
        return "", errors.New("cipher text too short")
    }
    
    iv := cipherTextBytes[:aes.BlockSize]
    cipherTextBytes = cipherTextBytes[aes.BlockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    
    // PKCS#7 padding
    padding := cipherTextBytes[len(cipherTextBytes)-1]
    cipherTextBytes = cipherTextBytes[:len(cipherTextBytes)-int(padding)]
    mode.CryptBlocks(cipherTextBytes, cipherTextBytes)
    
    return string(cipherTextBytes), nil
}

func main() {
    revel.OnAppStart(func() {
        // Initialize Revel
        revel.InitRoutes(revel.DefaultRouteRegex)
    })
    
    // Start the server
    revel.RunProd()
}
