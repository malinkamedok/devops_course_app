package usecase

type AlertUseCase struct {
	tgBot TelegramReq
}

func NewAlertUseCase(tgBot TelegramReq) *AlertUseCase {
	return &AlertUseCase{tgBot: tgBot}
}

var _ AlertContract = (*AlertUseCase)(nil)
