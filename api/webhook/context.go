package webhook

type Context struct {
	Forwarded           *bool            `json:"forwarded,omitempty"`            // Se a mensagem foi encaminhada
	FrequentlyForwarded *bool            `json:"frequently_forwarded,omitempty"` // Se a mensagem foi encaminhada muitas vezes
	From                *string          `json:"from,omitempty"`                 // ID do WhatsApp do remetente original (para respostas)
	ID                  *string          `json:"id,omitempty"`                   // ID da mensagem original (para respostas ou reações)
	ReferredProduct     *ReferredProduct `json:"referred_product,omitempty"`     // Produto referido em uma resposta a mensagem de produto
	MentionedJid        []string         `json:"mentioned_jid,omitempty"`        // IDs do WhatsApp dos usuários mencionados
}
