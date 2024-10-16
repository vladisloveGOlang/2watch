package legalentities

import (
	"github.com/krisch/crm-backend/pkg/postgres"
)

type BankAccounts interface {
	DeleteBankAccount(a BankAcc) (Ans, error)
	GetAllBankAccounts() ([]BankAcc, error)
	CreateBankAccount(a BankAcc) (Ans, error)
	UpdateBankAccount(a BankAcc) (Ans, error)
}

type LegalEntitiRepository struct {
	db *postgres.GDB
}

func NewLegalEntitiesRepository(db *postgres.GDB) *LegalEntitiRepository {
	return &LegalEntitiRepository{
		db: db,
	}
}

func (r *LegalEntitiRepository) GetAllBankAccounts() ([]BankAcc, error) {
	var accList []BankAcc
	err := r.db.DB.Find(&accList).Error
	return accList, err
}

func (r *LegalEntitiRepository) CreateBankAccount(a BankAcc) (Ans, error) {
	result := r.db.DB.Create(&a)
	if result.Error != nil {

		ans := Ans{
			ChangedCount: int(result.RowsAffected),
			CorrAcc:      a.CorAcc,
		}
		return ans, result.Error
	}
	ans := Ans{
		ChangedCount: int(result.RowsAffected),
		CorrAcc:      a.CorAcc,
	}
	return ans, result.Error
}

func (r *LegalEntitiRepository) DeleteBankAccount(a BankAcc) (Ans, error) {
	result := r.db.DB.Model(&a).Where("cur_acc= ?", a.CurAcc).Delete(&a)
	var ans Ans
	if result.RowsAffected == 0 {
		ans.ChangedCount = int(result.RowsAffected)
		ans.CorrAcc = a.CorAcc
	}
	ans.ChangedCount = int(result.RowsAffected)
	ans.CorrAcc = a.CorAcc

	return ans, result.Error
}

func (r *LegalEntitiRepository) UpdateBankAccount(a BankAcc) (Ans, error) {
	res := r.db.DB.Model(&a).Where("cur_acc= ?", a.CurAcc).Omit("cur_acc").Updates(map[string]interface{}{
		"name": a.Name, "address": a.Address,
		"cur_acc": a.CurAcc, "cor_acc": a.CorAcc, "currency": a.Currency,
		"comment": a.Comment, "bik": a.Bik,
	})

	var ans Ans
	ans.ChangedCount = int(res.RowsAffected)
	ans.CorrAcc = a.CorAcc

	return ans, res.Error
}
