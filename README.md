<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/2/2a/Mastercard-logo.svg" alt="Mastercard Logo" width="500">
</p>

<br>
<br>

# 💳 Mastercard ISO8583-1993 Parser (Go)
Parser developed in **Go** for reading and interpreting messages in the **ISO8583-1993** standard, with a focus on **IPM (Interchange File)** files from Mastercard.

---

## 📌 About the Project
This project parses ISO8583 messages found in Mastercard IPM files, converting raw data into organized structures and exporting the results to **CSV** and **JSON**.
It was developed to facilitate analyses, reconciliations, and integrations with financial systems that use the ISO8583-1993 standard in **EBCDIC (CP500)** encoding.

---

## ⚙️ Features
- Reading of Mastercard IPM files in binary EBCDIC format
- Complete parsing of ISO8583-1993 messages:
  - **MTI** (Message Type Indicator)
  - Primary and secondary **Bitmap**
  - Fixed and variable **Data Elements** (DE001–DE128)
- Customized parsing for the most complex fields:
  - **DE048** — Additional Data (format: type prefix + 2-byte key + 2-byte ASCII length + N-byte value)
  - **DE055** — ICC Data / EMV (exported in hex ASCII)
  - **PDS** (Private Data Subelements) within DE048
- Support for **CIC1**, **CIC2**, and **CIC3** processing cycles
- Parallel processing with **goroutines**
- Export to **CSV** (semicolon delimiter `;`) and ordered **JSON**
- Custom CSV headers with sorting (`ID` and `MTI` always first)

---

## 🧠 How It Works
The parser performs the following steps for each message:
1. Decodes the file from **EBCDIC to ASCII**
2. Identifies the **MTI** (4 bytes)
3. Reads the primary **Bitmap** (8 bytes) and, if necessary, the secondary one (additional 8 bytes)
4. For each active bit in the bitmap, interprets the corresponding **Data Element** based on its type and length as defined in the ISO8583-1993 specification
5. Applies specific parsers for **DE048** and **DE055**
6. Serializes the result to **CSV** and **JSON**

---

## 🛠️ Technologies Used
- **Go 1.22+**
- [`github.com/moov-io/iso8583`](https://github.com/moov-io/iso8583) — Main Parser
- [`github.com/iancoleman/orderedmap`](https://github.com/iancoleman/orderedmap) — Ordered JSON
- [`github.com/bytedance/sonic`](https://github.com/bytedance/sonic) — High-performance JSON serialization
- `encoding/csv` — CSV Generation

---

## 📄 License
This project is licensed under the [MIT](LICENSE) license.

---

## 👨🏻‍💻 Author
Developed by **Heitor Dairel**  
[GitHub](https://github.com/heitor-dairel) · [LinkedIn](https://linkedin.com/in/heitor-dairel)