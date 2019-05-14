package queryGenerator

import (
	"context"
	"TopNService/models/Criteria"
)

type EsQueryContext struct {
	url      string
	user     string
	password string
	criteria *Criteria
	parsedFilter interface
}

type EsQuery struct {
	//Abstract for now code to be developed
}

func NewEsQueryContext(ctx context.Context, criteria *Criteria, url, user, password string) *EsQueryContext {
	//Created EsQueryContext and return

	return &EsQueryContext{}
}

func (e *EsQueryContext) ParseFilter () {

	//Need to decide on type of result
}

func (e *EsQueryContext) Query ()(*TopResponseValue) {

	//TODO
	return TopResponseValue{}
}
