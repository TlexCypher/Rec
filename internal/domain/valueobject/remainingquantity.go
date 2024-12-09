package valueobject

type RemainingQuantity struct {
	LiteralBase[int64]
}

func (rq RemainingQuantity) Validate() error {
	return nil
}

func (rq RemainingQuantity) Value() int64 {
	return rq.v
}

func NewRemainingQuantityFromInt64(src int64) (RemainingQuantity, error) {
	rq := RemainingQuantity{
		LiteralBase[int64]{
			v: src,
		},
	}
	if err := rq.Validate(); err != nil {
		return RemainingQuantity{}, err
	}
	return rq, nil
}
