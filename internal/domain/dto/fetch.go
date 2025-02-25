package dto

import (
	"time"

	"github.com/xgmsx/go-tiny-url/internal/domain/entity"
)

type FetchLinkInput struct {
	Alias string `json:"alias"`
}

func (i FetchLinkInput) Validate() error {
	if len(i.Alias) < 2 {
		return entity.ErrInputValidation
	}
	return nil
}

type FetchLinkOutput struct {
	URL       string    `json:"url"`
	Alias     string    `json:"alias"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (o FetchLinkOutput) Load(l *entity.Link) FetchLinkOutput {
	o.URL = l.URL
	o.Alias = l.Alias
	o.ExpiredAt = l.ExpiredAt

	return o
}
