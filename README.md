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

### Estrutura do projeto
- Cada pasta numerada representa um tema de estudo com exemplos autocontidos.
- Os exemplos que possuem dependências externas possuem seu próprio `go.mod`.
- Use o workspace (`go.work`) na raiz para navegar entre módulos com facilidade.

```sh
  # baixa dependências e mantém o workspace sincronizado
  go work sync

  # executa um exemplo específico
  go run ./4-Funcoes
```

### Mapa de estudos
- 1-Pacotes: organização de pacotes e módulos.
- 2-Variaveis: variáveis e tipos básicos.
- 3-TiposDeDados: tipos primitivos e conversões.
- 4-Funcoes: funções e padrões comuns.
- 5-Structs: structs e composição básica.
- 6-Pseudo-Heranca-Composicao: composição e embedding.
- 7-Ponteiros: ponteiros e passagem por referência.
- 8-Array-Slice: arrays, slices e matrizes.
- 9-Maps: mapas e operações.
- 10-Estruturas-De-Controle: if/else e controle de fluxo.
- 11-Switch: uso de switch.
- 12-Loops: laços for.
- 13-Funcoes Avancadas: defer, closures, init e panic/recover.
- 14-Metodos: métodos em structs e interfaces básicas.
- 15-Interfaces: interfaces e exemplos genéricos.
- 16-AplicacaoLinhaDeComando: app CLI com dependências externas.
- 17-Concorrencia: goroutines, channels e wait groups.
- 18-ConcurrencyPatterns: padrões de concorrência.
- 19-Testes: exemplos de testes e estrutura de pacotes.
- 20-Go1.22-Novidades: estudos sobre novidades do Go 1.22.

### Novidades do Go 1.22
- `for range` sobre inteiros (ex.: `for i := range 5`) para iterações simples.
- Variáveis de `range` passam a ser recriadas por iteração, evitando capturas erradas em closures.
- Confira exemplos na pasta `20-Go1.22-Novidades`.

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
### Packages
- Compilar os pacotes e dependencias
```sh
  go build
```
- Adicionar um pacote externo a aplicacao
```sh
  got get package_name
```
- Remover dependencias nao utilizadas
```sh
  go mod tidy
```

## Concorrência x Paralelismo
- Concorrência é sobre lidar com várias coisas ao mesmo tempo e paralelismo é sobre fazer várias coisas ao mesmo tempo. Concorrência é um conceito mais a nível de software e paralelismo mais a nível de hardware.
