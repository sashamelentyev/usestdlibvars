package time

// DAY

func sunday() {
	_ = "Sunday" // want `"Sunday" can be replaced by time.Sunday.String\(\)`
}

func monday() {
	_ = "Monday" // want `"Monday" can be replaced by time.Monday.String\(\)`
}

func tuesday() {
	_ = "Tuesday" // want `"Tuesday" can be replaced by time.Tuesday.String\(\)`
}

func wednesday() {
	_ = "Wednesday" // want `"Wednesday" can be replaced by time.Wednesday.String\(\)`
}

func thursday() {
	_ = "Thursday" // want `"Thursday" can be replaced by time.Thursday.String\(\)`
}

func friday() {
	_ = "Friday" // want `"Friday" can be replaced by time.Friday.String\(\)`
}

func saturday() {
	_ = "Saturday" // want `"Saturday" can be replaced by time.Saturday.String\(\)`
}

// MONTH

func january() {
	_ = "January" // want `"January" can be replaced by time.January.String\(\)`
}

func february() {
	_ = "February" // want `"February" can be replaced by time.February.String\(\)`
}

func march() {
	_ = "March" // want `"March" can be replaced by time.March.String\(\)`
}

func april() {
	_ = "April" // want `"April" can be replaced by time.April.String\(\)`
}

func may() {
	_ = "May" // want `"May" can be replaced by time.May.String\(\)`
}

func june() {
	_ = "June" // want `"June" can be replaced by time.June.String\(\)`
}

func july() {
	_ = "July" // want `"July" can be replaced by time.July.String\(\)`
}

func august() {
	_ = "August" // want `"August" can be replaced by time.August.String\(\)`
}

func september() {
	_ = "September" // want `"September" can be replaced by time.September.String\(\)`
}

func october() {
	_ = "October" // want `"October" can be replaced by time.October.String\(\)`
}

func november() {
	_ = "November" // want `"November" can be replaced by time.November.String\(\)`
}

func december() {
	_ = "December" // want `"December" can be replaced by time.December.String\(\)`
}

// LAYOUT

func layout() {
	_ = "01/02 03:04:05PM '06 -0700" // want `"01/02 03:04:05PM '06 -0700" can be replaced by time.Layout`
}

func ansic() {
	_ = "Mon Jan _2 15:04:05 2006" // want `"Mon Jan _2 15:04:05 2006" can be replaced by time.ANSIC`
}

func unixDate() {
	_ = "Mon Jan _2 15:04:05 MST 2006" // want `"Mon Jan _2 15:04:05 MST 2006" can be replaced by time.UnixDate`
}

func rubyDate() {
	_ = "Mon Jan 02 15:04:05 -0700 2006" // want `"Mon Jan 02 15:04:05 -0700 2006" can be replaced by time.RubyDate`
}

func rfc822() {
	_ = "02 Jan 06 15:04 MST" // want `"02 Jan 06 15:04 MST" can be replaced by time.RFC822`
}

func rfc822Z() {
	_ = "02 Jan 06 15:04 -0700" // want `"02 Jan 06 15:04 -0700" can be replaced by time.RFC822Z`
}

func rfc850() {
	_ = "Monday, 02-Jan-06 15:04:05 MST" // want `"Monday, 02-Jan-06 15:04:05 MST" can be replaced by time.RFC850`
}

func rfc1123() {
	_ = "Mon, 02 Jan 2006 15:04:05 MST" // want `"Mon, 02 Jan 2006 15:04:05 MST" can be replaced by time.RFC1123`
}

func rfc1123Z() {
	_ = "Mon, 02 Jan 2006 15:04:05 -0700" // want `"Mon, 02 Jan 2006 15:04:05 -0700" can be replaced by time.RFC1123Z`
}

func rfc3339() {
	_ = "2006-01-02T15:04:05Z07:00" // want `"2006-01-02T15:04:05Z07:00" can be replaced by time.RFC3339`
}

func rfc3339Nano() {
	_ = "2006-01-02T15:04:05.999999999Z07:00" // want `"2006-01-02T15:04:05.999999999Z07:00" can be replaced by time.RFC3339Nano`
}

func kitchen() {
	_ = "3:04PM" // want `"3:04PM" can be replaced by time.Kitchen`
}

func stamp() {
	_ = "Jan _2 15:04:05" // want `"Jan _2 15:04:05" can be replaced by time.Stamp`
}

func stampMilli() {
	_ = "Jan _2 15:04:05.000" // want `"Jan _2 15:04:05.000" can be replaced by time.StampMilli`
}

func stampMicro() {
	_ = "Jan _2 15:04:05.000000" // want `"Jan _2 15:04:05.000000" can be replaced by time.StampMicro`
}

func stampNano() {
	_ = "Jan _2 15:04:05.000000000" // want `"Jan _2 15:04:05.000000000" can be replaced by time.StampNano`
}
