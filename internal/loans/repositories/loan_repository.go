package repositories

import (
	"errors"
	"library/internal/loans/models"
	"sync"
)

type LoanRepository struct {
	loans  map[int64]*models.Loan
	mu     sync.RWMutex
	nextId int64
}

func NewLoanRepository() models.LoanRepository {
	return &LoanRepository{
		loans:  make(map[int64]*models.Loan),
		nextId: 1,
	}
}

// CreateLoan implements [models.LoanRepository].
func (l *LoanRepository) CreateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan.ID = l.nextId
	l.nextId++
	l.loans[loan.ID] = loan
	return nil
}

// GetActiveUserLoans implements [models.LoanRepository].
func (l *LoanRepository) GetActiveUserLoans(userId int64) ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var activeLoans []*models.Loan
	for _, loan := range l.loans {
		if loan.UserID == userId && loan.Status == "active" {
			activeLoans = append(activeLoans, loan)
		}
	}
	return activeLoans, nil
}

// GetAllLoans implements [models.LoanRepository].
func (l *LoanRepository) GetAllLoans() ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loans := make([]*models.Loan, 0, len(l.loans))
	for _, loan := range l.loans {
		loans = append(loans, loan)
	}

	return loans, nil
}

// GetLoan implements [models.LoanRepository].
func (l *LoanRepository) GetLoan(id int64) (*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[id]
	if !exists {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

// ReturnBook implements [models.LoanRepository].
func (l *LoanRepository) ReturnBook(loanId int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[loanId]
	if !exists {
		return errors.New("loan not found")
	}

	loan.Status = "returned"

	return nil
}

// UpdateLoan implements [models.LoanRepository].
func (l *LoanRepository) UpdateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exists := l.loans[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}

	l.loans[loan.ID] = loan

	return nil
}
