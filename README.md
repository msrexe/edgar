# edgar

`edgar` is a command-line interface (CLI) application developed in Go using the Cobra library. It provides AES256 encryption and decryption functionality for text or file inputs.

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/msrexe/edgar.git
   ```
2. Navigate to the project directory:
   ```
   cd edgar
   ```
3. Build the binary:
   ```
   go build -o edgar cmd/edgar/main.go
   ```
4. Run the application:
   ```
   ./edgar [command]
   ```

## Usage

- Encrypt a text:
   ```
   edgar enc 'your-text' -key enc_key.txt
   ```
- Encrypt the contents of a file:
   ```
   edgar enc -f plain.txt -key enc_key.txt > encryptedText.txt
   ```
- Decrypt the encrypted text:
   ```
   edgar dec 'encrypted-text' -key enc_key.txt
   ```
- Decrypt the contents of a file:
   ```
   edgar dec -f encryptedText.txt -key enc_key.txt > decryptedText.txt
   ```

Make sure to replace `enc_key.txt` with your own encryption key file.

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
