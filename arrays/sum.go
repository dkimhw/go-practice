package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Sum(arr []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(arr, add, 0)
}

func SumAll(numsToSum ...[]int) []int {
	var result []int

	for _, nums := range numsToSum {
		result = append(result, Sum(nums))
	}

	return result
}

func SumAllTails(numsToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numsToSum, sumTail, []int{})
}

// We've added a second type constraint which has allowed us to loosen the constraints on Reduce.
// This allows people to Reduce from a collection of A into a B.
func Reduce[A, B any](collection []A, f func(B, A) B, initialvalue B) B {
	result := initialvalue
	for _, x := range collection {
		result = f(result, x)
	}

	return result
}

// More generics

type Person struct {
	Name string
}

func Find[A any](collection []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	return
}
