package commands

type Commands struct {
	CreateUser *CreateUserCommandHandler
	UpdateUser *UpdateUserCommandHandler

	SetBTCPrice *SetBTCPriceCommandHandler
}
