package utils

/*
type VehicleType string

const (
	Tractor   VehicleType = "Tractor"
	Emergency  = "Emergency"
	Diplomat   = "Tractor"
	Foreign    = "Foreign"
	Military   = "Military"
)

func isVehicalFeeFree(vehicleType string) (bool) {
	return false
} */

func IsHoliday() bool {
	return false
}

func IsWeekend() bool {
	return false
}

type FreeDays struct {
	Month string
	Day   int
}

func IsFeeFreeDay(month string, day int) bool {

	feeFreeDays := []FreeDays{
		{Month: "May", Day: 16},
		{Month: "Sep", Day: 19},
		{Month: "Oct", Day: 30},
		{Month: "Dec", Day: 31},
	}

	for _, item := range feeFreeDays {
		if month == item.Month && day == item.Day {
			return true
		}
	}
	return false
}
