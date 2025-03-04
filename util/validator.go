/***********************************************************************
     Copyright (c) 2025 GNU/Linux Users' Group (NIT Durgapur)
     Author: Dhruba Sinha
************************************************************************/

package util

import "time"

// check whether start_time and end_time are valid according to the RFC3339 format and are in the future
func IsValidDuration(startTime, endTime string) bool {
	startTimeParsed, err1 := time.Parse(time.RFC3339, startTime)
	endTimeParsed, err2 := time.Parse(time.RFC3339, endTime)
	return ((err1 == nil) && (err2 == nil) && (startTimeParsed.Before(endTimeParsed)) && (startTimeParsed.After(time.Now())))
}

// check whether the given string is a valid email
func IsValidEmail(email string) bool {
	/*
		Email validation logic goes here
	*/
	return true
}

// check whether the password is strong enough
func IsValidPassword(password string) bool {
	/*
		Password validation logic goes here
	*/
	return true
}
