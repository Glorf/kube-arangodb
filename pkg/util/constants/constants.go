//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
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
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package constants

const (
	EnvOperatorNodeName     = "MY_NODE_NAME"
	EnvOperatorPodName      = "MY_POD_NAME"
	EnvOperatorPodNamespace = "MY_POD_NAMESPACE"
	EnvOperatorPodIP        = "MY_POD_IP"

	EnvArangodJWTSecret          = "ARANGOD_JWT_SECRET"          // Contains JWT secret for the ArangoDB cluster
	EnvArangoSyncMonitoringToken = "ARANGOSYNC_MONITORING_TOKEN" // Constains monitoring token for ArangoSync servers

	SecretEncryptionKey = "key"   // Key in a Secret.Data used to store an 32-byte encryption key
	SecretKeyJWT        = "token" // Key inside a Secret used to hold a JW token
	SecretKeyMonitoring = "token" // Key inside a Secret used to hold a monitoring token

	SecretCACertificate = "ca.crt" // Key in Secret.data used to store a PEM encoded CA certificate (public key)
	SecretCAKey         = "ca.key" // Key in Secret.data used to store a PEM encoded CA private key

	SecretTLSKeyfile = "tls.keyfile" // Key in Secret.data used to store a PEM encoded TLS certificate in the format used by ArangoDB (`--ssl.keyfile`)

	FinalizerPodDrainDBServer = "dbserver.database.arangodb.com/drain"       // Finalizer added to DBServers, indicating the need for draining that dbserver
	FinalizerPodAgencyServing = "agent.database.arangodb.com/agency-serving" // Finalizer added to Agents, indicating the need for keeping enough agents alive
	FinalizerPVCMemberExists  = "pvc.database.arangodb.com/member-exists"    // Finalizer added to PVCs, indicating the need to keep is as long as its member exists
)
