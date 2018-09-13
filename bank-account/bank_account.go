/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 08-09-2018
 * |
 * | File Name:     bank_account.go
 * +===============================================
 */

package account

import "sync"

// Account is a bank account
type Account struct {
	close  bool         // indicates whether the account is closed or open
	amount int64        // indicates current balance of the account
	lck    sync.RWMutex // read-write lock
}

// Open opens new back account with given inital deposit
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}

	return &Account{
		amount: deposit,
	}
}

// Close closes account a. it returns accounts payout and status of opertation
func (a *Account) Close() (payout int64, ok bool) {
	a.lck.Lock()
	defer a.lck.Unlock()

	if a.close {
		return 0, false
	}

	a.close = true

	return a.amount, true
}

// Balance returns current balance of account a. it returns balance and status of operation
func (a *Account) Balance() (balance int64, ok bool) {
	a.lck.RLock()
	defer a.lck.RUnlock()

	if a.close {
		return 0, false
	}

	return a.amount, true
}

// Deposit adds a given amount into account a's balance. The amount could be negative, to represent withdraws
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.lck.Lock()
	defer a.lck.Unlock()

	if a.close {
		return 0, false
	}

	if a.amount+amount < 0 {
		return 0, false
	}

	a.amount += amount

	return a.amount, true
}
