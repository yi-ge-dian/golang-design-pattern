package main

import (
	"github.com/yi-ge-dian/golang-design-pattern/src/behavior_type_mode"
)

func main() {
	//creation_type_mode.CallSinglePattern()              // 单例模式
	//creation_type_mode.CallSampleFactoryPattern()       // 简单工厂模式
	//creation_type_mode.CallAbstractFactoryPattern()     // 抽象工厂模式
	//creation_type_mode.CallBuilderPattern()             // 建造者模式
	//creation_type_mode.CallConfigurePatternLowLevel()   // 建造者模式初级
	//creation_type_mode.CallConfigurePatternHighLevel()  // 建造者模式高级
	//creation_type_mode.CallPrototypePattern()           // 原型模式
	//structure_type_mode.CallAdaptorPattern()            // 适配器模式
	//structure_type_mode.CallBridgePattern()             // 桥接模式
	//structure_type_mode.CallObjectTreePattern()         // 对象树模式
	//structure_type_mode.CallDecoratorPattern()          // 装饰器模式
	//structure_type_mode.CallPipelinePattern()           // 管道模式
	//structure_type_mode.CallPluginPattern()             // 插件模式
	//behavior_type_mode.CallChainedCallPattern()         // 链式调用模式
	//behavior_type_mode.CallResponsibilityChainPattern() //责任链模式
	//behavior_type_mode.CallObserverPattern()            //观察者模式
	//behavior_type_mode.CallCacheProxyPattern()          //缓存代理模式
	behavior_type_mode.CallPolicyPattern() // 策略模式
}
