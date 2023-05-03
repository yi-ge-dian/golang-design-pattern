package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Instances interface{}
type Zones interface{}

var providersMutex sync.Mutex
var providers = make(map[string]Factory_)

type Interface interface {
	Instances() (Instances, bool)
	Zones() (Zones, bool)
}

type Factory_ func(filePath string) (Interface, error)

func RegisterCloudProvider(name string, cloud Factory_) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	providers[name] = cloud
}

func GetCloudProvider(name string, filePath string) (Interface, error) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	f, found := providers[name]
	if !found {
		return nil, nil
	}
	return f(filePath)
}

func CallPluginPattern() {
	RegisterCloudProvider("gce", newGCECloud)
	gceCloud, _ := GetCloudProvider("gce", "/Users/wenlongdong/GolandProjects/golang-design-pattern/src/12_plugin_exmaple.json")
	fmt.Printf("%v", gceCloud)
}

// GCECloud is a simplified GCE Cloud provider
type GCECloud struct {
	ProjectID string `json:"ProjectID"`
	Zone      string `json:"Zone"`
}

func newGCECloud(filePath string) (Interface, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	// 读取 JSON 数据并解码到结构体中
	var gceCloud GCECloud
	for {
		err := decoder.Decode(&gceCloud)
		if err == io.EOF {
			break // 已经读完文件
		} else if err != nil {
			log.Fatal(err)
		}
		// 打印解码后的字段值
		fmt.Printf("ProjectID: %s\n", gceCloud.ProjectID)
		fmt.Printf("Zone: %s\n", gceCloud.Zone)
	}

	return &GCECloud{
		ProjectID: gceCloud.ProjectID,
		Zone:      gceCloud.Zone,
	}, nil
}

func (gce *GCECloud) Instances() (Instances, bool) {
	return gce, true
}

func (gce *GCECloud) Zones() (Zones, bool) {
	return gce, true
}
