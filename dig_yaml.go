/*
Copyright © 2020 esakat <esaka.tom@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package dig_yaml

import (
	"errors"
	"fmt"
	"reflect"
)

func DigYaml(y interface{}, keys ...interface{}) (interface{}, error) {

	if len(keys) == 0 {
		return nil, errors.New("keys is required")
	}

	head := keys[0]
	t := reflect.TypeOf(head).Kind()

	if len(keys) == 1 {
		if t != reflect.String {
			return nil, errors.New("key should be string")
		}
		item, ok := y.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("it's not yaml format")
		}
		for key, value := range item {
			if key == head.(string) {
				return value, nil
			}
		}
		errorMessage := fmt.Sprintf("key: %v not exists", head)
		return nil, errors.New(errorMessage)
	}

	// recursive
	switch t {
	case reflect.Int:
		// Array
		item, ok := y.([]interface{})
		if !ok {
			return nil, errors.New("it's not array")
		}
		return DigYaml(item[head.(int)], keys[1:]...)
	case reflect.String:
		// key
		item, ok := y.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("it's not yaml format")
		}
		for key, value := range item {
			if key == head.(string) {
				return DigYaml(value, keys[1:]...)
			}
		}

		errorMessage := fmt.Sprintf("key: %v not exists", head)
		return nil, errors.New(errorMessage)
	default:
		return nil, errors.New("keys should be string or int")
	}
}
