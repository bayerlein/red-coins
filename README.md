# Rotas de usuário
Existe apenas umas única rota para usuário: POST /v1/api/user/
Esta rota é publica, então não precisa estar autenticado para acessa-la
Exemplo de requisição:
Headers: Content-Type: application/x-www-form-urlencoded
Body: 
email: email@email.com
password: senha
full_name: nome completo
date_of_birth: 2018-10-25

# Rotas de bitcoin
Contanto com os relatórios, existem 4 rotas:
Vender bitcoin: GET /v1/api/bitcoin/sell/{amount}
Comprar bitcoin: GET /v1/api/bitcoin/buy/{amount}
Amount é do tipo float
Exemplo de requisição:
Header: 
Authorization: BEARER + <TOKEN>
GET http://localhost:8080/v1/api/bitcoin/sell/918.2
GET http://localhost:8080/v1/api/bitcoin/buy/918.2

Relatório por usuário: GET /v1/api/bitcoin/reports/byuser/{user_id}
user_id é do tipo int
Relatório por data: GET /v1/api/bitcoin/reports/byday/{date}
date tem o seguinte formato: yyyy-MM-dd
Exemplo de requisição:
Authorization: BEARER + <TOKEN>
GET http://localhost:8080/v1/api/bitcoin/reports/byuser/5
GET http://localhost:8080/v1/api/bitcoin/reports/byday/2018-10-22

# BEARER Token
O token pode ser gerado pelo site: https://jwt.io/ usando a chave My secret
No atual estado da aplicação, somente o user_id é usado. O seguinte header foi usado como teste:
BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJnaW92YW5uaWJheWVybGVpbjEyQGdtYWlsLmNvbWp3dCIsImlhdCI6MTUxNjIzOTAyMn0.wT2QiwiW0LK4qo6IxeIKBIkIBFrW2ucUHNp0I8HwWfE
este toke representa o seguinte usuário:
"user_id": 5,
"sub": "1234567890",
"name": "John Doe",
"email": "giovannibayerlein12@gmail.comjwt",