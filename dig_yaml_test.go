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

import "testing"

func TestDigYaml(t *testing.T) {
	testdata := map[interface{}]interface{}{
		"hoge": "foo",
		"mo":   true,
		"2": map[interface{}]interface{}{
			"hoge": 2,
			"fuga": 3,
		},
		"sample": []interface{}{
			map[interface{}]interface{}{
				"test": 1,
			},
			map[interface{}]interface{}{
				"foo": 2,
			},
			map[interface{}]interface{}{
				"mac": "testmsg",
			},
		},
	}

	actual1, err := DigYaml(testdata, "2", "hoge")
	if err != nil {
		t.Fatalf("failed to dig yaml: %v", err)
	}
	expected1 := 2

	if actual1 != expected1 {
		t.Fatalf("expected: %v, got: %v", expected1, actual1)
	}

	actual2, err := DigYaml(testdata, "sample", 2, "mac")
	if err != nil {
		t.Fatalf("failed to dig yaml: %v", err)
	}
	expected2 := "testmsg"

	if actual2 != expected2 {
		t.Fatalf("expected: %v, got: %v", expected2, actual2)
	}
}
