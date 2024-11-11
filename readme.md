# Documentação da API de Importação de Dados de Acidentes

Esta API foi criada para permitir o upload e processamento de arquivos CSV contendo dados de acidentes de trânsito. Existem dois tipos de arquivos que podem ser importados: um com dados da Polícia Rodoviária Federal (PRF) e outro com dados da Secretaria de Estado de Saúde de Minas Gerais (SES-MG).

## Visão Geral da API

- **Base URL**: `http://localhost:8081`
- **Versão da API**: `v1`
- **Formato de Dados**: `JSON`
- **Autenticação**: Não há autenticação necessária para os endpoints descritos nesta documentação.

## Endpoints de Upload

### 1. Upload de Arquivo PRF

- **URL**: `/api/v1/upload/prf`
- **Método**: `POST`
- **Descrição**: Este endpoint permite o upload de um arquivo contendo dados de acidentes de trânsito coletados pela Polícia Rodoviária Federal (PRF).
- **Campo no FormData**: O arquivo deve ser enviado no campo `acidentes_prf`.
- **Formato do Arquivo**: O arquivo deve ser um CSV com os dados de acidentes da PRF.

#### Exemplo de Requisição

Para realizar o upload de um arquivo PRF, você pode usar o seguinte exemplo de requisição com `curl`:

```curl -X POST http://localhost:8081/api/v1/upload/prf -F "acidentes_prf=@path_to_your_file.csv"```

#### Resposta Esperada

- **Status HTTP 200**: Se o upload e o processamento do arquivo forem bem-sucedidos.

```{ "message": "Upload realizado com sucesso.", "status": "success" }```

### 2. Upload de Arquivo SES-MG

- **URL**: `/api/v1/upload/sesmg`
- **Método**: `POST`
- **Descrição**: Este endpoint permite o upload de um arquivo contendo dados de acidentes de trânsito coletados pela Secretaria de Estado de Saúde de Minas Gerais (SES-MG).
- **Campo no FormData**: O arquivo deve ser enviado no campo `acidentes_sesmg`.
- **Formato do Arquivo**: O arquivo deve ser um CSV com os dados de acidentes da SES-MG.

#### Exemplo de Requisição

Para realizar o upload de um arquivo SES-MG, você pode usar o seguinte exemplo de requisição com `curl`:

```curl -X POST http://localhost:8081/api/v1/upload/sesmg -F "acidentes_sesmg=@path_to_your_file.csv"```


#### Resposta Esperada

- **Status HTTP 200**: Se o upload e o processamento do arquivo forem bem-sucedidos.

```{ "message": "Upload realizado com sucesso.", "status": "success" }```

## Observações

- O arquivo enviado para cada endpoint deve estar no formato CSV.
- Certifique-se de que o nome do campo no FormData esteja correto: para o PRF, use `acidentes_prf` e para o SES-MG, use `acidentes_sesmg`.
