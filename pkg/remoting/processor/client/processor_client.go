/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"github.com/seata/seata-go/pkg/protocol/message"
	"github.com/seata/seata-go/pkg/remoting/getty"
)
// Init is initialization of the current package
func Init() {
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_HeartbeatMsg, &clientHeartBeatProcesson{})
	clientOnResponseProcessor := &clientOnResponseProcessor{}
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_SeataMergeResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_BranchRegisterResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_BranchStatusReportResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalLockQueryResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_RegRmResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalBeginResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalCommitResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalReportResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalRollbackResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_GlobalStatusResult, clientOnResponseProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_RegCltResult, clientOnResponseProcessor)
	rmBranchCommitProcessor := &rmBranchCommitProcessor{}
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_BranchCommit, rmBranchCommitProcessor)
	rmBranchRollbackProcessor := &rmBranchRollbackProcessor{}
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_BranchRollback, rmBranchRollbackProcessor)
	getty.GetGettyClientHandlerInstance().RegisterProcessor(message.MessageType_BranchCommit, rmBranchCommitProcessor)

}
