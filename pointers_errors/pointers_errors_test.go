package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	checkError := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		//  &wallet.balance ao colocar o & na frente eu descubro o endereco da memoria em que esse valor esta presente

		checkError(t, wallet, Bitcoin(20))
		// want := Bitcoin(20)
		// // Aqui para transformar um int normal para um typ Bitcoin que eu criei eu passo um valor
		// // Nesse caso como ambos sao int fica mais simples

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// 	// Usando o %s eles utilizam a interface Stringer e o metodo String() que formatamos manualmente para o tipo Bitcoin
		// 	// Conseguimos ver a formatacao durante o erro
		// }
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(10))

		checkError(t, wallet, Bitcoin(0))
	})
}

// In Go, if you want to indicate an error it is idiomatic for your function to return an err for the caller to check and act on.
