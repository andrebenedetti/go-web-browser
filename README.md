### My implementation of https://browser.engineering/http.html in Go

#### Chapter 1 
- [x] Request an url (not using net/http to closely follow the book :)
- [x] Parse status code
- [x] Parse headers
- [x] Render HTML texts in the console (currently text inside <script> and <style> tags are rendered too)

- [ ] alternate encoding (non UTF-8)
- [ ] encoding/chunking
- [ ] HTTP/1.1
- [ ] File URLs (file://)
- [ ] Data URLs (data://)
- [ ] Body tag (only show text inside <body></body>
- [ ] Entities (&lt;) (&gt;)
- [ ] view-source scheme
- [ ] Compresion
- [ ] Redirects
- [ ] Caching
