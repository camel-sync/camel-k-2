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

package maven

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/apache/camel-k/v2/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestRunAndLogErrorMvn(t *testing.T) {
	mavenCmd, ok := os.LookupEnv("MAVEN_CMD")
	if !ok {
		mavenCmd = "mvn"
	}

	cmd := exec.CommandContext(context.Background(), mavenCmd, "package", "-B")
	err := util.RunAndLog(context.Background(), cmd, mavenLogHandler, mavenLogHandler)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "[ERROR] The goal you specified requires a project to execute but there is no POM in this directory")
}
