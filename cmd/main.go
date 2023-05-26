package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"edgar"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "edgar",
		Short: "AES256 encryption and decryption CLI",
	}

	var encCmd = &cobra.Command{
		Use:   "enc",
		Short: "Encrypt the provided text",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().Changed("file") {
				encryptFile(cmd, args)
			} else {
				if len(args) != 1 {
					fmt.Println("Must provide text to encrypt")
					return
				}

				encryptText(cmd, args)
			}
		},
	}

	var decCmd = &cobra.Command{
		Use:   "dec",
		Short: "Decrypt the provided text",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().Changed("file") {
				decryptFile(cmd, args)
			} else {
				if len(args) != 1 {
					fmt.Println("Must provide text to decrypt")
					return
				}

				decryptText(cmd, args)
			}
		},
	}

	encCmd.Flags().StringP("key", "k", "key.txt", "File containing the encryption key")
	decCmd.Flags().StringP("key", "k", "key.txt", "File containing the decryption key")
	encCmd.Flags().StringP("file", "f", "", "File to encrypt")
	decCmd.Flags().StringP("file", "f", "", "File to decrypt")

	rootCmd.AddCommand(encCmd)
	rootCmd.AddCommand(decCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func encryptText(cmd *cobra.Command, args []string) {
	text := args[0]
	keyFile, _ := cmd.Flags().GetString("key")

	fmt.Printf("keyFile: %s\n", keyFile)
	key, err := loadKeyFile(keyFile)
	if err != nil {
		fmt.Println("Error loading encryption key:", err)
		return
	}

	encryptedText, err := edgar.Encrypt([]byte(text), key)
	if err != nil {
		fmt.Println("Error encrypting text:", err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(encryptedText))
}

func encryptFile(cmd *cobra.Command, args []string) {
	fileName, _ := cmd.Flags().GetString("file")
	keyFile, _ := cmd.Flags().GetString("key")

	key, err := loadKeyFile(keyFile)
	if err != nil {
		fmt.Println("Error loading encryption key:", err)
		return
	}

	plainText, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	encryptedText, err := edgar.Encrypt(plainText, key)
	if err != nil {
		fmt.Println("Error encrypting file contents:", err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(encryptedText))
}

func decryptText(cmd *cobra.Command, args []string) {
	text := args[0]
	keyFile, _ := cmd.Flags().GetString("key")

	key, err := loadKeyFile(keyFile)
	if err != nil {
		fmt.Println("Error loading decryption key:", err)
		return
	}

	decodedText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println("Error decoding text:", err)
		return
	}

	decryptedText, err := edgar.Decrypt(decodedText, key)
	if err != nil {
		fmt.Println("Error decrypting text:", err)
		return
	}

	fmt.Println(string(decryptedText))
}

func decryptFile(cmd *cobra.Command, args []string) {
	fileName, _ := cmd.Flags().GetString("file")
	keyFile, _ := cmd.Flags().GetString("key")

	key, err := loadKeyFile(keyFile)
	if err != nil {
		fmt.Println("Error loading decryption key:", err)
		return
	}

	encryptedText, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	decodedText, err := base64.StdEncoding.DecodeString(string(encryptedText))
	if err != nil {
		fmt.Println("Error decoding text:", err)
		return
	}

	decryptedText, err := edgar.Decrypt(decodedText, key)
	if err != nil {
		fmt.Println("Error decrypting file contents:", err)
		return
	}

	fmt.Println(string(decryptedText))
}

func loadKeyFile(keyFile string) ([]byte, error) {
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return key, nil
}
