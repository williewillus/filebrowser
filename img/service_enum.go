// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package img

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const (
	// FormatJpeg is a Format of type Jpeg.
	FormatJpeg Format = iota
	// FormatPng is a Format of type Png.
	FormatPng
	// FormatGif is a Format of type Gif.
	FormatGif
	// FormatTiff is a Format of type Tiff.
	FormatTiff
	// FormatBmp is a Format of type Bmp.
	FormatBmp
	// FormatHeif is a Format of type Heif.
	FormatHeif
)

var ErrInvalidFormat = errors.New("not a valid Format")

const _FormatName = "jpegpnggiftiffbmpheif"

var _FormatMap = map[Format]string{
	FormatJpeg: _FormatName[0:4],
	FormatPng:  _FormatName[4:7],
	FormatGif:  _FormatName[7:10],
	FormatTiff: _FormatName[10:14],
	FormatBmp:  _FormatName[14:17],
	FormatHeif: _FormatName[17:21],
}

// String implements the Stringer interface.
func (x Format) String() string {
	if str, ok := _FormatMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Format(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Format) IsValid() bool {
	_, ok := _FormatMap[x]
	return ok
}

var _FormatValue = map[string]Format{
	_FormatName[0:4]:   FormatJpeg,
	_FormatName[4:7]:   FormatPng,
	_FormatName[7:10]:  FormatGif,
	_FormatName[10:14]: FormatTiff,
	_FormatName[14:17]: FormatBmp,
	_FormatName[17:21]: FormatHeif,
}

// ParseFormat attempts to convert a string to a Format.
func ParseFormat(name string) (Format, error) {
	if x, ok := _FormatValue[name]; ok {
		return x, nil
	}
	return Format(0), fmt.Errorf("%s is %w", name, ErrInvalidFormat)
}

// MarshalText implements the text marshaller method.
func (x Format) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Format) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseFormat(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errFormatNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Format) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Format(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = Format(v)
	case string:
		*x, err = ParseFormat(v)
	case []byte:
		*x, err = ParseFormat(string(v))
	case Format:
		*x = v
	case int:
		*x = Format(v)
	case *Format:
		if v == nil {
			return errFormatNilPtr
		}
		*x = *v
	case uint:
		*x = Format(v)
	case uint64:
		*x = Format(v)
	case *int:
		if v == nil {
			return errFormatNilPtr
		}
		*x = Format(*v)
	case *int64:
		if v == nil {
			return errFormatNilPtr
		}
		*x = Format(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = Format(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errFormatNilPtr
		}
		*x = Format(*v)
	case *uint:
		if v == nil {
			return errFormatNilPtr
		}
		*x = Format(*v)
	case *uint64:
		if v == nil {
			return errFormatNilPtr
		}
		*x = Format(*v)
	case *string:
		if v == nil {
			return errFormatNilPtr
		}
		*x, err = ParseFormat(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x Format) Value() (driver.Value, error) {
	return x.String(), nil
}

const (
	// QualityHigh is a Quality of type High.
	QualityHigh Quality = iota
	// QualityMedium is a Quality of type Medium.
	QualityMedium
	// QualityLow is a Quality of type Low.
	QualityLow
)

var ErrInvalidQuality = errors.New("not a valid Quality")

const _QualityName = "highmediumlow"

var _QualityMap = map[Quality]string{
	QualityHigh:   _QualityName[0:4],
	QualityMedium: _QualityName[4:10],
	QualityLow:    _QualityName[10:13],
}

// String implements the Stringer interface.
func (x Quality) String() string {
	if str, ok := _QualityMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Quality(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Quality) IsValid() bool {
	_, ok := _QualityMap[x]
	return ok
}

var _QualityValue = map[string]Quality{
	_QualityName[0:4]:   QualityHigh,
	_QualityName[4:10]:  QualityMedium,
	_QualityName[10:13]: QualityLow,
}

// ParseQuality attempts to convert a string to a Quality.
func ParseQuality(name string) (Quality, error) {
	if x, ok := _QualityValue[name]; ok {
		return x, nil
	}
	return Quality(0), fmt.Errorf("%s is %w", name, ErrInvalidQuality)
}

// MarshalText implements the text marshaller method.
func (x Quality) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Quality) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseQuality(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errQualityNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Quality) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Quality(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = Quality(v)
	case string:
		*x, err = ParseQuality(v)
	case []byte:
		*x, err = ParseQuality(string(v))
	case Quality:
		*x = v
	case int:
		*x = Quality(v)
	case *Quality:
		if v == nil {
			return errQualityNilPtr
		}
		*x = *v
	case uint:
		*x = Quality(v)
	case uint64:
		*x = Quality(v)
	case *int:
		if v == nil {
			return errQualityNilPtr
		}
		*x = Quality(*v)
	case *int64:
		if v == nil {
			return errQualityNilPtr
		}
		*x = Quality(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = Quality(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errQualityNilPtr
		}
		*x = Quality(*v)
	case *uint:
		if v == nil {
			return errQualityNilPtr
		}
		*x = Quality(*v)
	case *uint64:
		if v == nil {
			return errQualityNilPtr
		}
		*x = Quality(*v)
	case *string:
		if v == nil {
			return errQualityNilPtr
		}
		*x, err = ParseQuality(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x Quality) Value() (driver.Value, error) {
	return x.String(), nil
}

const (
	// ResizeModeFit is a ResizeMode of type Fit.
	ResizeModeFit ResizeMode = iota
	// ResizeModeFill is a ResizeMode of type Fill.
	ResizeModeFill
)

var ErrInvalidResizeMode = errors.New("not a valid ResizeMode")

const _ResizeModeName = "fitfill"

var _ResizeModeMap = map[ResizeMode]string{
	ResizeModeFit:  _ResizeModeName[0:3],
	ResizeModeFill: _ResizeModeName[3:7],
}

// String implements the Stringer interface.
func (x ResizeMode) String() string {
	if str, ok := _ResizeModeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ResizeMode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ResizeMode) IsValid() bool {
	_, ok := _ResizeModeMap[x]
	return ok
}

var _ResizeModeValue = map[string]ResizeMode{
	_ResizeModeName[0:3]: ResizeModeFit,
	_ResizeModeName[3:7]: ResizeModeFill,
}

// ParseResizeMode attempts to convert a string to a ResizeMode.
func ParseResizeMode(name string) (ResizeMode, error) {
	if x, ok := _ResizeModeValue[name]; ok {
		return x, nil
	}
	return ResizeMode(0), fmt.Errorf("%s is %w", name, ErrInvalidResizeMode)
}

// MarshalText implements the text marshaller method.
func (x ResizeMode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ResizeMode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseResizeMode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errResizeModeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ResizeMode) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ResizeMode(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = ResizeMode(v)
	case string:
		*x, err = ParseResizeMode(v)
	case []byte:
		*x, err = ParseResizeMode(string(v))
	case ResizeMode:
		*x = v
	case int:
		*x = ResizeMode(v)
	case *ResizeMode:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = *v
	case uint:
		*x = ResizeMode(v)
	case uint64:
		*x = ResizeMode(v)
	case *int:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = ResizeMode(*v)
	case *int64:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = ResizeMode(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = ResizeMode(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = ResizeMode(*v)
	case *uint:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = ResizeMode(*v)
	case *uint64:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x = ResizeMode(*v)
	case *string:
		if v == nil {
			return errResizeModeNilPtr
		}
		*x, err = ParseResizeMode(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x ResizeMode) Value() (driver.Value, error) {
	return x.String(), nil
}
