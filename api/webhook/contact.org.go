package webhook

type ContactOrg struct {
	Company    *string `json:"company,omitempty"`    // Nome da empresa
	Department *string `json:"department,omitempty"` // Departamento
	Title      *string `json:"title,omitempty"`      // Cargo
}
