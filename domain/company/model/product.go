package model

type (
	// Product is an item that company provides.
	// We don't focus on concrete attributes cause different companies may provide
	// absolutely different type of products or goods.
	Product struct {
		Name       string
		Attributes map[ProductAttribute]ProductAttributeValue
	}

	// ProductAttribute is description of a concrete product attribute, like length, weight etc.
	ProductAttribute string
	// ProductAttributeValue is a concrete attribute value, for instance 100 for length attribute.
	ProductAttributeValue string
)
