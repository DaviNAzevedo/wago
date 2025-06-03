package webhook

type NFMReply struct {
	ResponseJSON string `json:"response_json"` // String JSON contendo a resposta do Flow
	Body         string `json:"body"`          // Mensagem de confirmação para o usuário
	Name         string `json:"name"`          // Nome do Flow
}
