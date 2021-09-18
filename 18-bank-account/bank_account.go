package account

import "sync"

type Account struct {
	isOpen  bool
	balance int64
	mutex   sync.Mutex
}

// Open creates a new bank account with an initial deposit
// the initial deposit cannot be negative
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{isOpen: true, balance: initialDeposit}
}

// Close closes the account if it is open and cashes out the remaining balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen {
		return
	}
	payout, a.balance = a.balance, 0
	a.isOpen, ok = false, true
	return
}

// Balance returns the remaining account balance if it is still open
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen {
		return
	}
	balance, ok = a.balance, true
	return
}

// Deposit handles deposits and withdrawals
// If it is not posible to proceed, returns false
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.isOpen || a.balance+amount < 0 {
		return
	}
	a.balance += amount
	newBalance, ok = a.balance, true
	return
}
