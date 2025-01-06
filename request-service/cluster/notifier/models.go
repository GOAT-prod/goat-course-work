package notifier

import "request-service/domain"

var (
	SubjectByRequestType = map[domain.Type]string{
		domain.ApproveType: "Ваш продукт %s успешно подтвержден",
		domain.SupplyType:  "Необходима поставка продукции %s",
	}

	SupplyBodyTemplate = "Цвет: %s\nРазмер: %d\nКоличество: %d\n\n"
)

type Mail struct {
	Subject string
	Body    string
}
