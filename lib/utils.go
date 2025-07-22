package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(input string) string {
	// Remove all characters that are not alphanumeric or spaces or underscores
	s := regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(input, "")

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Title case s
	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

func WriteTextToFile(filename string, text string) error {
	// Get the directory part of the filename
	dir := filepath.Dir(filename)

	// Create the directory path if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Open the file for writing. Create it if it doesn't exist, or truncate it if it does.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

type Task func()

func BackgroundTask(ctx context.Context, task Task) {
	for {
		select {
		case <-ctx.Done():
			log.Println("background task canceled.")
			return
		default:
			var funcName string

			funcValue := reflect.ValueOf(task)
			if funcValue.Kind() == reflect.Func {
				funcName = runtime.FuncForPC(funcValue.Pointer()).Name()
			} else {
				funcName = ""
			}

			log.Println("running task ...", funcName)

			task()
		}
	}
}

func Int32Ptr(i int32) *int32 { return &i }

func StructToMap(s interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("error marshaling struct to JSON: %v", err)
	}

	var mapData map[string]interface{}
	err = json.Unmarshal(jsonData, &mapData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON to map: %v", err)
	}

	return mapData, nil
}

func UseDefault[T interface{}](value interface{}, defaultValue T) T {
	if value == nil || value == "" {
		return defaultValue
	}

	return value.(T)
}

func UseDefaultValueIf[T interface{}](target interface{}, value interface{}, defaultValue T) T {
	if value == target {
		return defaultValue
	}

	return value.(T)
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panicf("%s is empty", key)
	}

	return value
}

func TranslateHandlerError(c fiber.Ctx, err error) error {
	switch err := err.(type) {
	case HttpError:
		return c.
			Status(err.StatusCode).
			JSON(fiber.Map{"message": err.Error()})
	default:
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": err.Error()})
	}
}

func TranslateDAOError(err error) error {
	switch err := err.(type) {
	case DatabaseError:
		var errorCode int = UseDefault[int](err.ErrorCode, ErrorCodeInvalid)
		var httpErrorCode int

		if mappedCode, exists := DatabaseErrorCodeMappings[errorCode]; exists {
			httpErrorCode = mappedCode
		}

		return HttpError{Message: err.Error(), StatusCode: httpErrorCode}
	default:
		return HttpError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	}
}

func Slugify(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Replace special characters and spaces with hyphens
	regexp := regexp.MustCompile("[^a-z0-9]+")
	text = regexp.ReplaceAllString(text, "-")

	// Remove leading and trailing hyphens
	text = strings.Trim(text, "-")

	return text
}

func GetRepoRef(text string) string {
	splitedText := strings.Split(text, "/")
	length := len(splitedText)

	return splitedText[length-1]
}

func ListHasitem[T interface{}](item T, list []T) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}

	return false
}

type PostgresConnectionParam struct {
	User      string
	Password  string
	Host      string
	Port      int
	Database  string
	ExtraArgs string
}

func GetPostgresConnectionString(c PostgresConnectionParam) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", c.User, c.Password, c.Host, c.Port, c.Database, c.ExtraArgs)
}

func CopyStructFields(src, dst interface{}) error {
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		fieldName := srcValue.Type().Field(i).Name
		srcField := srcValue.Field(i)
		dstField := dstValue.FieldByName(fieldName)

		if dstField.IsValid() && dstField.CanSet() {
			if srcField.Kind() == reflect.Ptr {
				if !srcField.IsNil() {
					dstField.Set(srcField)
				}
			} else {
				if dstField.Kind() == reflect.Ptr {
					if dstField.IsNil() {
						dstField.Set(reflect.New(dstField.Type().Elem()))
					}
					dstField.Elem().Set(srcField)
				} else {
					dstField.Set(srcField)
				}
			}
		}
	}

	return nil
}

// RemoveFieldOption defines the type for options that can be passed to RemoveField.
type RemoveFieldOption func(map[string]interface{})

// Option to remove a specific field
func RemoveFieldOptionFunc(fieldName string) RemoveFieldOption {
	return func(m map[string]interface{}) {
		delete(m, fieldName)
	}
}

// RemoveField removes specified fields from the struct and returns the modified struct.
func RemoveField[T any](target T, opts ...RemoveFieldOption) T {
	// Convert struct to map
	mapData := structToMap(target)

	// Apply all options to remove fields
	for _, opt := range opts {
		opt(mapData)
	}

	// Convert map back to struct
	modifiedStruct := mapToStruct[T](mapData)

	return modifiedStruct
}

// structToMap converts a struct to a map.
func structToMap(target interface{}) map[string]interface{} {
	data, _ := json.Marshal(target)
	var mapData map[string]interface{}
	_ = json.Unmarshal(data, &mapData)

	return mapData
}

// mapToStruct converts a map to a struct of type T.
func mapToStruct[T any](mapData map[string]interface{}) T {
	data, _ := json.Marshal(mapData)
	var target T
	_ = json.Unmarshal(data, &target)

	return target
}

func GetNestedValue(m map[string]interface{}, keyPath string) interface{} {
	keys := strings.Split(keyPath, ".")
	if len(keys) == 1 {
		return m[keys[0]]
	}

	if innerMap, ok := m[keys[0]].(map[string]interface{}); ok {
		return GetNestedValue(innerMap, strings.Join(keys[1:], "."))
	}

	return nil
}
