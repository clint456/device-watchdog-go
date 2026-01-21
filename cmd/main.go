// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2021 IOTech Ltd
// Copyright (C) 2025 Sichuan Huiyuan Optical Communication Co., Ltd. Clinton Luo
// SPDX-License-Identifier: Apache-2.0

package main

import (
	demo "device-demo-go"
	"device-demo-go/internal/driver"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/startup"
)

const (
	serviceName string = "device-demo"
)

func main() {
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, demo.Version, sd)
}
