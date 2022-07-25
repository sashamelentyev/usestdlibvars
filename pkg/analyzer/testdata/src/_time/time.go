package _time

func sunday() {
	_ = "Sunday" // want `can use time.Sunday.String\(\) instead "Sunday"`
}

func monday() {
	_ = "Monday" // want `can use time.Monday.String\(\) instead "Monday"`
}

func tuesday() {
	_ = "Tuesday" // want `can use time.Tuesday.String\(\) instead "Tuesday"`
}

func wednesday() {
	_ = "Wednesday" // want `can use time.Wednesday.String\(\) instead "Wednesday"`
}

func thursday() {
	_ = "Thursday" // want `can use time.Thursday.String\(\) instead "Thursday"`
}

func friday() {
	_ = "Friday" // want `can use time.Friday.String\(\) instead "Friday"`
}

func saturday() {
	_ = "Saturday" // want `can use time.Saturday.String\(\) instead "Saturday"`
}

func january() {
	_ = "January" // want `can use time.January.String\(\) instead "January"`
}

func february() {
	_ = "February" // want `can use time.February.String\(\) instead "February"`
}

func march() {
	_ = "March" // want `can use time.March.String\(\) instead "March"`
}

func april() {
	_ = "April" // want `can use time.April.String\(\) instead "April"`
}

func may() {
	_ = "May" // want `can use time.May.String\(\) instead "May"`
}

func june() {
	_ = "June" // want `can use time.June.String\(\) instead "June"`
}

func july() {
	_ = "July" // want `can use time.July.String\(\) instead "July"`
}

func august() {
	_ = "August" // want `can use time.August.String\(\) instead "August"`
}

func september() {
	_ = "September" // want `can use time.September.String\(\) instead "September"`
}

func october() {
	_ = "October" // want `can use time.October.String\(\) instead "October"`
}

func november() {
	_ = "November" // want `can use time.November.String\(\) instead "November"`
}

func december() {
	_ = "December" // want `can use time.December.String\(\) instead "December"`
}

func layout() {
	_ = "01/02 03:04:05PM '06 -0700" // want `can use time.Layout instead "01/02 03:04:05PM '06 -0700"`
}

func ansic() {
	_ = "Mon Jan _2 15:04:05 2006" // want `can use time.ANSIC instead "Mon Jan _2 15:04:05 2006"`
}

func unixDate() {
	_ = "Mon Jan _2 15:04:05 MST 2006" // want `can use time.UnixDate instead "Mon Jan _2 15:04:05 MST 2006"`
}

func rubyDate() {
	_ = "Mon Jan 02 15:04:05 -0700 2006" // want `can use time.RubyDate instead "Mon Jan 02 15:04:05 -0700 2006"`
}

func rfc822() {
	_ = "02 Jan 06 15:04 MST" // want `can use time.RFC822 instead "02 Jan 06 15:04 MST"`
}

func rfc822Z() {
	_ = "02 Jan 06 15:04 -0700" // want `can use time.RFC822Z instead "02 Jan 06 15:04 -0700"`
}

func rfc850() {
	_ = "Monday, 02-Jan-06 15:04:05 MST" // want `can use time.RFC850 instead "Monday, 02-Jan-06 15:04:05 MST"`
}

func rfc1123() {
	_ = "Mon, 02 Jan 2006 15:04:05 MST" // want `can use time.RFC1123 instead "Mon, 02 Jan 2006 15:04:05 MST"`
}

func rfc1123Z() {
	_ = "Mon, 02 Jan 2006 15:04:05 -0700" // want `can use time.RFC1123Z instead "Mon, 02 Jan 2006 15:04:05 -0700"`
}

func rfc3339() {
	_ = "2006-01-02T15:04:05Z07:00" // want `can use time.RFC3339 instead "2006-01-02T15:04:05Z07:00"`
}

func rfc3339Nano() {
	_ = "2006-01-02T15:04:05.999999999Z07:00" // want `can use time.RFC3339Nano instead "2006-01-02T15:04:05.999999999Z07:00"`
}

func kitchen() {
	_ = "3:04PM" // want `can use time.Kitchen instead "3:04PM"`
}

func stamp() {
	_ = "Jan _2 15:04:05" // want `can use time.Stamp instead "Jan _2 15:04:05"`
}

func stampMilli() {
	_ = "Jan _2 15:04:05.000" // want `can use time.StampMilli instead "Jan _2 15:04:05.000"`
}

func stampMicro() {
	_ = "Jan _2 15:04:05.000000" // want `can use time.StampMicro instead "Jan _2 15:04:05.000000"`
}

func stampNano() {
	_ = "Jan _2 15:04:05.000000000" // want `can use time.StampNano instead "Jan _2 15:04:05.000000000"`
}
