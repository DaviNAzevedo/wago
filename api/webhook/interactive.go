package webhook

type Interactive struct {
	Type        string       `json:"type"` // Tipo de resposta interativa: "button_reply", "list_reply", "nfm_reply"
	ButtonReply *ButtonReply `json:"button_reply,omitempty"`
	ListReply   *ListReply   `json:"list_reply,omitempty"`
	NFMReply    *NFMReply    `json:"nfm_reply,omitempty"` // Non-Flow Message Reply
}
