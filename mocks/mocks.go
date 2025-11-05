package mocks

// mocks package contém os dados de mock usados pelo servidor

type User struct {
	Password string
	Payload  map[string]interface{}
}

var Scenarios = map[string]User{
	"cenario1": {
		Password: "Test1234!",
		Payload: map[string]interface{}{
			"id_usuario": 1,
			"id_pessoa":  10,
			"nome":       "Cenário 1",
			"email":      "cenario1@mail.com",
			"cpf":        "379.137.150-90",
			"perfil":     []string{"docente"},
		},
	},
	"cenario2": {
		Password: "Test5678!",
		Payload: map[string]interface{}{
			"id_usuario": 2,
			"id_pessoa":  20,
			"nome":       "Cenário 2",
			"email":      "cenario2@mail.com",
			"cpf":        "343.028.030-37",
			"perfil":     []string{"discente"},
		},
	},
	"cenario3": {
		Password: "Test9012!",
		Payload: map[string]interface{}{
			"id_usuario": 3,
			"id_pessoa":  30,
			"nome":       "Cenário 3",
			"email":      "cenario3@mail.com",
			"cpf":        "205.188.000-08",
			"perfil":     []string{"docente", "discente"},
		},
	},
	"cenario4": {
		Password: "Test1234!",
		Payload: map[string]interface{}{
			"id_usuario": 4,
			"id_pessoa":  40,
			"nome":       "Cenário 4",
			"email":      "cenario4@mail.com",
			"cpf":        "603.559.990-72",
			"perfil":     []string{"discente"},
		},
	},
}
