package pointer

import (
	"errors"
	"fmt"
	"testing"
)

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Desposit(b Bitcoin) {
	fmt.Println("address of balance :", &w.balance)
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("InsufficientFundsError")

func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return InsufficientFundsError
	}
	w.balance -= b
	return nil
}

func TestPointer(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	}

	assertError := func(t *testing.T, want error, got error) {
		if want != got {
			t.Errorf("expect error '%s',but got '%s'", want, got)
		}
	}

	t.Run("desposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Desposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(2)

		assertBalance(t, wallet, Bitcoin(8))
	})

	t.Run("test insufficient bunds", func(t *testing.T) {
		start := Bitcoin(10)
		wallet := Wallet{start}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(10))
		// assertError(t, nil, err)
		assertError(t, InsufficientFundsError, err)
	})

	t.Run("stringer", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Desposit(10)

		got := wallet.balance.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})

}
