package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Esse código nao funciona pois o ponteiro nao foi passado (*)
// Quando faco (w Wallet) significa que os valores dessa struct nao podem ser alterados
// Ao fazer (w *Wallet) o campo da struct consegue ser alterado

// func (w Wallet) Deposit(number int) {
// 	w.balance += number
// }

// func (w Wallet) Deposit(number int) {
// 	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
// 	w.balance += number
// 	fmt.Printf("address of balance in Deposit is %v \n", w.balance)
// 	// Interessante de se notar que dentro da funcao Deposit eu consigo alterar o valor na copia, porém quando ele retorna ele continua 0
// }

func (w *Wallet) Deposit(number Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += number
	fmt.Printf("address of balance in Deposit is %v \n", w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// error pode ser nil (NULL) pois ele é uma interface
// Se qualquer outra interface for retornada ela pode ser nil
func (w *Wallet) Withdraw(number Bitcoin) error {
	actualBalance := w.balance
	if actualBalance < number {
		return ErrInsufficientFunds
	}

	w.balance -= number

	return nil
}

// Aqui nao seria necessario o ponteiro, mas é melhor deixar para consistencia

// Em go quando (variables, types, functions) comecar com letra minuscula quer dizer
// que é privado para o package que ele esta

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Essa funcao String() é definida dentro do pacote fmt
// Dentro da interface Stringer
// Ela permite que formatemos um type para que quando ele for printado ele apareca de determinada maneira
// Sendo assim o type Bitcoin agora implementou a interface Stringer de dentro da biblioteca fmt
