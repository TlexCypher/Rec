package valueobject

type CategoryName struct {
	LiteralBase[string]
}

func (c CategoryName) Value() string {
	return c.v
}

func (c CategoryName) Validate() error {
	return nil
}

func NewCategoryNameFromString(src string) (CategoryName, error) {
	c := CategoryName{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := c.Validate(); err != nil {
		return CategoryName{}, err
	}
	return c, nil
}
