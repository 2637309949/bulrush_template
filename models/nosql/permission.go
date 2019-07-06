// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package nosql

import (
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/globalsign/mgo/bson"
)

// Permission info
type Permission struct {
	Base `bson:",inline"`
	Code string        `bson:"code,comment:编码,"`
	Name string        `bson:"name,comment:名称,"`
	Pid  bson.ObjectId `bson:"pid,comment:父级ID,"`
	Type uint          `bson:"type,comment:类型,enum:一级菜单=1 二级菜单=2 三级菜单=3 按钮=4 自定义=5,"`
}

var _ = addition.MGOExt.Register(&mgoext.Profile{
	DB:        "test",
	Name:      "permission",
	Reflector: &Permission{},
})
