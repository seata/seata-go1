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

package resource

import (
	"context"
	"sync"

	"github.com/seata/seata-go/pkg/protocol/branch"
)

// Resource that can be managed by Resource Manager and involved into global transaction
type Resource interface {
	GetResourceGroupId() string
	GetResourceId() string
	GetBranchType() branch.BranchType
}

// ResourceManagerInbound Control a branch transaction commit or rollback
type ResourceManagerInbound interface {
	BranchCommit(ctx context.Context, branchType branch.BranchType, xid string, branchId int64, resourceId string, applicationData []byte) (branch.BranchStatus, error)  // BranchCommit Commit a branch transaction
	BranchRollback(ctx context.Context, ranchType branch.BranchType, xid string, branchId int64, resourceId string, applicationData []byte) (branch.BranchStatus, error) // BranchRollback Rollback a branch transaction
}

// ResourceManagerOutbound Resource Manager: send outbound request to TC
type ResourceManagerOutbound interface {
	BranchRegister(ctx context.Context, ranchType branch.BranchType, resourceId, clientId, xid, applicationData, lockKeys string) (int64, error)         // BranchRegister Branch register long
	BranchReport(ctx context.Context, ranchType branch.BranchType, xid string, branchId int64, status branch.BranchStatus, applicationData string) error // BranchReport Branch report
	LockQuery(ctx context.Context, ranchType branch.BranchType, resourceId, xid, lockKeys string) (bool, error)                                          // LockQuery Lock query boolean
}

// ResourceManager Resource Manager: common behaviors
type ResourceManager interface {
	ResourceManagerInbound
	ResourceManagerOutbound

	RegisterResource(resource Resource) error   // RegisterResource Register a Resource to be managed by Resource Manager
	UnregisterResource(resource Resource) error //  UnregisterResource Unregister a Resource from the Resource Manager
	GetManagedResources() *sync.Map             // GetManagedResources Get all resources managed by this manager
	GetBranchType() branch.BranchType           // GetBranchType Get the BranchType
}

type ResourceManagerGetter interface {
	GetResourceManager() ResourceManager
}
