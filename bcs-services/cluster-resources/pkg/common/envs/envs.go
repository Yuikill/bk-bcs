/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package envs

import (
	"os"
	"path/filepath"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util/envx"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util/path"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util/stringx"
)

// 以下变量值可通过环境变量指定
var (
	// BCSApiGWHost 容器服务网关 Host
	BCSApiGWHost = os.Getenv("BCS_API_GW_HOST")
	// BCSApiGWAuthToken 网关 Auth Token
	BCSApiGWAuthToken = os.Getenv("BCS_API_GW_AUTH_TOKEN")
	// ExampleFileBaseDir Example 配置文件目录
	ExampleFileBaseDir = envx.Get(
		"EXAMPLE_FILE_BASE_DIR", filepath.Dir(filepath.Dir(path.GetCurPKGPath()))+"/resource/example",
	)
)

// 以下变量值可通过环境变量指定（仅用于单元测试）
var (
	// TestProjectID 单测指定的项目 ID
	TestProjectID = envx.Get("TEST_PROJECT_ID", stringx.Rand(32, ""))
	// TestClusterID 单测指定的集群 ID
	TestClusterID = envx.Get("TEST_CLUSTER_ID", "BCS-K8S-T"+stringx.Rand(5, "1234567890"))
	// TestNamespace 单测指定的命名空间
	TestNamespace = envx.Get("TEST_NAMESPACE", "default")
)