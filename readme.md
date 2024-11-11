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

----------

# Documentação do Processo ETL - Sistema de Análise de Acidentes

## Visão Geral
Este documento descreve a implementação de um sistema ETL (Extract, Transform, Load) desenvolvido em Go para processar dados de acidentes de trânsito provenientes de duas fontes distintas: Polícia Rodoviária Federal (PRF) e Secretaria de Estado de Saúde de Minas Gerais (SES-MG). O sistema foi projetado seguindo princípios de arquitetura limpa e padrões de projeto, visando manter a qualidade e escalabilidade do código.

## Arquitetura

### Estrutura do Projeto
O projeto segue uma arquitetura em camadas bem definidas:
- **Handlers**: Responsáveis pela interface HTTP
- **Usecases**: Contém a lógica de negócio
- **Repositories**: Gerenciam o acesso ao banco de dados
- **Entities**: Definem as estruturas de dados
- **Infrastructure**: Contém implementações específicas (parsers CSV, conexões de banco)

## Processo ETL

### 1. Extract (Extração)
A extração dos dados é realizada através de dois endpoints HTTP distintos, um para cada fonte de dados:

#### PRF Data
```go
func (h *UploadHandler) HandlePRF(c *gin.Context) {
    file, err := c.FormFile("acidentes_prf")
    // ... processamento do arquivo
}
```

#### SES-MG Data
```go
func (h *UploadHandler) HandleSESMG(c *gin.Context) {
    file, err := c.FormFile("acidentes_sesmg")
    // ... processamento do arquivo
}
```

### 2. Transform (Transformação)

#### Estrutura Intermediária
Foi criada uma estrutura intermediária `AcidenteTemp` para normalizar os dados de ambas as fontes:

```go
type AcidenteTemp struct {
    DataCompleta           time.Time
    Hora                   string
    PeriodoDia            string
    Municipio             string
    // ... outros campos
}
```

#### Processo de Transformação
O processo de transformação é realizado em duas etapas:

1. **Parsing CSV**
   - Implementação de parsers específicos para cada fonte
   - Tratamento de tipos de dados (conversão de strings para tipos apropriados)
   - Padronização de formatos de data e hora
   - Tratamento de valores nulos ou inválidos

2. **Normalização de Dados**
   - Padronização de nomenclaturas
   - Preenchimento de campos obrigatórios
   - Adequação ao modelo dimensional

### 3. Load (Carregamento)

#### Modelo Dimensional
O sistema utiliza um modelo dimensional (Star Schema) com as seguintes dimensões:
- Dim_Tempo
- Dim_Localizacao
- Dim_Veiculo
- Dim_Pessoa
- Dim_Condicoes
- Fato_Acidentes

#### Processo de Carga
O carregamento é realizado de forma transacional, seguindo a sequência:

1. **Inserção nas Dimensões**
   ```go
   // Exemplo para Dim_Tempo
   dimTempo := &entity.DimTempo{
       DataCompleta: acidente.DataCompleta,
       Hora:        acidente.Hora,
       // ... outros campos
   }
   idTempo, err := u.dimTempoRepo.Insert(dimTempo)
   ```

2. **Inserção na Tabela Fato**
   ```go
   fatoAcidente := &entity.FatoAcidentes{
       IDTempo:        idTempo,
       IDLocalizacao: idLocalizacao,
       // ... outros campos
   }
   _, err = u.fatoAcidentesRepo.Insert(fatoAcidente)
   ```

## Tratamento de Diferenças entre Fontes

### PRF
- Dados mais completos
- Informações detalhadas sobre veículos
- Dados sobre condições das vias
- Múltiplos tipos de vítimas (ilesos, feridos leves/graves, mortos)

### SES-MG
- Foco em óbitos
- Dados mais limitados
- Informações específicas sobre causa da morte (CID)
- Sempre conta como uma vítima fatal por registro

## Garantias de Qualidade

### Transações
- Implementação de transações SQL para garantir atomicidade
- Rollback automático em caso de falhas

### Validações
- Verificação de tipos de dados
- Tratamento de valores nulos
- Validação de datas

### Logs e Monitoramento
- Registro de erros durante o processamento
- Contagem de registros processados
- Identificação de registros problemáticos

## Considerações de Performance

### Otimizações Implementadas
1. Uso de prepared statements
2. Transações em lote
3. Índices nas tabelas dimensionais

### Pontos de Atenção
1. Volume de dados por carga
2. Tempo de processamento
3. Consumo de memória durante o parsing

## Conclusão
O sistema ETL implementado oferece uma solução robusta para a integração de dados de acidentes de diferentes fontes, mantendo a qualidade e consistência dos dados. A arquitetura modular permite fácil manutenção e extensão para novas fontes de dados no futuro.

