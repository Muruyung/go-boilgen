package logger

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Muruyung/go-utilities/converter"
)

type key string

const (
	// ActivityID generate uuid
	ActivityID key = "activityID"

	// CreatedBy who created logger process
	CreatedBy key = "createdBy"

	// Email who created logger process
	Email key = "email"

	// PhoneNumber who created logger process
	PhoneNumber key = "phoneNumber"

	// UserRole who created logger process
	UserRole key = "role"

	// CMSRole who created logger process
	CMSRole key = "cmsRole"

	// IsInfluencer check log using elastic
	IsInfluencer key = "isInfluencer"
)

// APILogger api logger
func APILogger(ctx context.Context, method, path string, body interface{}) {
	var dataStr string
	conv, err := converter.ConvertInterfaceToJSON(body)
	if err != nil {
		dataStr = fmt.Sprint(body)
	} else {
		dataStr = string(conv)
	}

	var (
		data = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   path,
			"data":       dataStr,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"method": method,
			"path":   path,
			"body":   data,
		}

		Logger.WithFields(logFields).Info("HTTP")
	}
}

// DetailLoggerInfo detail logger info
func DetailLoggerInfo(ctx context.Context, command, details string, logData interface{}) {
	var dataStr string
	conv, err := converter.ConvertInterfaceToJSON(logData)
	if err != nil {
		dataStr = fmt.Sprint(logData)
	} else {
		dataStr = string(conv)
	}

	var (
		data = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"data":       dataStr,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}

		Logger.WithFields(logFields).Info("Function")
	}
}

// DetailLoggerError detail logger error
func DetailLoggerError(ctx context.Context, command, details string, err ...interface{}) {
	var (
		reflectData = reflect.ValueOf(err)
		dataString  = fmt.Sprintf("%v", reflectData.Interface())
		data        = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"error":      dataString,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}

		Logger.WithFields(logFields).Error("Function")
	}
}

// DetailLoggerWarn detail logger warning
func DetailLoggerWarn(ctx context.Context, command, details string, warn ...interface{}) {
	var (
		reflectData = reflect.ValueOf(warn)
		dataString  = fmt.Sprintf("%v", reflectData.Interface())
		data        = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"warning":    dataString,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}

		Logger.WithFields(logFields).Warn("Function")
	}
}
