// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_27

import "gitea.dev/models/db"

func AddRepositoryToPackageBlobUpload(x db.EngineMigration) error {
	type PackageBlobUpload struct {
		OwnerID int64  `xorm:"INDEX(owner_image) NOT NULL DEFAULT 0"`
		Image   string `xorm:"INDEX(owner_image) NOT NULL DEFAULT ''"`
	}

	return x.Sync(new(PackageBlobUpload))
}
