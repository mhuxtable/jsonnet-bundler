// Copyright 2018 jsonnet-bundler authors
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

package main

import (
	"testing"

	"github.com/jsonnet-bundler/jsonnet-bundler/spec"
	"github.com/stretchr/testify/assert"
)

func TestParseDepedency(t *testing.T) {
	tests := []struct {
		name string
		path string
		want *spec.Dependency
	}{
		{
			name: "Empty",
			path: "",
			want: nil,
		},
		{
			name: "Invalid",
			path: "github.com/foo",
			want: nil,
		},
		{
			name: "GitHub",
			path: "github.com/jsonnet-bundler/jsonnet-bundler",
			want: &spec.Dependency{
				Name: "jsonnet-bundler",
				Source: spec.Source{
					GitSource: &spec.GitSource{
						Remote: "https://github.com/jsonnet-bundler/jsonnet-bundler",
						Subdir: "",
					},
				},
				Version: "master",
			},
		},
		{
			name: "SSH",
			path: "git+ssh://git@github.com:jsonnet-bundler/jsonnet-bundler.git",
			want: &spec.Dependency{
				Name: "jsonnet-bundler",
				Source: spec.Source{
					GitSource: &spec.GitSource{
						Remote: "git@github.com:jsonnet-bundler/jsonnet-bundler",
						Subdir: "",
					},
				},
				Version: "master",
			},
		},
	}
	for _, tt := range tests {
		_ = t.Run(tt.name, func(t *testing.T) {
			dependency := parseDependency(tt.path)

			if tt.path == "" {
				assert.Nil(t, dependency)
			} else {
				assert.Equal(t, tt.want, dependency)
			}
		})
	}
}
