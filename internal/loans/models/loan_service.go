package models

type LoanService interface {
	CreateLoan(bookId, userId int64) (*Loan, error)
	ReturnBook(loanId int64) error
	GetLoan(id int64) (*Loan, error)
	GetUserLoans(userId int64) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}
