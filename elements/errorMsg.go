package elements

import "github.com/condemo/home-inventory/styles"

func NewErrorView(err error) string {
	return styles.ErrorContainer.Render(err.Error())
}
