package alarm

type AlarmDatasource interface {
	AlarmCounter() (*AlarmCounter, error)
}
