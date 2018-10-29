package models

// DataEvent structure for data send to web interface
type DataEvent struct { // structure for data send to web interface
	PackageNumber int // time stamp
	Address       string
	HexString     string
	ASCIIString   string
}

// RecMessage structure for data received from web
type RecMessage struct {
	Name     string
	Port     int
	Action   string
	Protocol string
}
