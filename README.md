SIGAA Mock

Resumo

Este repositório contém um mock simples de autenticação inspirado no SIGAA. Ele expõe um endpoint `/login` que valida username+senha (senha enviada em base64 via query string) e retorna um payload JSON específico para 4 cenários mockados.

Rodando localmente

1) No Windows (PowerShell), no diretório do projeto:

```powershell
Set-Location -LiteralPath 'D:\Projects\Estágio IF\Backend\sigaa-mock'
go run .
```

O servidor escuta em http://localhost:8080

Endpoint

GET /login

Query params:
- username: `cenario1` | `cenario2` | `cenario3` | `cenario4` (ou `cenarioX@mail.com` — o mock aceita ambos)
- password: senha codificada em base64 (ex.: `Test1234!` -> `VGVzdDEyMzQh`)

Header:
- Authorization: Bearer apitoken123

Comportamento

- Se Authorization incorreto: 401 Unauthorized com corpo `unauthorized`.
- Se username ou senha inválidos: 401 com mensagem `invalid credentials for <user>/<pass>`.
- Se credenciais válidas: 201 Created com payload JSON do usuário.

Contas mock (username / senha em texto)

- cenario1 / Test1234!  => email: cenario1@mail.com, cpf: 379.137.150-90, perfil: ["docente"]
- cenario2 / Test5678!  => email: cenario2@mail.com, cpf: 343.028.030-37, perfil: ["discente"]
- cenario3 / Test9012!  => email: cenario3@mail.com, cpf: 205.188.000-08, perfil: ["docente"]
- cenario4 / Test1234!  => email: cenario4@mail.com, cpf: 603.559.990-72, perfil: ["discente"]

Exemplos

PowerShell (ex.: cenario1):

```powershell
$u = "cenario1"
$p = [System.Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("Test1234!"))
Invoke-WebRequest -Uri "http://localhost:8080/login?username=$u&password=$p" -Headers @{ Authorization = "Bearer apitoken123" } -UseBasicParsing
```

curl (Git Bash / WSL):

```bash
u=cenario1
p=$(echo -n "Test1234!" | base64)
curl -i -H "Authorization: Bearer apitoken123" "http://localhost:8080/login?username=${u}&password=${p}"
```

Onde editar os mocks

Os dados de mock estão no arquivo `mocks.go` (mesmo pacote). Altere `var scenarios` se quiser modificar ou adicionar usuários.

Observações

- O mock foi organizado para facilitar edição rápida dos cenários. Não há persistência — é apenas em memória.
- Se quiser, posso adicionar testes unitários (`main_test.go`) ou alterar para um POST JSON mais realista.

