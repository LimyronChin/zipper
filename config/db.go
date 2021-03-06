// Copyright (C) 2017, Zipper Team.  All rights reserved.
//
// This file is part of zipper
//
// The zipper is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The zipper is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"path/filepath"

	"github.com/zipper-project/zipper/common/db"
)

// DBConfig returns configurations for database
func DBConfig() *db.Config {
	dbConfig := db.DefaultConfig()
	dbConfig.DbPath = filepath.Join(DataDir, ChainDataDirName)
	dbConfig.Columnfamilies = getStringSlice("db.columnfamilies", dbConfig.Columnfamilies)
	dbConfig.KeepLogFileNumber = getInt("db.keepLogFileNumber", dbConfig.KeepLogFileNumber)
	dbConfig.MaxLogFileSize = getInt("db.maxLogFileSize", dbConfig.MaxLogFileSize)
	dbConfig.LogLevel = getString("db.loglevel", dbConfig.LogLevel)
	return dbConfig
}
