// Copyright (c) 2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"context"
	"testing"

	"github.com/uber/peloton/.gen/peloton/private/jobmgrsvc"
	jobmgrsvcmocks "github.com/uber/peloton/.gen/peloton/private/jobmgrsvc/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/yarpc/yarpcerrors"
)

type jobmgrActionsTestSuite struct {
	suite.Suite
	ctx    context.Context
	client Client

	ctrl         *gomock.Controller
	jobmgrClient *jobmgrsvcmocks.MockJobManagerServiceYARPCClient
}

func (suite *jobmgrActionsTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.jobmgrClient = jobmgrsvcmocks.NewMockJobManagerServiceYARPCClient(suite.ctrl)
	suite.ctx = context.Background()
	suite.client = Client{
		Debug:        false,
		jobmgrClient: suite.jobmgrClient,
		dispatcher:   nil,
		ctx:          suite.ctx,
	}
}

func (suite *jobmgrActionsTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestJobmgrActions(t *testing.T) {
	suite.Run(t, new(jobmgrActionsTestSuite))
}

// TestGetThrottledPodsSuccess tests getting list of throttled pods
func (suite *jobmgrActionsTestSuite) TestGetThrottledPodsSuccess() {
	suite.jobmgrClient.EXPECT().
		GetThrottledPods(gomock.Any(), gomock.Any()).
		Return(&jobmgrsvc.GetThrottledPodsResponse{}, nil)
	suite.NoError(suite.client.JobMgrGetThrottledPods())
}

// TestGetThrottledPodsFailure tests getting a failure when
// fetching list of throttled pods
func (suite *jobmgrActionsTestSuite) TestGetThrottledPodsFailure() {
	suite.jobmgrClient.EXPECT().
		GetThrottledPods(gomock.Any(), gomock.Any()).
		Return(nil, yarpcerrors.InternalErrorf("test error"))
	suite.Error(suite.client.JobMgrGetThrottledPods())
}
