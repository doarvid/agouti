package appium

import "github.com/doarvid/agouti/internal/element"

type elementRepository interface {
	Get() ([]element.Element, error)
	GetAtLeastOne() ([]element.Element, error)
	GetExactlyOne() (element.Element, error)
}
