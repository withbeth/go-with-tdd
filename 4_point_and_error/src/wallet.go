package main

import "fmt"

// Requirements :
// Going to implement BITCOIN wallet.. you have to remember that
type BitCoin int

type OverDraftError struct {
	insufficient BitCoin
}
func (o *OverDraftError) Error() string {
	return fmt.Sprintf("Insufficient %d BTC", o.insufficient)
}

// type aliasing makes you can declare methods on them
// Stringer interface is defined in the fmt package.
// That interface let's you define how your type is printed when used with the %s format string in prints.
func (b BitCoin) String() (string) {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// If a symbol starts with a lowercase, then it is "private" outside the package which it is defined in.
	balance BitCoin
}

// In Go, arguments of functions or methods will be copied when every time we call
// Which means, we have to point the arguments if we want some change those states
func (w *Wallet) Deposit(income BitCoin) {
	w.balance += income
}

func (w *Wallet) Withdraw(amount BitCoin) (BitCoin, error) {
	current := w.balance
	insufficient := current - amount
	if insufficient < 0 {
		// Need Withdraw error
		// In Go, if you want to indicate an error, you have to make your function return error
		// for the caller to check and act on.
		return w.balance, &OverDraftError{insufficient}
	}
	w.balance = insufficient
	return w.balance, nil
}

func (w *Wallet) Balance() (BitCoin) {
	return w.balance
}


