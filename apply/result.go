// Copyright © 2016 Asteris, LLC
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

package apply

import "github.com/asteris-llc/converge/plan"

// Result of application
type Result struct {
	Ran    bool
	Status string

	Plan *plan.Result
}

// Fields returns the fields that changed
func (r *Result) Fields() map[string][2]string {
	previous := "<unknown>"
	if r.Plan != nil {
		previous = r.Plan.Status
	}

	return map[string][2]string{
		"state": [2]string{previous, r.Status},
	}
}

// HasChanges indicates if this result ran
func (r *Result) HasChanges() bool { return r.Ran }

// Error always returns nil since application results never fail
func (r *Result) Error() error { return nil }
