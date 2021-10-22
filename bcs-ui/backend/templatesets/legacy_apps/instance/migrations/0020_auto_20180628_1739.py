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
# Generated by Django 1.11.5 on 2018-06-28 09:39
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('instance', '0019_auto_20180409_1119'),
    ]

    operations = [
        migrations.AlterField(
            model_name='instanceconfig',
            name='category',
            field=models.CharField(choices=[('application', 'Application'), ('deployment', 'Deplpyment'), ('service', 'Service'), ('configmap', 'ConfigMap'), ('secret', 'Secret'), ('K8sSecret', 'K8sSecret'), ('K8sConfigMap', 'K8sConfigMap'), ('K8sService', 'K8sService'), ('K8sDeployment', 'K8sDeployment'), ('K8sDaemonSet', 'K8sDaemonSet'), ('K8sJob', 'K8sJob'), ('K8sStatefulSet', 'K8sStatefulSet')], max_length=32, verbose_name='资源类型'),
        ),
        migrations.AlterField(
            model_name='instanceconfig',
            name='instance_id',
            field=models.IntegerField(db_index=True, verbose_name='关联的 VersionInstance ID'),
        ),
    ]
