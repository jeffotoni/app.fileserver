# app.plataform.static.go

Este exemplo feito em Go é puramente didatico para melhor entendimento da funcionalidade de embuir em nosso binário arquivos.

Este repo é um exemplo de como compilar seu projeto web que server HTML dinâmico para dentro do seu binário. Isto mesmo que você leu embutir dentro do binário: HTML, CSS, JS e IMG.

O nosso diretório html é onde encontra-se uma página html um carrinho de compras, uma simples simulação para que possa entender o comportamento do nosso build usando Go.

Na versão go1.16beta este recurso já é nativo da linguagem, bacana né?! 😊 [embed go1.16](https://tip.golang.org/doc/go1.16#embed)


Para versões anteriores usamos uma lib externa para que possamos fazer nosso Embed.

Usamos para este projeto o [go static](github.com/rakyll/statik) para fazer o Embedded. 

Vamos primeiramente instalar o statik
```bash
$ go get github.com/rakyll/statik
```

Agora vamos transformar nosso diretório com os arquivos para o formato Go.
```bash
$ statik -src=html
```

Depois que fizemos isto agora é fazer o build do nosso projeto.
```bash
$ GOARCH=386 GOOS=linux go build -o web main.go

```

Agora é executar e conferir o resultado..!!!❤️


Na pasta go1.16 é onde irá encontrar os codigos nativos usando a lib nativa "embed".

Agora é ir  para cima!!!
