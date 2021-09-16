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

| Protocol Buffers | JSON |
| :---: | :---: |
| Arquivos binários | JSON |
| Processamento mais leve | - |
| Gasta menos recursos de rede | - |
| Processo mais veloz | - |

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

| REST | gRPC |
| :---: | :---: |
| Texto / JSON | Protocol Buffers (arquivos binários) |
| Unidirecional | Bidirecional e Assíncrono |
| Alta latência | Baixa latência |
| Sem contrato (maior chance de erros) | Contrato definido (.proto) |
| Sem suporte a streaming (Request / Response) | Suporte a streaming |
| Design pré-definido | Design é livre |
| Bibliotecas de terceiro | Geração de código |

## Links
- [**Protocol buffers**](https://developers.google.com/protocol-buffers/)
- [**gRPC**](https://grpc.io/)
- [**gRPC Web**](https://grpc.io/blog/state-of-grpc-web/)
