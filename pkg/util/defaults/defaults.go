/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package defaults

// ***********************
//  DO NOT EDIT THIS FILE
// ***********************

const (
	// Version --
	Version = "2.0.0-SNAPSHOT"

	// DefaultRuntimeVersion --
	DefaultRuntimeVersion = "1.18.0-SNAPSHOT"

	// BuildahVersion --
	BuildahVersion = "1.23.3"

	// KanikoVersion --
	KanikoVersion = "0.17.1"

	// baseImage --
	baseImage = "docker.io/eclipse-temurin:11"

	// LocalRepository --
	LocalRepository = "/etc/maven/m2"

	// DefaultPVC --
	DefaultPVC = "camel-k-pvc"

	// ImageName --
	ImageName = "docker.io/apache/camel-k"

	// installDefaultKamelets --
	installDefaultKamelets = true
)

// GitCommit must be provided during application build
var GitCommit string
