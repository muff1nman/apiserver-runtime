/*
Copyright 2020 The Kubernetes Authors.

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

// Package loopback is a set of utilities for apiserver loopback connections.
package loopback

import (
	"sync"

	"k8s.io/client-go/rest"
)

var setLoopbackClientOnce sync.Once
var loopbackClientConfig *rest.Config

// SetLoopbackClientConfig provides loopback client config for one time
func SetLoopbackClientConfig(c *rest.Config) {
	setLoopbackClientOnce.Do(func() {
		loopbackClientConfig = c
	})
}

// GetLoopbackClientConfig gets loopback client config
func GetLoopbackClientConfig() *rest.Config {
	return loopbackClientConfig
}
