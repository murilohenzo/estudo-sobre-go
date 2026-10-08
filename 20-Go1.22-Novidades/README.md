# Novidades do Go 1.22

Este diretório reúne exemplos de estudo das mudanças introduzidas no Go 1.22.

## Range sobre inteiros

O `range` pode iterar diretamente sobre um inteiro, iniciando em `0` e indo até `n-1`.

```go
for i := range 5 {
	fmt.Println(i)
}
```

## Variáveis de range por iteração

As variáveis do `range` passam a ser recriadas a cada iteração, o que evita capturas
incorretas em closures quando usamos goroutines.

```go
for i := range 3 {
	go func() {
		fmt.Println(i)
	}()
}
```

> Observação: a ordem de impressão não é garantida por causa da concorrência,
> mas cada goroutine imprime o valor correto de `i`.
