package cache

import "github.com/karlseguin/ccache/v2"

type ADRCache struct {
	cc *ccache.Cache
}

func NewADRCache() *ADRCache {
	return &ADRCache{
		cc: nil,
	}
}

func (m *ADRCache) Init() {
	m.cc = ccache.New(ccache.Configure())
}

func (m *ADRCache) Set(key string, val interface{}) {
}

func (m *ADRCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (m *ADRCache) Delete(key string) {
}

func (m *ADRCache) Clean() {

}

var DomainCache *ADRCache
var DomainLogicCache *ADRCache

func init() {
	// TODO:
	DomainCache = NewADRCache()
	DomainLogicCache = NewADRCache()
}
