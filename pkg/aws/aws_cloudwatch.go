package aws

type CloudWatchService interface {
	Info(a ...interface{}) bool
	Error(a ...interface{}) bool
	Warning(a ...interface{}) bool
	SendLog(flag string, a ...interface{}) bool
}
