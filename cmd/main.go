// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2021 IOTech Ltd
// Copyright (C) 2025 Sichuan Huiyuan Optical Communication Co., Ltd. Clinton Luo
// SPDX-License-Identifier: Apache-2.0

package main

import (
	watchdog "device-watchdog-go"
	"device-watchdog-go/internal/driver"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/startup"
)

const (
	serviceName string = "device-watchdog"
)

func main() {
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, watchdog.Version, sd)
}
