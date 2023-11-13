package models

import v1 "github.com/alwaysbespoke/coba/pkg/crds/sbc/v1"

type ListSbcsResponse struct {
	Sbcs []*v1.SBC `json:"sbcs"`
}
