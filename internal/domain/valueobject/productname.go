package valueobject

type ProductName struct {
	LiteralBase[string]
}

func (pn ProductName) Validate() error {
	return nil
}

func (pn ProductName) Value() string {
	return pn.v
}

func NewProductNameFromString(src string) (ProductName, error) {
	pn := ProductName{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := pn.Validate(); err != nil {
		return ProductName{}, err
	}
	return pn, nil
}
