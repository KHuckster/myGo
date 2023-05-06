package point

import (
	"errors"
	"fmt"
	"testing"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func TestWallet(t *testing.T) {

	amountCheck := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	}

	errCheck := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0)}
		wallet.Deposit(10)

		amountCheck(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.Withdraw(10)

		amountCheck(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with error", func(t *testing.T) {
		startBitcoin := Bitcoin(10)
		wallet := Wallet{balance: startBitcoin}
		err := wallet.Withdraw(20)

		amountCheck(t, wallet, Bitcoin(10))
		errCheck(t, err, InsufficientFundsError)
	})

}
