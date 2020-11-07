// Copyright 2020 astaxie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package beego

var (
	BuildVersion     string
	BuildGitRevision string
	BuildStatus      string
	BuildTag         string
	BuildTime        string

	GoVersion string

	GitBranch string
)

const (
	// VERSION represent beego web framework version.
	VERSION = "2.0.0-alpha"
)
