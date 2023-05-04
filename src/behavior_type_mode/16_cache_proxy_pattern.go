package behavior_type_mode

import (
	"fmt"
	"time"
)

type DataProvider interface {
	GetData(id string) string
}

type RealDataProvider struct{}

func (rdp *RealDataProvider) GetData(id string) string {
	// Simulate an expensive operation
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Real data for ID: %s", id)
}

type CachedDataProvider struct {
	provider DataProvider
	cache    map[string]string
}

func (cdp *CachedDataProvider) GetData(id string) string {
	data, ok := cdp.cache[id]
	if !ok {
		data = cdp.provider.GetData(id)
		cdp.cache[id] = data
	}
	return data
}

func NewCachedDataProvider(provider DataProvider) *CachedDataProvider {
	return &CachedDataProvider{
		provider: provider,
		cache:    make(map[string]string),
	}
}

func CallCacheProxyPattern() {
	realProvider := &RealDataProvider{}
	cachedProvider := NewCachedDataProvider(realProvider)
	startTime := time.Now()
	data := cachedProvider.GetData("1")
	fmt.Println("Data:", data)
	fmt.Println("Time spent:", time.Since(startTime))
	startTime = time.Now()
	data = cachedProvider.GetData("1")
	fmt.Println("Data:", data)
	fmt.Println("Time spent:", time.Since(startTime))
}
