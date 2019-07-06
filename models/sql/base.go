// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sql

import (
	gormext "github.com/2637309949/bulrush-addition/gorm"
)

// Base common fields
type Base struct {
	gormext.Model
	Creator    *User
	CreatorID  uint `gorm:"comment:创建人ID;"`
	Modifier   *User
	ModifierID uint `gorm:"comment:修改人ID;"`
}
