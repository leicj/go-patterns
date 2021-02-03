package main

import (
	"fmt"
	template "p21/template"
)

func main() {
	smsOTP := &template.Sms{}
	o := template.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &template.Email{}
	o = template.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)
}
