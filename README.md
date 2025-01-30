# CLI de Teste de Carga

Este projeto é uma ferramenta de Interface de Linha de Comando (CLI) escrita em Go para realizar testes de carga em um serviço web.

## Funcionalidades

- Realizar requisições HTTP para uma URL especificada.
- Controlar o número total de requisições e o nível de concorrência.
- Gerar um relatório com as seguintes informações:
  - Tempo total gasto no teste.
  - Número total de requisições realizadas.
  - Número de requisições com status HTTP 200.
  - Distribuição de outros códigos de status HTTP.

## Uso

### Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de requisições simultâneas.

### Exemplo

```sh
 docker build -t main . 
 docker run main --url=http://google.com --requests=1000 --concurrency=10
```