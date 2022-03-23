package main

import (
	"fmt"
	"github.com/cube-group/golib/cache"
	"github.com/cube-group/golib/uuid"
	"math/rand"
	"time"
)

var ramer *cache.Ram
var content string = `sakdjfalsjfdlajsfklajflj埃里克森江东父老卡金飞达将防盗链卡见附件发送的减法精神分裂撒加打发时间打法就发啦发啦街坊邻居安饭量但是就是
按数量就大法试卷得分卡所肩负的垃圾分类卡束带结发加分的卡夫拉设计费
阿斯利康打飞机阿里发动机辣椒粉垃圾费拉咖啡机拉发动机拉升发动机暗室逢灯垃圾啊是劳动法发了疯咖啡机啊
阿斯利康的房价阿斯顿发婕拉开发机奥拉夫静安里看风景阿卡发动机阿凡达卡拉街坊邻居阿法拉克发顺丰
sakdjfalsjfdlajsfklajflj埃里克森江东父老卡金飞达将防盗链卡见附件发送的减法精神分裂撒加打发时间打法就发啦发啦街坊邻居安饭量但是就是
按数量就大法试卷得分卡所肩负的垃圾分类卡束带结发加分的卡夫拉设计费
阿斯利康打飞机阿里发动机辣椒粉垃圾费拉咖啡机拉发动机拉升发动机暗室逢灯垃圾啊是劳动法发了疯咖啡机啊
阿斯利康的房价阿斯顿发婕拉开发机奥拉夫静安里看风景阿卡发动机阿凡达卡拉街坊邻居阿法拉克发顺丰
sakdjfalsjfdlajsfklajflj埃里克森江东父老卡金飞达将防盗链卡见附件发送的减法精神分裂撒加打发时间打法就发啦发啦街坊邻居安饭量但是就是
按数量就大法试卷得分卡所肩负的垃圾分类卡束带结发加分的卡夫拉设计费
阿斯利康打飞机阿里发动机辣椒粉垃圾费拉咖啡机拉发动机拉升发动机暗室逢灯垃圾啊是劳动法发了疯咖啡机啊
阿斯利康的房价阿斯顿发婕拉开发机奥拉夫静安里看风景阿卡发动机阿凡达卡拉街坊邻居阿法拉克发顺丰
sakdjfalsjfdlajsfklajflj埃里克森江东父老卡金飞达将防盗链卡见附件发送的减法精神分裂撒加打发时间打法就发啦发啦街坊邻居安饭量但是就是
按数量就大法试卷得分卡所肩负的垃圾分类卡束带结发加分的卡夫拉设计费
阿斯利康打飞机阿里发动机辣椒粉垃圾费拉咖啡机拉发动机拉升发动机暗室逢灯垃圾啊是劳动法发了疯咖啡机啊
阿斯利康的房价阿斯顿发婕拉开发机奥拉夫静安里看风景阿卡发动机阿凡达卡拉街坊邻居阿法拉克发顺丰`

func main() {
	ramer = cache.NewRam()
	var input string
	fmt.Scan(&input)
	startTime := time.Now().UnixNano()

	for i := 0; i < 10; i++ {
		go do()
	}

	fmt.Println("useTime:", (time.Now().UnixNano()-startTime)/1e6)
	for {
		fmt.Scan(&input)
		fmt.Println(ramer.DebugInfo())
	}
}

func do() {
	for {
		for i := 0; i < 1000; i++ {
			dt := 1 + rand.Int63n(2)
			ramer.Set(fmt.Sprintf("%s%d", uuid.GetUUID(), i), "1", time.Duration(dt)*time.Second)
		}
		time.Sleep(time.Second * 5)
	}

}
