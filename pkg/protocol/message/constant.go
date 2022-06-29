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

package message

var MAGIC_CODE_BYTES = [2]byte{0xda, 0xda}

type (
	MessageType      int
	GettyRequestType byte
	GlobalStatus     int64
)

const (
	MessageTypeGlobalBegin              MessageType = iota + 1 // The constant TYPE_GLOBAL_BEGIN.
	MessageTypeGlobalBeginResult                               // The constant TYPE_GLOBAL_BEGIN_RESULT.
	MessageTypeBranchCommit                                    // The constant TYPE_BRANCH_COMMIT.
	MessageTypeBranchCommitResult                              // The constant TYPE_BRANCH_COMMIT_RESULT.
	MessageTypeBranchRollback                                  // The constant TYPE_BRANCH_ROLLBACK.
	MessageTypeGlobalCommit                                    // The constant TYPE_GLOBAL_COMMIT.
	MessageTypeGlobalCommitResult                              // The constant TYPE_GLOBAL_COMMIT_RESULT.
	MessageTypeGlobalRollback                                  // The constant TYPE_GLOBAL_ROLLBACK.
	MessageTypeBranchRollbackResult                            // The constant TYPE_BRANCH_ROLLBACK_RESULT.
	MessageTypeGlobalRollbackResult                            // The constant TYPE_GLOBAL_ROLLBACK_RESULT.
	MessageTypeBranchRegister                                  // The constant TYPE_BRANCH_REGISTER.
	MessageTypeBranchRegisterResult                            // The constant TYPE_BRANCH_REGISTER_RESULT.
	MessageTypeBranchStatusReport                              // The constant TYPE_BRANCH_STATUS_REPORT.
	MessageTypeBranchStatusReportResult                        // The constant TYPE_BRANCH_STATUS_REPORT_RESULT.
	MessageTypeGlobalStatus                                    // The constant TYPE_GLOBAL_STATUS.
	MessageTypeGlobalStatusResult                              // The constant TYPE_GLOBAL_STATUS_RESULT.
	MessageTypeGlobalReport                                    // The constant TYPE_GLOBAL_REPORT.
	MessageTypeGlobalReportResult                              // The constant TYPE_GLOBAL_REPORT_RESULT.
	MessageTypeGlobalLockQuery          MessageType = 21       // The constant TYPE_GLOBAL_LOCK_QUERY.
	MessageTypeGlobalLockQueryResult    MessageType = 22       // The constant TYPE_GLOBAL_LOCK_QUERY_RESULT.
	MessageTypeSeataMerge               MessageType = 59       // The constant TYPE_SEATA_MERGE.
	MessageTypeSeataMergeResult         MessageType = 60       // The constant TYPE_SEATA_MERGE_RESULT.
	MessageTypeRegClt                   MessageType = 101      // The constant TYPE_REG_CLT.
	MessageTypeRegCltResult             MessageType = 102      // The constant TYPE_REG_CLT_RESULT.
	MessageTypeRegRm                    MessageType = 103      // The constant TYPE_REG_RM.
	MessageTypeRegRmResult              MessageType = 104      // The constant TYPE_REG_RM_RESULT.
	MessageTypeRmDeleteUndolog          MessageType = 111      // The constant TYPE_RM_DELETE_UNDOLOG.
	MessageTypeHeartbeatMsg             MessageType = 120      // The constant TYPE_HEARTBEAT_MSG

)

const (
	VERSION = 1

	// MaxFrameLength max frame length
	MaxFrameLength = 8 * 1024 * 1024

	// V1HeadLength v1 head length
	V1HeadLength = 16

	GettyRequestTypeRequestSync       GettyRequestType = iota // Request message type
	GettyRequestTypeResponse                                  // Response message type
	GettyRequestTypeRequestOneway                             // Request which no need response
	GettyRequestTypeHeartbeatRequest                          // Heartbeat Request
	GettyRequestTypeHeartbeatResponse                         // Heartbeat Response
)

const (
	GlobalStatusUnKnown                 GlobalStatus = iota // Un known global status. Unknown
	GlobalStatusBegin                                       // The GlobalStatusBegin. PHASE 1: can accept new branch registering.
	GlobalStatusCommitting                                  // PHASE 2: Running Status: may be changed any time. Committing.
	GlobalStatusCommitRetrying                              // The Commit retrying. Retrying commit after a recoverable failure.
	GlobalStatusRollBacking                                 // Roll backing global status. Roll backing
	GlobalStatusRollbackRetrying                            // The Rollback retrying. Retrying rollback after a recoverable failure.
	GlobalStatusTimeoutRollBacking                          // The Timeout roll backing. roll backing since timeout
	GlobalStatusTimeoutRollbackRetrying                     // The Timeout rollback retrying. Retrying rollback  GlobalStatus = since timeout) after a recoverable failure.
	GlobalStatusAsyncCommitting                             // All branches can be async committed. The committing is NOT done yet, but it can be seen as committed for TM/RM client.
	GlobalStatusCommitted                                   // PHASE 2: Final Status: will NOT change any more. Finally: global transaction is successfully committed.
	GlobalStatusCommitFailed                                // The Commit failed.  Finally: failed to commit
	GlobalStatusRollBacked                                  // The roll backed. Finally: global transaction is successfully roll backed.
	GlobalStatusRollbackFailed                              // The Rollback failed. Finally: failed to rollback
	GlobalStatusTimeoutRollBacked                           // The Timeout roll backed. Finally: global transaction is successfully roll backed since timeout.
	GlobalStatusTimeoutRollbackFailed                       // The Timeout rollback failed. Finally: failed to rollback since timeout
	GlobalStatusFinished                                    // The Finished. Not managed in session MAP any more
)
