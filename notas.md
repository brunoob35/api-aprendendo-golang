Objetivos entrega 3:
1. Completar o Uders Model com os métodos que faltam
2. Implementar funcionalidades do CRUD de Users
3. Criar um script de utilidades para validação de CPF

### Completando os metodos do Users Model:
1. Completar Prep()
    - já está praticamente pronto já que temos as funções que ele precisa chamar já preparadas

2. Completar validatee()
    - verificar strings em branco
    - verificar se o email existe de fato
    - validar se o CPF é válido
    - verifica se a senha está em branco caso passo for cadastro

3. Completar format()
    - Trim spaces
    - Adequa todo para lower case
    - no futuro vamos criptografar a senha aqui também

4. Validador de CPF
   Como fazemos uma validação de CPF? https://www.macoratti.net/alg_cpf.htm
- Crimaos uma nova pasta na raiz do projeto "utils"
- Fazemos a verificação com os tres passos

### Create()
- Lê o body
- Descomprimi o conteudo
- Prepara o caminho para o Repo
- Responde

### Criar um pacote de respostas para padronizar (basicamente a função que já tinhamos dentro do books_handler)
- Criamos dois metodos diferentes, uma para as respostas e outro para os erros

### Criando o Banco
- Criamos a pasta sql
- Criamos o ddl.sql
- Colocamos o script do ddl ali


### Preparando ambiente de configuração da aplicação
- Criamos o novo modulo "config"
- Criamos o .env no config
- Criamos a função LoadEnv() em config
- Editamos o main para passar a chamar config.LoadEnv()
- Criamos o gitignore e adicionamos o env nele para que não suba ao github
```
DB_USER=admin
DB_PASSWORD=sua_senha_aqui
DB_ADDR=treehouse-db.ctuc44aqwph2.us-east-2.rds.amazonaws.com:3306
DB_DATABASE=treehouse-db
```
- cria o .gitignore na raiz do projeto
```
.env
```

### Criar o pacote persistency
- Criamos a nova pasta persistency a partir da raiz do projeto
- Criamos o arquivo database dentro de persistency
- usamos o pacote database/sql para configurar o banco na função Connect()
- testamos a conexão com o banco

### Criando a estrutura do Repositório
- Dentro do repository/users vamos criar o type UsersRepo
- Criar ao metodo UsersNewRepo
- Editar o metodo Create dentro do repositorio
- Voltamos no CRUD para conectar nosso controller com o repo

### Desafio 4: Voltar no CRUD e fazer os ajustes fazendo as chamadas para o banco depois chamando o repositório para salvar os dados
- Senhas precisam ser salvas criptografadas!!!
- vamos criar outra camada (security) para colocar as funções de criptografia da senha
- essa camada deve ter dois metodos: um para criptografar a senha e salvar o Hash no banco
- outro para ler um hash vindo do banco e fazer o processo de comparação com o conteudo vindo na requeisição para validar a entrada de um novo usuário
