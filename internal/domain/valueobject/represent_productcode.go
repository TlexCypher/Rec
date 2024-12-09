package valueobject

type RepresentProductCode struct {
	LiteralBase[string]
}

func (rpn RepresentProductCode) Value() string {
	return rpn.v
}

func (rpn RepresentProductCode) Validate() error {
	return nil
}

func NewRepresentProductCodeFromString(src string) (RepresentProductCode, error) {
	rpc := RepresentProductCode{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := rpc.Validate(); err != nil {
		return RepresentProductCode{}, err
	}
	return rpc, nil
}
