// Copyright 2021 The etcd Authors
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

package integration

import (
	"os"
	"path/filepath"
	"testing"

	"go.etcd.io/etcd/pkg/v3/testutil"
)

func BeforeTest(t testing.TB) {
	testutil.BeforeTest(t)

	previousWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	os.Chdir(t.TempDir())
	t.Cleanup(func() {
		os.Chdir(previousWD)
	})

}

func MustAbsPath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return abs
}