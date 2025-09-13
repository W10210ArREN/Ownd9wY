// 代码生成时间: 2025-09-13 18:10:07
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "log"
    "os"
)

// CipherText represents the encrypted data
type CipherText struct {
    Data string `json:"data"`
}

// PlainText represents the decrypted data
type PlainText struct {
    Data string `json:"data"`
}

// Config is the configuration for encryption and decryption
type Config struct {
    Key string
}

// NewConfig creates a new configuration with a given key
func NewConfig(key string) *Config {
    return &Config{Key: key}
}

// Encrypt encrypts plain text using AES encryption
func Encrypt(plainText string, config *Config) (*CipherText, error) {
    if len(config.Key) < 32 {
        return nil, fmt.Errorf("key must be at least 32 bytes long")
    }
    
    key := []byte(config.Key)
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)
    
    return &CipherText{Data: base64.StdEncoding.EncodeToString(ciphertext)}, nil
}

// Decrypt decrypts encrypted data using AES decryption
func Decrypt(cipherText string, config *Config) (*PlainText, error) {
    if len(config.Key) < 32 {
        return nil, fmt.Errorf("key must be at least 32 bytes long")
    }
    
    key := []byte(config.Key)
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    encryptedData, err := base64.StdEncoding.DecodeString(cipherText)
    if err != nil {
        return nil, err
    }
    
    ciphertext := encryptedData
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    
    plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    
    return &PlainText{Data: string(plainText)}, nil
}

func main() {
    config := NewConfig("your-32-byte-long-key")
    
    plainText := "Hello, World!"
    cipherText, err := Encrypt(plainText, config)
    if err != nil {
        log.Fatalf("Error encrypting data: %v", err)
    }
    fmt.Println("Encrypted data: ", cipherText.Data)
    
    decryptedText, err := Decrypt(cipherText.Data, config)
    if err != nil {
        log.Fatalf("Error decrypting data: %v", err)
    }
    fmt.Println("Decrypted data: ", decryptedText.Data)
}