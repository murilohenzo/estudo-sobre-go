// concorrencia != paralelismo
package main

import (
	"fmt"
	"time"
)

func main() {
	// Uma goroutine é uma thread leve gerenciada pelo runtime Go.
	//Goroutines são executados no mesmo espaço de endereço,
	//portanto, o acesso à memória compartilhada deve ser sincronizado.
	go write("Hello World!")
	write("Programming go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
