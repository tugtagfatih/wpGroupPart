# 📞 WhatsApp Phone Number Extractor

This simple Go application extracts Turkish phone numbers from an exported WhatsApp `.html` chat file and saves the **unique numbers** to a text file.

---

## ✨ Features

- Extracts Turkish-format phone numbers: `+90 5XX XXX XX XX`
- Removes duplicate entries
- Saves the result to a clean output file (`numbers.txt`)
- Efficient scanning with large buffer support (10MB)

---

## 📂 Input Format

Export a WhatsApp chat **without media** and select the `.html` format if possible. The file should contain phone numbers in this format:

```
+90 532 123 45 67
```

---

## 🚀 How It Works

1. Opens `whatsapp.html`.
2. Scans the content for Turkish phone numbers using regex.
3. Saves the matches to `temp_numbers.txt`.
4. Removes duplicates and writes to `numbers.txt`.
5. Deletes the temporary file.

---

## 🛠️ Installation

1. Make sure [Go](https://golang.org/dl/) is installed.
2. Clone this repository:

```bash
git clone https://github.com/yourusername/whatsapp-number-extractor.git
cd whatsapp-number-extractor
```

3. Build the application:

```bash
go build -o extractor main.go
```

---

## 📦 Usage

Place your `whatsapp.html` file in the same directory and run:

```bash
./extractor
```

After running, you'll find the output in:

```
numbers.txt
```

---

## 📄 Example Output

```
+90 532 123 45 67
+90 544 765 43 21
+90 501 234 56 78
```

---

## 🤝 Contributions

Contributions, issues, and feature requests are welcome! Feel free to open an issue or submit a pull request.

---

## 👨‍💻 Author

Developed by [tugtagfatih](https://github.com/tugtagfatih)
