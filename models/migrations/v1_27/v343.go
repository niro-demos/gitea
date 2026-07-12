// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_27

import "gitea.dev/models/db"

func BindTemporaryUploads(x db.EngineMigration) error {
	type Upload struct {
		UploaderID int64 `xorm:"INDEX NOT NULL DEFAULT 0"`
		RepoID     int64 `xorm:"INDEX NOT NULL DEFAULT 0"`
	}
	type PackageBlobUpload struct {
		OwnerID int64  `xorm:"INDEX(owner_image) NOT NULL DEFAULT 0"`
		Image   string `xorm:"INDEX(owner_image) NOT NULL DEFAULT ''"`
	}

	return x.Sync(new(Upload), new(PackageBlobUpload))
}
