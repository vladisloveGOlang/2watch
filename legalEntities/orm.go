package legalentities

// Банковский счет.
type BankAcc struct {
	Bik      string // БИК
	Name     string // Имя банка
	Address  string // Юр. Ардес
	CurAcc   string // рассчетный счет
	CorAcc   string // корреспондентский счет
	Currency string // Тип валюьы
	Comment  string // Клментарий
}

type Ans struct {
	ChangedCount int    `json:"changed,omitempty"`
	CorrAcc      string `json:"corr_acc_,omitempty"`
}

func (BankAcc) TableName() string {
	return "legal_entities"
}
