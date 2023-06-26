package gopointers

import (
	"testing"
)

// Go copies values when you pass them to functions/methods,
//so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	//fmt.Printf("Address of balance in test is %v\n", &wallet.balance)
	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %s but want %s", got.String(), want.String())
	}
}

// The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something,
// in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance
// is necessary (like database connection pools).
func TestWallet1(t *testing.T) {

	assertMsg := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got.String(), want.String())
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(50))

		want := Bitcoin(50)
		assertMsg(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(40)}
		err := wallet.Withdraw(Bitcoin(20))

		want := Bitcoin(20)
		assertNoError(t, err)
		assertMsg(t, wallet, want)
	})

	t.Run("insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficentFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get error but want one")
	}

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, sB Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != sB {
		t.Errorf("got %s but want %s", got.String(), sB.String())
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}
