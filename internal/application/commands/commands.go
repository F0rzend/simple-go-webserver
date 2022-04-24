package commands

type Commands struct {
	CreateUser       CreateUserCommandHandler
	UpdateUser       UpdateUserCommandHandler
	ChangeBTCBalance ChangeBTCBalanceCommandHandler
	ChangeUSDBalance ChangeUSDBalanceCommandHandler

	SetBTCPrice SetBTCPriceCommandHandler
}
