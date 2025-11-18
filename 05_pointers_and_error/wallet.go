// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors
package pointersanderror

type Wallet struct{ balance int }

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int { return w.balance }
