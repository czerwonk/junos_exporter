package isis

type IsisDatasource interface {
	IsisAdjancies() (*IsisAdjacencies, error)
}
