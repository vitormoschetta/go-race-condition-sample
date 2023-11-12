# Race Condition Sample

A ideia é mostrar como um dado pode ser alterado por mais de uma thread ao mesmo tempo, causando inconsistência de dados.

## Bombardier

Para simular isso vamos utilizar o [bombardier](https://github.com/codesenberg/bombardier):

```bash
go install github.com/codesenberg/bombardier@latest
```

Mover o binário para o diretório de binários do sistema operacional. Exemplo:
```bash
sudo mv $(go env GOPATH)/bin/bombardier /usr/local/bin
```

## Subindo a aplicação

```bash
go run main.go
```

## Executando o bombardier

```bash
bombardier -n 50 -m GET -l http://localhost:8080/api/execute1
```

Legenda de parâmetros:
- `-n` número de requisições
- `-m` método HTTP

Veja a inconsistência de dados na propriedade `count` na saída de logs no terminal.  
`count` deveria ser um número constante, mas como a propriedade é alterada por mais de uma thread ao mesmo tempo, o valor é inconsistente.

## Executando o bombardier com mutex

```bash
bombardier -n 50 -m GET -l http://localhost:8080/api/execute2
```

Veja que agora o valor de `count` é consistente. 

**Nota**: O mutex mantém o dado consistente, pois ele bloqueia a thread que está alterando o dado até que a thread termine de alterar o dado.  
Perceba que isso pode impactar no desempenho da aplicação. Faça o seguinte teste:
    
```bash
bombardier -n 10000 -m GET -l http://localhost:8080/api/execute1
```

```bash
bombardier -n 10000 -m GET -l http://localhost:8080/api/execute2
```

Veja o benchmark de cada um dos testes acima e compare os resultados.

## Referências
https://github.com/codesenberg/bombardier  
https://pkg.go.dev/github.com/codesenberg/bombardier


