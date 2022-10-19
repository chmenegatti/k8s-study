package responses

import (
	"time"

	"github.com/chmenegatti/k8s-study/src/models"
	"github.com/google/uuid"
)

type nomes struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NomesResponse(nome models.Nomes) nomes {
	return nomes{ID: nome.ID, Nome: nome.Nome, Email: nome.Email,
		CreatedAt: nome.CreatedAt, UpdatedAt: nome.UpdatedAt}
}
