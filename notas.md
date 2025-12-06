Objetivos entrega 4:
- Criptografia e salvamento das senhas criptgrafadas

### Criando o pacote Security
1. Função Hash
2. Função ValidatePassword
3. Adicionada a chamada para Hash() dentro da de models/users.format()
4. Confirmar se a senha salva no DB agora está "hasheada"

### Login
1. Como agora teremos rotas que exigem autenticação do Usuário, queremos marcar rotas que deverão passar por uma verificação de autenticação
2. Vamos atualizar nossa struct de rotas e acrescentar o campo de Auth
3. Vamos atualizar nossas rotas já criadas e adicionar o campo Auth em cada uma
4. Vamos criar uma rota para Login também com um novo arquivo
5. Vamos criar a função de Login no controller

### Controller Login
1. Processo padrão que já fazemos, lê o body
2. Unmarshal do Body em uma var do tipo Users
3. Abrimos conexão com o banco (lembra do defer na conexão também)
4. Chamamos um novo repositório de Usuários

### Repositório de Users
1. Aqui vamos criar um novo metodo de consulta no banco que vai retornar só o Id, email a senha do User
2. Pq no repo Users e não criar um novo repo de login?
3. Criamos o metodo FetchByEmail que faz um Select no Banco e reotrna apenas os dados que queremos

### Coltando para o Controller Login
1. Agora com os dados salvos no DB e os dados da requisição em mãos
2. Podemos chamar o security.ValidatePassword()
3. Você está logado
    ```w.Write([]byte("User logged in"))```
4. Vamos testar!

### Token de autenticação JWT (JSON Web Token)
1. Criamos o pacote token e o arquivo token.go
2. Importamos a lib JWT-go
3. Criamos uma função GenerateToken() que receberá um Id de usuário
```
    // para teste geramos com um secret simples
   permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // tempo para expirar o token
	permissions["userId"] = userID

	// jwt.SigningMethodHS256 é um de muitos diferentes métodos para assinatura do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("secret"))
```
```
    // em Login() somente para vermos o token gerado
    token, _ := auth.GenerateToken(uint64(userSalvoemDB.ID))
	fmt.Println(token)
```
4. Agora que foi testada a função que gera Token, vamos criar um Secret melhor!
5. Para este metodo de assinatura é recomendado uma chave de 64bytes

### Criando o Secrets
1. Conseguimos criar essa chave usando pacotes do próprio Go
2. Dentro do pacote main vamos criar uma função init() que vai rodar antes mesmo da main()
```
func init() {
	chave := make([]byte, 64)
	fmt.Println(chave)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(chave))

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
```
3. Essa função vai gerar um slice de 64bytes que vamos salvar no .env em SECRET_KEY
4. Depois disso podemos deletar a função Init()

### Middlewares - Camada que fica entre a requisição e a resposta
1. Primeiro criamos o pacote e a função Autenticate, recebendo os parametros w e r
2. Atualizamos nosso roteador para que no loop que carrega nossas rotas, seja feita também a chamada do middleware
3. Criamos a função Logger também para criar um registro das chamadas dentro da API
4. Esse logger chamamos antes do Autenticate e também dentro das rotas que não precisam de autenticação

### Autenticação
1. Agora que temos o middleware sendo executado em todas as rotas, vamos continuar o processo de validação do token
2. vamos criar o extractToken() que vai seprar o token do seu container
3. ValidateToken() vai validar este token extraido para autenticação de fato
4. Antes de terminar a validação vou verificar se a chave de autenticação é a mesma do token recebido, por segurança

### Desafio 5 Biblioteca Funcional:
- Agora que temos Users e um Login funcional, o desafio é criar as funcionalidades de CRUD dos livros
- Queremos que quando o usuário estiver logado ele seja capaz de pesquisar por um livro, que será vinculado aquele usuario
- Assim poderemos ver a biblioteca particular daquele usuario
- Para isso, além da pesquisa de livros, sejamos capazes de vincular o livro ao user
- Editar livro do user
- Deletar livro da biblioteca deste user
- Visualizar todos os livros da biblioteca deste User
- Você terá que modelar a camada de Livros
- E também deverá criar tabelas no banco que se relacionem a tabela de users, a fim de vincular estes livros salvos a algum user

