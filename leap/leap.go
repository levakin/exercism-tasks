// Package leap implements leap year check
package leap

// IsLeapYear checks if the year is leap or not
// leap year is every year that is evenly divisible by 4
//  except every year that is evenly divisible by 100
//    unless the year is also evenly divisible by 400
func IsLeapYear(y int) bool {
	return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
}
