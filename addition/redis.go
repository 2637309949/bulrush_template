/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:50:15
 * @modify date 2019-01-16 20:50:15
 * @desc [description]
 */

package addition

import (
	"github.com/2637309949/bulrush-addition/redis"
	"github.com/2637309949/bulrush-template/conf"
)

// Redis application redis store
var Redis = redis.New(conf.Cfg)