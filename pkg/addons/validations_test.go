/*
Copyright 2019 The Kubernetes Authors All rights reserved.

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

package addons

import "testing"

func TestIsAddonValid(t *testing.T) {
	tests := []struct {
		description string
		name        string
		isValid     bool
	}{
		{
			description: "valid addon",
			name:        "gvisor",
			isValid:     true,
		}, {
			description: "invalid addon",
			name:        "invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, valid := isAddonValid(test.name)
			if test.isValid != valid {
				t.Fatalf("expected: %t\nactual:%t\n", test.isValid, valid)
			}
		})
	}
}
