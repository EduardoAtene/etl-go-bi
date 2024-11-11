# Documentação da API de Importação de Dados de Acidentes

Esta API foi criada para permitir o upload e processamento de arquivos CSV contendo dados de acidentes de trânsito. Existem dois tipos de arquivos que podem ser importados: um com dados da Polícia Rodoviária Federal (PRF) e outro com dados da Secretaria de Estado de Saúde de Minas Gerais (SES-MG).

## Visão Geral da API

- **Base URL**: `http://localhost:8081`
- **Versão da API**: `v1`
- **Formato de Dados**: `JSON`
- **Autenticação**: Não há autenticação necessária para os endpoints descritos nesta documentação.

## Endpoints de Upload

A API possui os seguintes endpoints para o upload dos arquivos:

### 1. Upload de Arquivo PRF

- **URL**: `/api/v1/upload/prf`
- **Método**: `POST`
- **Descrição**: Este endpoint permite o upload de um arquivo contendo dados de acidentes de trânsito coletados pela Polícia Rodoviária Federal (PRF).
- **Campo no FormData**: O arquivo deve ser enviado no campo `acidentes_prf`.
- **Formato do Arquivo**: O arquivo deve ser um CSV com os dados de acidentes da PRF.

#### Exemplo de Requisição

```bash
curl -X POST http://localhost:8081/api/v1/upload/prf \
  -F "acidentes_prf=@/caminho/para/o/arquivo/prf.csv"
Resposta Esperada:
Status HTTP 200: Se o upload e o processamento do arquivo forem bem-sucedidos.

Corpo da Resposta:

json
Copiar código
{
  "message": "Upload realizado com sucesso.",
  "status": "success"
}
Erro no Processamento:

Se ocorrer um erro durante o processamento do arquivo, o corpo da resposta será semelhante a:

json
Copiar código
{
  "message": "Erro ao processar o arquivo.",
  "status": "error"
}
2. Upload de Arquivo SES-MG
URL: /api/v1/upload/sesmg
Método: POST
Descrição: Este endpoint permite o upload de um arquivo contendo dados de acidentes de trânsito coletados pela Secretaria de Estado de Saúde de Minas Gerais (SES-MG).
Campo no FormData: O arquivo deve ser enviado no campo acidentes_sesmg.
Formato do Arquivo: O arquivo deve ser um CSV com os dados de acidentes da SES-MG.
Exemplo de Requisição
bash
Copiar código
curl -X POST http://localhost:8081/api/v1/upload/sesmg \
  -F "acidentes_sesmg=@/caminho/para/o/arquivo/sesmg.csv"
Resposta Esperada:
Status HTTP 200: Se o upload e o processamento do arquivo forem bem-sucedidos.

Corpo da Resposta:

json
Copiar código
{
  "message": "Upload realizado com sucesso.",
  "status": "success"
}
Erro no Processamento:

Se ocorrer um erro durante o processamento do arquivo, o corpo da resposta será semelhante a:

json
Copiar código
{
  "message": "Erro ao processar o arquivo.",
  "status": "error"
}
Detalhes sobre o Formato dos Arquivos CSV
Os arquivos CSV enviados devem seguir as especificações de cada fonte de dados.

1. Formato do Arquivo PRF
O arquivo de dados da PRF deve conter as seguintes colunas (exemplo):

Índice	Coluna	Descrição
0	data_acidente	Data do acidente no formato YYYY-MM-DD
1	hora_acidente	Hora do acidente no formato HH:mm:ss
2	municipio	Município onde o acidente ocorreu
3	dia_semana	Dia da semana do acidente
4	fase_dia	Período do dia (manhã, tarde, noite)
5	tipo_veiculo	Tipo do veículo envolvido
6	marca_veiculo	Marca do veículo
7	ano_fabricacao_veiculo	Ano de fabricação do veículo
8	km	Quilometragem da rodovia onde ocorreu o acidente
9	latitude	Latitude do local do acidente
10	longitude	Longitude do local do acidente
11	causa_acidente	Causa do acidente
Nota: Outras colunas podem ser incluídas dependendo do formato específico da PRF. O importante é garantir que as colunas essenciais estejam presentes.

2. Formato do Arquivo SES-MG
O arquivo de dados da SES-MG deve conter as seguintes colunas (exemplo):

Índice	Coluna	Descrição
0	data_acidente	Data do acidente no formato DD/MM/YYYY
1	municipio	Município onde o acidente ocorreu
2	tipo_envolvido	Tipo de envolvido (motorista, pedestre, etc.)
3	idade	Idade do envolvido
4	sexo	Sexo do envolvido
5	raca_cor	Raça/Cor do envolvido
6	causa_acidente	Causa do acidente
7	cid_causa_morte	CID da causa de morte
8	descricao_causa_morte	Descrição da causa da morte (se aplicável)
Nota: Assim como o arquivo da PRF, outras colunas podem ser adicionadas, dependendo do formato e dados disponíveis da SES-MG.

Exemplos de Erros Comuns
Arquivo com Formato Incorreto: Se o arquivo CSV não estiver no formato esperado (por exemplo, colunas ausentes ou delimitador incorreto), a resposta será:

json
Copiar código
{
  "message": "Formato de arquivo inválido.",
  "status": "error"
}
Tamanho do Arquivo: Se o arquivo for maior que o limite permitido, você receberá um erro de limite de tamanho de arquivo, como:

json
Copiar código
{
  "message": "Arquivo muito grande. O limite de tamanho é 10MB.",
  "status": "error"
}
Considerações
Limite de Tamanho do Arquivo: O tamanho máximo para o upload de um arquivo é de 10MB. Arquivos maiores não serão processados.
Encoding do Arquivo CSV: O arquivo CSV deve ser codificado em UTF-8 sem BOM (Byte Order Mark).
Verificação dos Dados: A API realiza uma verificação básica dos dados, como a validade das datas e a presença dos campos essenciais. Caso algum dado esteja incorreto ou ausente, ele será ignorado e a importação continuará com os dados válidos.
Performance: A API é otimizada para lidar com grandes volumes de dados, mas é importante garantir que o arquivo esteja no formato correto para evitar problemas durante o processamento.
Como Testar a API Localmente
Para testar a API localmente, siga os seguintes passos:

Inicie o servidor:

bash
Copiar código
go run main.go
Use curl ou Postman para enviar requisições de upload.

Verifique o status da resposta para garantir que o upload foi bem-sucedido.

A API estará disponível no endereço http://localhost:8081.

Conclusão
Esta API permite importar grandes volumes de dados de acidentes de forma eficiente e estruturada, facilitando o processamento e análise desses dados. Certifique-se de que os arquivos CSV sigam o formato especificado para evitar erros durante o upload.

markdown
Copiar código

### O que foi detalhado:
- Descrição completa de cada endpoint.
- Exemplo de uso de `curl` para o envio de arquivos.
- Explicação sobre o formato esperado dos arquivos CSV, com exemplos.
- Possíveis erros e respostas da API.
- Orientações para testar a API localmente.

Agora você tem uma documentação bem detalhada para incluir no seu `README.md`.