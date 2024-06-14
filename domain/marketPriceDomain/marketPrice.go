package marketPriceDomain

type domain struct {
	Token *Token
}

type Token struct {
	Symbol  string
	Chain   string
	Network string
	Price   float64
}

func InitDomain(token *Token) Domain {
	return &domain{Token: token}
}

type Domain interface {
	Get() *Token
	ChangePrice(price float64) *Token
	Set(token *Token)
}

func (d *domain) Get() *Token {
	return d.Token
}

func (d *domain) ChangePrice(price float64) *Token {
	d.Token.Price = price

	return d.Token
}

func (d *domain) Set(token *Token) {
	d.Token = token
}
