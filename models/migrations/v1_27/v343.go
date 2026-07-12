// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_27

import "gitea.dev/models/db"

func BindRepositoryUploads(x db.EngineMigration) error {
	type Upload struct {
		UploaderID int64 `xorm:"INDEX NOT NULL DEFAULT 0"`
		RepoID     int64 `xorm:"INDEX NOT NULL DEFAULT 0"`
	}
	return x.Sync(new(Upload))
}
