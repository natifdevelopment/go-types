package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var encryptionKey string

func SetEncryptionKey(key string) {
	encryptionKey = key
}

func GetEncryptionKey() string {
	return encryptionKey
}

type DateField time.Time

const dateFormat = "2006-01-02"

func (df *DateField) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	t, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	*df = DateField(t)
	return nil
}

func (df DateField) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(df).Format(dateFormat))
}

func (df DateField) Value() (driver.Value, error) {
	return time.Time(df), nil
}

func (df DateField) DateValue() (time.Time, error) {
	return time.Time(df), nil
}

func (df DateField) Time() time.Time {
	return time.Time(df)
}

func (df *DateField) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		*df = DateField(t)
		return nil
	}
	return fmt.Errorf("cannot convert %v to DateField", value)
}

var bulanID = [...]string{
	"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

const dateTimeFormatID = "%02d %s %d %s"

func (df DateField) FormatTanggalFull() string {
	t := time.Time(df)
	day := t.Day()
	month := bulanID[int(t.Month())]
	year := t.Year()
	return fmt.Sprintf("%02d %s %d", day, month, year)
}

func (df DateField) FormatTanggalDmy() string {
	return time.Time(df).Format("02-01-2006")
}

type DateTimeField time.Time

const dateTimeFormat = "2006-01-02 15:04:05"

func (df *DateTimeField) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	t, err := time.Parse(dateTimeFormat, str)
	if err != nil {
		return err
	}
	*df = DateTimeField(t)
	return nil
}

func (df DateTimeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(df).Format(dateTimeFormat))
}

func (df DateTimeField) Value() (driver.Value, error) {
	return time.Time(df), nil
}

func (df DateTimeField) DateValue() (time.Time, error) {
	return time.Time(df), nil
}

func (df DateTimeField) Time() time.Time {
	return time.Time(df)
}

func (df *DateTimeField) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		*df = DateTimeField(t)
		return nil
	}
	return fmt.Errorf("cannot convert %v to DateTimeField", value)
}

func (df *DateTimeField) DaysBetween(b DateTimeField) float64 {
	ta := time.Time(*df)
	tb := time.Time(b)
	diff := ta.Sub(tb)
	return diff.Hours() / 24
}

func (df DateTimeField) FormatTanggalWaktu() string {
	t := time.Time(df)
	day := t.Day()
	month := bulanID[int(t.Month())]
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	timeStr := fmt.Sprintf("%02d:%02d", hour, minute)

	return fmt.Sprintf(dateTimeFormatID, day, month, year, timeStr)
}

func (df DateTimeField) FormatTanggalWaktuFull() string {
	t := time.Time(df)
	day := t.Day()
	month := bulanID[int(t.Month())]
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	timeStr := fmt.Sprintf("Pukul %02d:%02d WIB", hour, minute)

	return fmt.Sprintf(dateTimeFormatID, day, month, year, timeStr)
}

func (df DateTimeField) FormatTanggalWaktuZonaWaktu(zonaWaktu string) string {
	t := time.Time(df)
	day := t.Day()
	month := bulanID[int(t.Month())]
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	timeStr := fmt.Sprintf("Pukul %02d:%02d %s", hour, minute, zonaWaktu)

	return fmt.Sprintf(dateTimeFormatID, day, month, year, timeStr)
}

func (df DateTimeField) FormatTanggalDmy() string {
	return time.Time(df).Format("02-01-2006")
}

func (df DateTimeField) FormatTanggalDmyHmSlash() string {
	return time.Time(df).Format("02/01/2006 15:04")
}

func DaysBetween(a, b time.Time) float64 {
	diff := a.Sub(b)
	return diff.Hours() / 24
}

func ParseDate(value string) (*time.Time, error) {
	date, err := time.Parse(dateFormat, value)
	if err != nil {
		return nil, err
	}
	return &date, nil
}

type YearMonthField time.Time

const yearMonthFormat = "2006-01"

func normalizeYmf(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		1,
		0, 0, 0, 0,
		t.Location(),
	)
}

func (ymf *YearMonthField) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	t, err := time.Parse(yearMonthFormat, str)
	if err != nil {
		return err
	}
	*ymf = YearMonthField(normalizeYmf(t))
	return nil
}

func (ymf YearMonthField) MarshalJSON() ([]byte, error) {
	return json.Marshal(normalizeYmf(time.Time(ymf)).Format(yearMonthFormat))
}

func (ymf YearMonthField) Value() (driver.Value, error) {
	return normalizeYmf(time.Time(ymf)), nil
}

func (ymf YearMonthField) DateValue() (time.Time, error) {
	return normalizeYmf(time.Time(ymf)), nil
}

func (ymf YearMonthField) Time() time.Time {
	return normalizeYmf(time.Time(ymf))
}

func (ymf *YearMonthField) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		*ymf = YearMonthField(normalizeYmf(t))
		return nil
	}
	return fmt.Errorf("cannot convert %v to DateField", value)
}

func (ymf YearMonthField) String() string {
	return normalizeYmf(time.Time(ymf)).Format(yearMonthFormat)
}

func (ymf *YearMonthField) Parse(str string) error {
	t, err := time.Parse(yearMonthFormat, str)
	if err != nil {
		return err
	}
	*ymf = YearMonthField(normalizeYmf(t))
	return nil
}

func (ymf YearMonthField) YearStr() string {
	return time.Time(ymf).Format("2006")
}

func (ymf YearMonthField) MonthStr() string {
	return time.Time(ymf).Format("01")
}

func (ymf YearMonthField) Year() int {
	return time.Time(ymf).Year()
}

func (ymf YearMonthField) Month() int {
	return int(time.Time(ymf).Month())
}

func (ymf *YearMonthField) AddMonth(monthCount int) {
	t := normalizeYmf(time.Time(*ymf))
	*ymf = YearMonthField(t.AddDate(0, monthCount, 0))
}

func (ymf *YearMonthField) FormatMonthYear() string {
	t := normalizeYmf(time.Time(*ymf))
	month := bulanID[int(t.Month())]
	year := t.Year()
	return fmt.Sprintf("%s %d", month, year)
}

func (ymf *YearMonthField) FormatYyyyMm() string {
	return ymf.Time().Format("2006-01")
}

func (ymf *YearMonthField) TimeReset() time.Time {
	t := normalizeYmf(time.Time(*ymf))
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

type EncryptedField string

func (e EncryptedField) Value() (driver.Value, error) {
	if encryptionKey == "" {
		return nil, errors.New("encryption key not set")
	}
	return Encrypt(string(e), encryptionKey)
}

func (e *EncryptedField) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("unsupported type for EncryptedField")
	}

	if encryptionKey == "" {
		return errors.New("encryption key not set")
	}

	decrypted, err := Decrypt(str, encryptionKey)
	if err != nil {
		return err
	}

	*e = EncryptedField(decrypted)
	return nil
}

func (e EncryptedField) ToString() string {
	return string(e)
}
