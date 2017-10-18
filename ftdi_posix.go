// +build !windows,cgo

package ftdi

/*
#cgo CFLAGS: -I /usr/include/libftdi1 -I /usr/local/include/libusb-1.0 -I /usr/local/include/libftdi1
#cgo LDFLAGS: -lftdi1 -lusb-1.0 -L/usr/local/lib
*/
import "C"
