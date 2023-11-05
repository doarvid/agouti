package appium

import "github.com/doarvid/agouti"

func NewTestDevice(session mobileSession) *Device {
	return &Device{
		Page:    &agouti.Page{},
		session: session,
	}
}
