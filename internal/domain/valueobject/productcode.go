package valueobject

type ProductCode struct {
	LiteralBase[string]
}

func (pc ProductCode) Validate() error {
	return nil
}

func (pc ProductCode) Value() string {
	return pc.v
}

func NewProductCodeFromString(src string) (ProductCode, error) {
	pc := ProductCode{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := pc.Validate(); err != nil {
		return ProductCode{}, err
	}
	return pc, nil
}
