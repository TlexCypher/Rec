package valueobject

type RepresentProductName struct {
	LiteralBase[string]
}

func (rpn RepresentProductName) Value() string {
	return rpn.v
}

func (rpn RepresentProductName) Validate() error {
	return nil
}

func NewRepresentProductNameFromString(src string) (RepresentProductName, error) {
	rpn := RepresentProductName{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := rpn.Validate(); err != nil {
		return RepresentProductName{}, err
	}
	return rpn, nil
}
