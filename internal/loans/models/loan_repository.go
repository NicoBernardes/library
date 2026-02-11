package models

type LoanRepository interface {
	CreateLoan(loan *Loan) error
	UpdateLoan(loan *Loan) error
	ReturnBook(loanId int64) error
	GetLoan(id int64) (*Loan, error)
	GetActiveUserLoans(userId int64) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}
