# Agente DNS Cloudflare

Aplicação para atualizar automaticamente o registro DNS no Cloudflare, apontando para o IP público atual da máquina onde o programa está sendo executado.

## Funcionalidades

- Obtém o IP público atual usando o serviço `ipify`.
- Atualiza um registro DNS A na Cloudflare com o IP público obtido.
- Usa a API da Cloudflare para interagir com os registros DNS.
- Requer autenticação via `X-Auth-Email` e `X-Auth-Key`.

## Pré-requisitos

- Conta na Cloudflare.
- Chave de API da Cloudflare.
- Configuração de um domínio e zona DNS na Cloudflare.
