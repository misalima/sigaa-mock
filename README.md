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

Os dados de mock agora estão no pacote `mocks`, dentro da pasta `mocks/mocks.go` (mesmo módulo). O arquivo exporta a variável `Scenarios` do tipo `map[string]mocks.User` — edite esse arquivo para modificar, remover ou adicionar cenários.

Exemplo: adicionar um novo usuário (edite `mocks/mocks.go`):

```go
// dentro do map Scenarios
"cenario5": {
    Password: "NovaSenha!",
    Payload: map[string]interface{}{
        "id_usuario": 5,
        "id_pessoa":  50,
        "nome":       "Cenário 5",
        "email":      "cenario5@mail.com",
        "cpf":        "000.000.000-00",
        "perfil":     []string{"discente"},
    },
},
```

Depois de editar, rode o servidor normalmente (`go run .` ou `go build` + executar o binário).

Observações finais

- Use `go run .` no diretório do módulo para garantir que todos os pacotes são compilados juntos.
- Se preferir, movo a documentação de edição para um arquivo `CONTRIBUTING.md` ou acrescento exemplos de testes automáticos.


Nota sobre o erro "undefined: scenarios"

Causa

Se você executar `go run main.go` —— sem fornecer os outros arquivos do pacote —— o comando compila apenas `main.go` e não inclui `mocks.go`. Como `scenarios` está definido em `mocks.go`, o compilador reclama de `undefined: scenarios`.

Soluções (PowerShell)

- Recomendo executar o pacote inteiro (inclui todos os arquivos .go do pacote):

```powershell
Set-Location -LiteralPath 'D:\Projects\Estágio IF\Backend\sigaa-mock'
go run .
```

- Alternativa: compilar e executar o binário:

```powershell
Set-Location -LiteralPath 'D:\Projects\Estágio IF\Backend\sigaa-mock'
go build -o sigaa-mock.exe
.\sigaa-mock.exe
```

- Outra alternativa (incluir manualmente os arquivos ao rodar):

```powershell
cd 'D:\Projects\Estágio IF\Backend\sigaa-mock'
go run main.go mocks.go
```

Por que usar `go run .`?

Ele garante que todos os arquivos do pacote atual são compilados juntos, evitando erros onde símbolos definidos em outros arquivos do pacote (como `scenarios`) não são encontrados.
