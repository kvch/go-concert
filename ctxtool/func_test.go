// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ctxtool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithFunc(t *testing.T) {
	t.Run("never executed on backgound context", func(t *testing.T) {
		count := 0
		ctx := WithFunc(context.Background(), func() { count++ })
		assert.NotNil(t, ctx)
		assert.Equal(t, 0, count)
	})

	t.Run("executed func on cancel", func(t *testing.T) {
		done := make(chan struct{})
		count := 0
		ctx, cancelFn := context.WithCancel(context.Background())
		ctx = WithFunc(ctx, func() { close(done); count++ })
		cancelFn()
		<-done
		assert.Equal(t, 1, count)
	})
}