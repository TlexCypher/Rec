package valueobject

type Category struct {
	LiteralBase[string]
}

func (c Category) Validate() error {
	return nil
}

func (c Category) Value() string {
	return c.v
}

func NewCategoryFromString(src string) (Category, error) {
	c := Category{
		LiteralBase[string]{
			v: src,
		},
	}
	if err := c.Validate(); err != nil {
		return Category{}, nil
	}
	return c, nil
}
