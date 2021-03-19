package ssecurity

import (
	"encoding/json"

	"github.com/go-redis/redis/v7"
	log "github.com/syncfuture/go/slog"
	"github.com/syncfuture/go/sproto"
	"github.com/syncfuture/go/sredis"
	"github.com/syncfuture/go/u"
)

type RedisPermissionProvider struct {
	redis         redis.Cmdable
	PermissionKey string
}

func NewRedisPermissionProvider(permissionKey string, config *sredis.RedisConfig) IPermissionProvider {
	if permissionKey == "" {
		log.Fatal("permissionKey key cannot be empty")
	}

	r := new(RedisPermissionProvider)

	r.redis = sredis.NewClient(config)

	r.PermissionKey = permissionKey

	return r
}

// *******************************************************************************************************************************
// Permission
func (x *RedisPermissionProvider) CreatePermission(in *sproto.PermissionDTO) error {
	j, err := json.Marshal(in)
	if err != nil {
		return err
	}

	cmd := x.redis.HSet(x.PermissionKey, in.ID, j)
	return cmd.Err()
}
func (x *RedisPermissionProvider) GetPermission(id string) (*sproto.PermissionDTO, error) {
	cmd := x.redis.HGet(x.PermissionKey, id)
	err := cmd.Err()
	if err != nil {
		return nil, err
	}

	r := new(sproto.PermissionDTO)
	j, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(j), r)
	u.LogError(err)
	return r, err
}
func (x *RedisPermissionProvider) UpdatePermission(in *sproto.PermissionDTO) error {
	j, err := json.Marshal(in)
	if err != nil {
		return err
	}

	cmd := x.redis.HSet(x.PermissionKey, in.ID, j)
	return cmd.Err()
}
func (x *RedisPermissionProvider) RemovePermission(id string) error {
	cmd := x.redis.HDel(x.PermissionKey, id)
	return cmd.Err()
}
func (x *RedisPermissionProvider) GetPermissions() (map[string]*sproto.PermissionDTO, error) {
	cmd := x.redis.HGetAll(x.PermissionKey)
	r, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]*sproto.PermissionDTO, len(r))
	for key, value := range r {
		dto := new(sproto.PermissionDTO)
		err = json.Unmarshal([]byte(value), dto)
		if err == nil {
			m[key] = dto
		} else {
			log.Errorf("%s, %v", value, err)
		}
	}
	return m, err
}