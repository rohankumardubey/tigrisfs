// Copyright 2015 - 2017 Ka-Hing Cheung
// Copyright 2021 Yandex LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows

package cfg

import (
	"os"

	"github.com/yandex-cloud/geesefs/lib"
	"github.com/yandex-cloud/geesefs/log"
)

func InitLoggers(flags *FlagStorage) {
	lf := flags.LogFile
	if lf == "" {
		lf = "stderr"
		if !flags.Foreground {
			lf = "syslog"
		}
	}

	log.InitLoggerRedirect(lf)

	log.DefaultLogConfig = &log.LogConfig{
		Level:  flags.LogLevel,
		Format: flags.LogFormat,
	}

	if (lib.IsTTY(os.Stdout) || lib.IsTTY(os.Stderr)) && log.DefaultLogConfig.Format == "" && lf == "stderr" {
		log.DefaultLogConfig.Format = "console"
	}

	log.DefaultLogConfig.Color = true
	if flags.NoLogColor {
		log.DefaultLogConfig.Color = false
	}

	log.SetLoggersConfig(log.DefaultLogConfig)

	log.DumpLoggers("InitLoggers")
}
