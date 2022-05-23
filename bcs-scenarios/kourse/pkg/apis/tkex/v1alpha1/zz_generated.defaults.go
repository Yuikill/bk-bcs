//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&GameStatefulSet{}, func(obj interface{}) { SetObjectDefaults_GameStatefulSet(obj.(*GameStatefulSet)) })
	scheme.AddTypeDefaultingFunc(&GameStatefulSetList{}, func(obj interface{}) { SetObjectDefaults_GameStatefulSetList(obj.(*GameStatefulSetList)) })
	return nil
}

func SetObjectDefaults_GameStatefulSet(in *GameStatefulSet) {
	SetDefaults_GameStatefulSet(in)
}

func SetObjectDefaults_GameStatefulSetList(in *GameStatefulSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_GameStatefulSet(a)
	}
}
