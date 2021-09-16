# gRPC (g Remote Procedure Call)
Framework desenvolvido pela google com o objetivo de facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida, leve, independente de linguagem

## Quanto utilizar

- Ideal para microserviços
- Mobile, Browsers e Backend
- Streaming bidirecional utilizando HTTP/2

## Linguagens com suporte

- gRPC - GO
- gRPC - JAVA
- gRPC - C
    - C++
    - Python
    - Ruby
    - Objective C
    - PHP
    - C#
    - Node.js
    - Dart
    - Kotlin / JVM

## Protocol Buffers vs Json

[Comparação](https://www.notion.so/b9f4c4f3f949491ba99868a96c26ce17)

## Protocol Buffers

### exemplo arquivo.proto:

```protobuf
syntax = "proto3";

message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;
}
```

## HTTP/2

- Nome original era SPDY
- Lançado em 2015
- Dados trafegados em binário e não texto (muito mais rápido)
- Utiliza a mesma conexão TCP para enviar e receber dados (buscar várias coisas com uma conexão aberta apenas)
- Server Push
- Headers são comprimidos (mais rápido)
- Gasta menos recursos de rede
- Processo é mais veloz

## gRPC

### API "unary"

- Forma padrão envia uma **request** e recebe uma **response**.

### API "Server streaming"

- Envia uma request e recebe várias respostas daquela requisição, ao invés de esperar todo o processamento para responder, é possível ir retornando as respostas aos poucos até fechar a conexão e terminar.

### API "Client streaming"

- Ao invés de enviar todos os dados de uma ves, vai enviando aos poucos na mesma conexão e quando tudo estiver pronto recebe a resposta do servidor.

**API "Bi directional streaming"**

- Ambos **Client streaming** e **Server streaming** trabalhando ao mesmo tempo.

## REST vc gRPC

[Comparação](https://www.notion.so/3fc1d93abee54fc1b05d8b79e36a9778)

## Links
- [**Protocol buffers**](https://developers.google.com/protocol-buffers/)
- [**gRPC**](https://grpc.io/)
- [**gRPC Web**](https://grpc.io/blog/state-of-grpc-web/)