# Go ile yazılmış, goroutine tabanlı hızlı TCP port tarayıcı.

- Concurrent tarama,
- JSON export,
- Banner Grabbing,
- Seçilen port aralığı

## Kurulum:

- git clone ...
- cd go-basic-port-scanner
- go build -o goscanner ./cmd/goscanner

## Örnek kullanım:

go run cmd\goscanner\main.go -host scanme.nmap.org -start 1 -end 100 -output results.json

### Çıktı - result.json:

[
{
"Port": 22,
"Open": true,
"Banner": "SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13"
},
{
"Port": 80,
"Open": true,
"Banner": ""
}
]

## Öğrenilen Konular

- **Goroutine & Channel:** 1000+ portu eş zamanlı taramak için worker pool pattern
- **sync.WaitGroup:** Tüm goroutine'lerin tamamlanmasını beklemek
- **net paketi:** TCP bağlantısı, timeout ve banner okuma
- **encoding/json:** Yapılandırılmış veri export
- **flag paketi:** CLI parametre yönetimi
