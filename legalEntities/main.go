package legalentities

type Service struct {
	repo *LegalEntitiRepository
}

func NewLegalEntitiesService(repo *LegalEntitiRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllBankAccounts() ([]BankAcc, error) {
	return s.repo.GetAllBankAccounts()
}

func (s *Service) CreateBankAccount(a BankAcc) (Ans, error) {
	return s.repo.CreateBankAccount(a)
}

func (s *Service) DeleteBankAccount(a BankAcc) (Ans, error) {
	return s.repo.DeleteBankAccount(a)
}

func (s *Service) UpdateBankAccount(a BankAcc) (Ans, error) {
	return s.repo.UpdateBankAccount(a)
}
