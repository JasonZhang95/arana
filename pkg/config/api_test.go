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

package config_test

import (
	"fmt"
	"testing"
)

import (
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

import (
	"github.com/arana-db/arana/pkg/config"
	"github.com/arana-db/arana/testdata"
)

func TestGetStoreOperate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// mockStore := NewMockStoreOperate(ctrl)
	tests := []struct {
		name    string
		want    config.StoreOperate
		wantErr assert.ErrorAssertionFunc
	}{
		{"GetStoreOperate_1", nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.GetStoreOperate()
			if !tt.wantErr(t, err, fmt.Sprintf("GetStoreOperate()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetStoreOperate()")
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		name    string
		options map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{"Init_1", args{"file", nil}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, config.Init(tt.args.name, tt.args.options), fmt.Sprintf("Init(%v, %v)", tt.args.name, tt.args.options))
		})
	}
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := testdata.NewMockStoreOperate(ctrl)
	mockStore.EXPECT().Name().Times(2).Return("nacos")
	type args struct {
		s config.StoreOperate
	}
	tests := []struct {
		name string
		args args
	}{
		{"Register", args{mockStore}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Register(tt.args.s)
		})
	}
}

func Test_api(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFileStore := testdata.NewMockStoreOperate(ctrl)
	mockEtcdStore := testdata.NewMockStoreOperate(ctrl)
	mockFileStore.EXPECT().Name().Times(2).Return("file")
	mockEtcdStore.EXPECT().Name().Times(2).Return("etcd")
	config.Register(mockFileStore)
	config.Register(mockEtcdStore)

	mockFileStore2 := testdata.NewMockStoreOperate(ctrl)
	mockFileStore2.EXPECT().Name().AnyTimes().Return("file")
	assert.Panics(t, func() {
		config.Register(mockFileStore2)
	}, "StoreOperate=[file] already exist")
}

func Test_Init(t *testing.T) {
	options := make(map[string]interface{}, 0)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFileStore := testdata.NewMockStoreOperate(ctrl)
	mockFileStore.EXPECT().Name().Times(2).Return("fake")
	mockFileStore.EXPECT().Init(options).Return(nil)
	err := config.Init("fake", options)
	assert.Error(t, err)

	config.Register(mockFileStore)
	err = config.Init("fake", options)
	assert.NoError(t, err)

	store, err := config.GetStoreOperate()
	assert.NoError(t, err)
	assert.NotNil(t, store)
}
