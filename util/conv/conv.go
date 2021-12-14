// Copyright © 2016 SurrealDB Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/surrealdb/surrealdb/sql"
)

func toNumber(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		val = 0.0
		err = fmt.Errorf("Expected a number, but found '%v'", str)

	}
	return float64(val), err
}

func toBoolean(str string) (bool, error) {
	val, err := strconv.ParseBool(str)
	if err != nil {
		val = false
		err = fmt.Errorf("Expected a boolean, but found '%v'", str)
	}
	return bool(val), err
}

// --------------------------------------------------

func MustBe(t, obj interface{}) (val interface{}) {
	switch t {
	default:
		return obj
	case "array":
		return MustBeArray(obj)
	case "object":
		return MustBeObject(obj)
	}
}

func MustBeArray(obj interface{}) (val interface{}) {
	if now, ok := obj.([]interface{}); ok {
		val = now
	} else {
		val = make([]interface{}, 0)
	}
	return
}

func MustBeObject(obj interface{}) (val interface{}) {
	if now, ok := obj.(map[string]interface{}); ok {
		val = now
	} else {
		val = make(map[string]interface{})
	}
	return
}

// --------------------------------------------------

func MightBe(obj interface{}) (val interface{}, ok bool) {
	switch now := obj.(type) {
	case string:
		if val, ok := MightBeDatetime(now); ok {
			return val, ok
		}
		if val, ok := MightBeRecord(now); ok {
			return val, ok
		}
	}
	return obj, false
}

func MightBeDatetime(obj string) (val interface{}, ok bool) {
	if val, err := time.Parse(time.RFC3339, obj); err == nil {
		return val, true
	}
	return nil, false
}

func MightBeRecord(obj string) (val interface{}, ok bool) {
	if val := sql.ParseThing(obj); val != nil {
		return val, true
	}
	return nil, false
}

// --------------------------------------------------

func ConvertTo(t, k string, obj interface{}) (val interface{}, err error) {
	switch t {
	default:
		return obj, nil
	case "array":
		return ConvertToArray(obj)
	case "object":
		return ConvertToObject(obj)
	case "string":
		return ConvertToString(obj)
	case "number":
		return ConvertToNumber(obj)
	case "boolean":
		return ConvertToBoolean(obj)
	case "datetime":
		return ConvertToDatetime(obj)
	case "record":
		return ConvertToRecord(obj, k)
	}
}

func ConvertToArray(obj interface{}) (val []interface{}, err error) {
	if now, ok := obj.([]interface{}); ok {
		val = now
	} else {
		err = fmt.Errorf("Expected an array, but found '%v'", obj)
	}
	return
}

func ConvertToObject(obj interface{}) (val map[string]interface{}, err error) {
	if now, ok := obj.(map[string]interface{}); ok {
		val = now
	} else {
		err = fmt.Errorf("Expected an object, but found '%v'", obj)
	}
	return
}

func ConvertToString(obj interface{}) (val string, err error) {
	switch now := obj.(type) {
	case string:
		return now, err
	case []interface{}, map[string]interface{}:
		return val, fmt.Errorf("Expected a string, but found '%v'", obj)
	default:
		return fmt.Sprintf("%v", obj), err
	}
}

func ConvertToNumber(obj interface{}) (val float64, err error) {
	switch now := obj.(type) {
	case int64:
		return float64(now), err
	case float64:
		return float64(now), err
	case string:
		return toNumber(now)
	default:
		return toNumber(fmt.Sprintf("%v", obj))
	}
}

func ConvertToBoolean(obj interface{}) (val bool, err error) {
	switch now := obj.(type) {
	case int64:
		return now > 0, err
	case float64:
		return now > 0, err
	case string:
		return toBoolean(now)
	default:
		return toBoolean(fmt.Sprintf("%v", obj))
	}
}

func ConvertToDatetime(obj interface{}) (val time.Time, err error) {
	switch now := obj.(type) {
	case time.Time:
		val = now
	case string:
		val, err = time.Parse(time.RFC3339Nano, now)
	default:
		err = fmt.Errorf("Expected a datetime, but found '%v'", obj)
	}
	return
}

func ConvertToRecord(obj interface{}, tb string) (val *sql.Thing, err error) {
	switch now := obj.(type) {
	case *sql.Thing:
		switch tb {
		case now.TB:
			val = now
		case "":
			val = now
		default:
			err = fmt.Errorf("Expected a record of type '%s', but found '%v'", tb, obj)
		}
	case string:
		if val = sql.ParseThing(now); val == nil {
			err = fmt.Errorf("Expected a record of type '%s', but found '%v'", tb, obj)
		}
	default:
		switch tb {
		default:
			err = fmt.Errorf("Expected a record of type '%s', but found '%v'", tb, obj)
		case "":
			err = fmt.Errorf("Expected a record of any type, but found '%v'", obj)
		}
	}
	return
}
