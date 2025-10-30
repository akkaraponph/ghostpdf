## Lesson 05 — Fonts and UTF‑8

Built‑in core fonts: Helvetica, Times, Courier, ZapfDingbats. For UTF‑8 text, add a Unicode TrueType font.

### Quick UTF‑8 with DejaVu
```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()

// Add UTF‑8 TTF
pdf.AddUTF8Font("DejaVu", "", "font/DejaVuSansCondensed.ttf")
pdf.SetFont("DejaVu", "", 14)
pdf.Cell(0, 10, "English: Hello")
pdf.Ln(8)
pdf.Cell(0, 10, "ไทย: สวัสดี")
pdf.Ln(8)
pdf.Cell(0, 10, "Русский: Привет")

_ = pdf.OutputFileAndClose("lesson05.pdf")
```

### Custom fonts
- Use `AddUTF8Font` for TTF/OTF files directly.
- Or pre‑generate JSON/Z font metrics with the included makefont utility for faster load.


