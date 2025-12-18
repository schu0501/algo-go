package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// luaScript 保证 原子性：检查库存 > 0 则扣减，否则返回 0
const luaScript = `
if redis.call("get", KEYS[1]) > "0" then
    return redis.call("decr", KEYS[1])
else
    return -1
end
`

// 1. 全局配置
var (
	rdb       *redis.Client
	orderChan chan Order     // 【关键】异步写入订单的 Channel
	wg        sync.WaitGroup // 限制同时写数据库的协程数
)

const (
	MaxWorkers = 5         // 消费者并发数
	Stock      = 100       // 库存
	BufferSize = Stock * 2 // Channel 缓冲区设置为库存倍数
)

type Order struct {
	UserID    string
	ProductID string
	Time      time.Time
}

func init() {
	// 初始化 Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "G62m50oigInC30sf",
	})
	// 初始化 Channel (缓冲区大小设置为库存的倍数，防止阻塞)
	orderChan = make(chan Order, BufferSize)

	// 初始化库存 (假设只有 100 个)
	rdb.Set(context.Background(), "product:1001", Stock, 0)
}

// HandleSeckill 2. HTTP 接口处理函数 (高并发入口)
// 这里的 QPS 可能会非常高，GMP 模型轻松扛住
func HandleSeckill(userID string, productID string) string {
	ctx := context.Background()

	// --- A. Redis 预减库存 (挡住 99.9% 的流量) ---
	//执行 Lua 脚本
	ret, err := rdb.Eval(ctx, luaScript, []string{"product:" + productID}).Result()
	if err != nil {
		return "系统繁忙"
	}

	// 此时 ret 是扣减后的库存
	stock := ret.(int64)
	if stock < 0 {
		return "秒杀已结束" // 库存不足，直接返回
	}

	// --- B. 只有抢到库存的人，才有资格进入 Channel (削峰) ---
	//这是一个非阻塞的操作（除非 Channel 满了）
	select {
	case orderChan <- Order{UserID: userID, ProductID: productID, Time: time.Now()}:
		return "排队中，请稍后查询结果..."
	// 快速响应用户default:
	default:
		return "系统繁忙"
		// 极端情况：Redis 扣成功了，但 Channel 满了（队列积压严重）
		//这里的策略需要根据业务定，可以报错回滚 Redis，也可以丢弃return "系统拥堵，请重试"
	}
}

// StartConsumers 3. 后台消费者 (Worker Pool)
// 它们慢慢地从 Channel 拿数据，往 MySQL 里写
func StartConsumers() {
	for i := 0; i < MaxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for order := range orderChan {
				// 循环读取，直到 Channel 被 close 且数据被读空
				createOrderInMySQL(order)
				fmt.Printf("Worker %d 处理了用户 %s 的订单\n", workerID, order.UserID)
			}
		}(i)
	}
}

func createOrderInMySQL(order Order) {
	// 模拟耗时 SQL 操作
	time.Sleep(500 * time.Millisecond)
	// INSERT INTO orders ...、
	fmt.Println("写入数据库", order)
}

func main() {
	// 启动消费者池
	StartConsumers()

	var userWg sync.WaitGroup

	fmt.Println("--- 秒杀开始 ---")
	start := time.Now()
	// 模拟高并发请求 (1000 人抢 100 个)
	for i := 0; i < 1000; i++ {
		userWg.Add(1)
		go func(userId int) {
			defer userWg.Done()
			userID := fmt.Sprintf("User_%d", userId)
			resp := HandleSeckill(userID, "1001")
			if resp != "秒杀已结束" {

			}
		}(i)
	}

	userWg.Wait()
	fmt.Printf("--- 所有用户请求已处理完毕 (Redis层)，耗时: %v ---\n", time.Since(start))
	close(orderChan)

	fmt.Println("--- 等待后台 Worker 落库... ---")
	wg.Wait()
	fmt.Println("--- 所有订单落库完成，程序安全退出 ---")
}
