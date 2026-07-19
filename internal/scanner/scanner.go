// paket adı yazılır.
package scanner

//import edilen paketler yazılır.

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Struct tanımlamaları yapılır. Struct tanımlamaları, tarayıcı tarafından kullanılacak verileri ve yapılandırmaları temsil eder. Örneğin, bir tarayıcı yapılandırması için bir struct tanımlanabilir.
type Scanner struct {
	Port   int
	Open   bool
	Banner string
}

// Fonksiyon tanımlamaları yapılır. Fonksiyonlar, tarayıcı tarafından gerçekleştirilecek işlemleri temsil eder. Örneğin, bir port tarama fonksiyonu veya bir banner alma fonksiyonu tanımlanabilir.

/*
@Param host: Hedef IP veya domain
@Param port: Tarama yapılacak port numarası
@Param timeout: Bağlantı zaman aşımı süresi

@Return Result: Tarama sonucu (port açık mı kapalı mı)
*/
func ScanPort(host string, port int, timeout time.Duration) Scanner {
	adress := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", adress, timeout)

	if err != nil {
		return Scanner{Port: port, Open: false, Banner: ""}
	}

	conn.SetReadDeadline(time.Now().Add(timeout))

	buf := make([]byte, 1024)
	banner := ""
	n, err := conn.Read(buf)
	if err == nil {
		banner = strings.TrimSpace(string(buf[:n]))
	}

	defer conn.Close()
	return Scanner{Port: port, Open: true, Banner: banner}

}
