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

package auth

import "github.com/matrix-org/gomatrixserverlib"

// IsServerAllowed returns true if there exists a event in authEvents
// which allows server to view this event. That is true when a client on the server
// can view the event. Otherwise returns false.
func IsServerAllowed(
	serverName gomatrixserverlib.ServerName,
	authEvents []gomatrixserverlib.Event,
) bool {
	for _, ev := range authEvents {
		membership, err := ev.Membership()
		if err != nil || membership != gomatrixserverlib.Join {
			continue
		}

		stateKey := ev.StateKey()
		if stateKey == nil {
			continue
		}

		_, domain, err := gomatrixserverlib.SplitID('@', *stateKey)
		if err != nil {
			continue
		}

		if domain == serverName {
			return true
		}
	}

	// TODO: Check if history visibility is shared and if the server is currently in the room
	return false
}
