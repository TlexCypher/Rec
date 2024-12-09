package valueobject

type Remark struct {
	LiteralBase[string]
}

func (r Remark) Validate() error {
	return nil
}

func (r Remark) Value() string {
	return r.v
}

func NewRemarkFromString(src string) (Remark, error) {
	r := Remark{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := r.Validate(); err != nil {
		return Remark{}, err
	}
	return r, nil
}
