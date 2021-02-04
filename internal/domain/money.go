package domain

const (
	CadToUsdRate = 0.79
)

// Money represents an amount of money.
//
// The zero value for this type is 0 usd.
type Money struct {
	usd float32
}

// NewMoneyFromUsd returns a Money with the amount in Usd.
func NewMoneyFromUsd(usd float32) Money {
	return Money{
		usd: usd,
	}
}

// NewMoneyFromCad returns a Money with the amount in Cad.
func NewMoneyFromCad(cad float32) Money {
	return NewMoneyFromUsd(CadToUsdRate * cad)
}

// Add returns a new Money with the amount added.
func (m Money) Add(money Money) Money {
	m.usd += money.usd
	return m
}

// Subtract returns a new Money with the amount subtracted.
func (m Money) Subtract(money Money) Money {
	m.usd -= money.usd
	return m
}

// IsNegative returns whether the money amount is negative.
func (m Money) IsNegative() bool {
	return m.usd < 0
}

// Usd returns the amount in Usd.
func (m Money) Usd() float32 {
	return m.usd
}

// Cad returns the amount in Cad.
func (m Money) Cad() float32 {
	return m.Usd() / CadToUsdRate
}
