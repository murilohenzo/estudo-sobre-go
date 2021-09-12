### Anotacoes

### Resumo
* Linguagem imperativa
* Digitado estaticamente
* Tokens de sintaxe semelhantes a C (mas menos parênteses e nenhum ponto-e-vírgula) e a estrutura de Oberon-2
* Compila para código nativo (sem JVM)
* Sem classes, mas estruturas com métodos
* Interfaces
* Sem herança de implementação. No entanto, há [type embedding] (http://golang.org/doc/effective%5Fgo.html#embedding).
* As funções podem retornar vários valores
* Possui fechamentos
* Ponteiros, mas não aritmética de ponteiros
* Primitivos de simultaneidade integrados: Goroutines e canais


### Modules
- Cria um modulo dentro do go, fazendo um arquivo auxiliar ao principal ser capaz de ser importado no arquivo raiz.
```sh
  go mod init name_module
```

### Functions

- Function com a primeira letra minuscula significa que ela eh visivel somente dentro do pacote, ou um convencao de ser um metodo privado do pacote pertencente.

```sh
  func write() {
    fmt.Println("Escrevendo do pacote auxiliar")
  }
```

- Function com a primeira letra maiscula significa que ela eh visivel por outros pacotes, ou um convencao de ser um metodo publico do pacote pertencente.

```sh
  func Write() {
    fmt.Println("Escrevendo do pacote auxiliar")
  }
```
