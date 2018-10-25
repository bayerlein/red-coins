# Rotas de usuário
<p>Existe apenas umas única rota para usuário: POST /v1/api/user/</p>
<p>Esta rota é publica, então não precisa estar autenticado para acessa-la</p>
<p>Exemplo de requisição:</p>
<p>Headers: Content-Type: application/x-www-form-urlencoded</p>
<p>Body: </p>
<p>email: email@email.com</p>
<p>password: senha</p>
<p>full_name: nome completo</p>
<p>date_of_birth: 2018-10-25</p>

# Rotas de bitcoin
<p>Contanto com os relatórios, existem 4 rotas:</p>
<p>Vender bitcoin: GET /v1/api/bitcoin/sell/{amount}</p>
<p>Comprar bitcoin: GET /v1/api/bitcoin/buy/{amount}</p>
<p>Amount é do tipo float</p>
<p>Exemplo de requisição:</p>
<p>Header: </p>
<p>Authorization: BEARER + <TOKEN></p>
<p>GET http://localhost:8080/v1/api/bitcoin/sell/918.2</p>
<p>GET http://localhost:8080/v1/api/bitcoin/buy/918.2</p>

<p>Relatório por usuário: GET /v1/api/bitcoin/reports/byuser/{user_id}</p>
<p>user_id é do tipo int</p>
<p>Relatório por data: GET /v1/api/bitcoin/reports/byday/{date}</p>
<p>date tem o seguinte formato: yyyy-MM-dd</p>
<p>Exemplo de requisição:</p>
<p>Authorization: BEARER + <TOKEN></p>
<p>GET http://localhost:8080/v1/api/bitcoin/reports/byuser/5</p>
<p>GET http://localhost:8080/v1/api/bitcoin/reports/byday/2018-10-22</p>

# BEARER Token
<p>O token pode ser gerado pelo site: https://jwt.io/ usando a chave My secret</p>
<p>No atual estado da aplicação, somente o user_id é usado. O seguinte header foi usado como teste:</p>
<p>BEARER </p>
<p>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJnaW92YW5uaWJheWVybGVpbjEy</p>
<p>QGdtYWlsLmNvbWp3dCIsImlhdCI6MTUxNjIzOTAyMn0.wT2QiwiW0LK4qo6IxeIKBIkIBFrW2ucUHNp0I8HwWfE</p>
<p>este toke representa o seguinte usuário:</p>
<p>"user_id": 5,</p>
<p>"sub": "1234567890",</p>
<p>"name": "John Doe",</p>
<p>"email": "giovannibayerlein12@gmail.comjwt",</p>