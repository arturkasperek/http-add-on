// /*
// Copyright 2023 The KEDA Authors.
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
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: operator/generated/listers/http/v1alpha1/expansion_generated.go
//
// Generated by this command:
//
//	mockgen -copyright_file=hack/boilerplate.go.txt -destination=operator/generated/listers/http/v1alpha1/mock/expansion_generated.go -package=mock -source=operator/generated/listers/http/v1alpha1/expansion_generated.go
//

// Package mock is a generated GoMock package.
package mock

import (
	gomock "go.uber.org/mock/gomock"
)

// MockHTTPScaledObjectListerExpansion is a mock of HTTPScaledObjectListerExpansion interface.
type MockHTTPScaledObjectListerExpansion struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScaledObjectListerExpansionMockRecorder
	isgomock struct{}
}

// MockHTTPScaledObjectListerExpansionMockRecorder is the mock recorder for MockHTTPScaledObjectListerExpansion.
type MockHTTPScaledObjectListerExpansionMockRecorder struct {
	mock *MockHTTPScaledObjectListerExpansion
}

// NewMockHTTPScaledObjectListerExpansion creates a new mock instance.
func NewMockHTTPScaledObjectListerExpansion(ctrl *gomock.Controller) *MockHTTPScaledObjectListerExpansion {
	mock := &MockHTTPScaledObjectListerExpansion{ctrl: ctrl}
	mock.recorder = &MockHTTPScaledObjectListerExpansionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScaledObjectListerExpansion) EXPECT() *MockHTTPScaledObjectListerExpansionMockRecorder {
	return m.recorder
}

// MockHTTPScaledObjectNamespaceListerExpansion is a mock of HTTPScaledObjectNamespaceListerExpansion interface.
type MockHTTPScaledObjectNamespaceListerExpansion struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScaledObjectNamespaceListerExpansionMockRecorder
	isgomock struct{}
}

// MockHTTPScaledObjectNamespaceListerExpansionMockRecorder is the mock recorder for MockHTTPScaledObjectNamespaceListerExpansion.
type MockHTTPScaledObjectNamespaceListerExpansionMockRecorder struct {
	mock *MockHTTPScaledObjectNamespaceListerExpansion
}

// NewMockHTTPScaledObjectNamespaceListerExpansion creates a new mock instance.
func NewMockHTTPScaledObjectNamespaceListerExpansion(ctrl *gomock.Controller) *MockHTTPScaledObjectNamespaceListerExpansion {
	mock := &MockHTTPScaledObjectNamespaceListerExpansion{ctrl: ctrl}
	mock.recorder = &MockHTTPScaledObjectNamespaceListerExpansionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScaledObjectNamespaceListerExpansion) EXPECT() *MockHTTPScaledObjectNamespaceListerExpansionMockRecorder {
	return m.recorder
}
