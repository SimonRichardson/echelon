package env

import "testing"

type RedisClusterStringIO struct {
	In     string
	Out    [][]string
	HasErr bool
}

func TestRedisClustersFromString(t *testing.T) {
	ins := []RedisClusterStringIO{
		RedisClusterStringIO{
			In:     "",
			Out:    [][]string{},
			HasErr: false,
		},
		RedisClusterStringIO{
			In:     "redis1:6379",
			Out:    [][]string{[]string{"redis1:6379"}},
			HasErr: false,
		},
		RedisClusterStringIO{
			In:     "redis1:6379,redis2:6379",
			Out:    [][]string{[]string{"redis1:6379", "redis2:6379"}},
			HasErr: false,
		},
		RedisClusterStringIO{
			In:     "redis1:6379,redis2:6379;redis3:6379,redis4:6379",
			Out:    [][]string{[]string{"redis1:6379", "redis2:6379"}, []string{"redis3:6379", "redis4:6379"}},
			HasErr: false,
		},
		RedisClusterStringIO{
			In:     "redis1:6379,redis2:6379;",
			Out:    [][]string{[]string{"redis1:6379", "redis2:6379"}},
			HasErr: false,
		},
		RedisClusterStringIO{
			In:     "redis1,redis2;",
			Out:    [][]string{},
			HasErr: true,
		},
		RedisClusterStringIO{
			In:     "redis1:,redis2:;",
			Out:    [][]string{},
			HasErr: true,
		},
		RedisClusterStringIO{
			In:     "redis1:not,redis2:port;",
			Out:    [][]string{},
			HasErr: true,
		},
		RedisClusterStringIO{
			In:     ";;;",
			Out:    [][]string{},
			HasErr: true,
		},
	}

	for _, v := range ins {
		out, err := RedisClustersFromString(v.In)

		if v.HasErr && err == nil {
			t.Errorf("%v(%v) should return error", "RedisClustersFromString", v.In)
			t.FailNow()
		} else {
			continue
		}

		if !stringSliceSliceEqual(v.Out, out) {
			t.Errorf("%v(%v) != %v (%v)", "RedisClustersFromString", v.In, v.Out, out)
			t.FailNow()
		}
	}
}

func stringSliceSliceEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !stringSliceEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
