// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package componentsuite

import (
	"github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/options"
)

// ComponentTestSuite contains metadata and test cases for running tests.
type ComponentTestSuite struct {
	// ComponentFile contains a path to a component file or component builder
	// file.
	ComponentFile string `json:"componentFile"`

	// RootDirectory is the path to a root-directory to run the tests. If not
	// specified, sets the root directory to the test-suite file directory.
	RootDirectory string `json:"rootDirectory"`

	// TestCases contains a list of component TestCases.
	TestCases []*TestCase `json:"testCases"`
}

// TestCase contains the schema expected for the test-cases.
type TestCase struct {
	// Description of the test.
	Description string `json:"description"`

	// Build contains parameters for the build-phase. This is roughly equivalent
	// to 'bundlectl build'
	Build Build `json:"build"`

	// Apply contains parameters for the apply-phase. This is roughly equivalent
	// to 'bundlectl apply'
	Apply Apply `json:"apply"`

	// Expect contains expectations to check against.
	Expect Expect `json:"expect"`
}

// Build contains build parameters
type Build struct {
	// Options for the PatchTemplate build process
	Options options.JSONOptions `json:"options"`
}

// Apply contains build paramaters
type Apply struct {
	// Options for apply process
	Options options.JSONOptions `json:"options"`
}

// Expect contains expectations that should be filled.
type Expect struct {
	// Objects contains expectations for objects.
	Objects []ObjectCheck `json:"objects"`

	// CanKubeDeserialize ensures that all objects can be deserialized as
	// concrete Kubernetes objects, using the built in schema.
	CanKubeDeserialize bool `json:"CanKubeDeserialize"`

	// BuildErrSubstr indicates a substring that's expected to be in an error in
	// the build-process.
	BuildErrSubstr string `json:"buildErrSubstr"`

	// ApplyErrSubstr indicates a substring that's expected to be in an error in
	// the apply-process.
	ApplyErrSubstr string `json:"applyErrSubstr"`
}

// ObjectCheck contains checks for a specific Object. Kind, and Name are used to find
// objects. Expects exactly one object to match.
type ObjectCheck struct {
	// Kind of the object (required).
	Kind string `json:"kind"`

	// Name of the object (required).
	Name string `json:"name"`

	// FindSubstrs contains a list of substrings that are expected to be found in
	// the object after the apply-phase.
	FindSubstrs []string `json:"findSubstrs"`

	// NotFindSubstrs contains a list of substrings that are not expecetd to be
	// found in the object after the apply-phase.
	NotFindSubstrs []string `json:"notFindSubstrs"`
}
