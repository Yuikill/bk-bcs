# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""
import mock
import pytest

from backend.tests.testing_utils.base import generate_random_string
from backend.tests.testing_utils.mocks.paas_cc import StubPaaSCCClient


@pytest.fixture
def namespace_name():
    return generate_random_string(16)


@pytest.fixture
def template_id():
    """生成一个随机模板集 ID"""
    return generate_random_string(32)


@pytest.fixture
def template_name():
    """生成一个随机模板集 name"""
    return generate_random_string(16)


@pytest.fixture(autouse=True)
def patch_paas_cc_client():
    with mock.patch('backend.iam.permissions.resources.namespace.PaaSCCClient', new=StubPaaSCCClient), mock.patch(
        'backend.iam.permissions.resources.namespace.get_client_access_token',
        new=lambda *args, **kwargs: {'access_token': generate_random_string(16)},
    ):
        yield
