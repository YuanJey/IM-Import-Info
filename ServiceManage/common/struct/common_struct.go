package common_struct

type CommonAskMessage struct {
	Code       int64
	ServerName string
}

type CommonReturnMessage struct {
	Code                              int64
	CentralDispatchServiceInformation []string
}

type ServerProviderAskMessage struct {
	CommonAskMessage
	ProviderAddress string
}
type ServerConsumerAskMessage struct {
	CommonAskMessage
	ConsumerAddress string
}

type ServerProviderReturnMessage struct {
	CommonReturnMessage
}
type ServerConsumerReturnMessage struct {
	CommonReturnMessage
	ServerProviderAddress string
}

type CenterAskMessage struct {
	CommonAskMessage
}
