/*
 * Copyright (c) 2026, WSO2 LLC. (https://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package templateengine

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestPreRender_ResolvesEnvInRawYAMLBytes(t *testing.T) {
	raw := []byte(`url: {{ env "TEST_URL" }}`)
	fm := template.FuncMap{
		"env": func(key string) string {
			return map[string]string{"TEST_URL": "https://example.com"}[key]
		},
	}

	got, err := PreRender(raw, fm)

	require.NoError(t, err)
	assert.Equal(t, "url: https://example.com", string(got))
	
	var parsed map[string]any
	require.NoError(t, yaml.Unmarshal(got, &parsed), "rendered output must be valid YAML")
	assert.Equal(t, "https://example.com", parsed["url"])
}
