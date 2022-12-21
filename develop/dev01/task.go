package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===
Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и go lint.
*/

func GetTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {

	time1, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to NTP-server failed: %s", err)
		os.Exit(0)
	}

	fmt.Printf("NTP-server time:\t%s", time1)
}
