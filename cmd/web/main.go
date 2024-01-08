package main

import (
	"fmt"
	// "runtime"
	"sync"

	// "sync/atomic"
	"time"
	// "golang-clean-architecture/internal/config"
)

func main() {
	/*
		go fmt.Println("With Goroutine")
		fmt.Println("Without Goroutine")
	*/

	// ================================

	/*
		channel := make(chan string)
		defer close(channel)

		go fromIn(channel)
		go fromOut(channel)

		time.Sleep(5 * time.Second)
	*/

	// ================================

	/*
		channel := make(chan string, 3)
		defer close(channel)

		go func() {
			channel <- "Eko"
			channel <- "Kurniawan"
			channel <- "Khannedy"
		}()

		go func() {
			fmt.Println(<-channel)
			fmt.Println(<-channel)
			fmt.Println(<-channel)
		}()

		time.Sleep(2 * time.Second)
		fmt.Println("Selesai")
	*/

	// ================================

	/*
		channel := make(chan string)

		go func() {
			for i := 0; i < 10; i++ {
				channel <- "Perulangan ke " + strconv.Itoa(i)
			}
			close(channel)
		}()

		for data := range channel {
			fmt.Println("Menerima data", data)
		}

		fmt.Println("Selesai")
	*/

	// ================================

	/*
		channel1 := make(chan string)
		channel2 := make(chan string)

		defer close(channel1)
		defer close(channel2)

		go func() {
			time.Sleep(2 * time.Second)
			channel1 <- "Eko Kurniawan Khannedy"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			channel2 <- "Eko Kurniawan Khannedy"
		}()

		counter := 0
		for {
			select {
			case data := <-channel1:
				fmt.Println("Data dari Channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari Channel 2", data)
				counter++
			default:
				fmt.Println("Menungg Data...")
			}
			if counter == 2 {
				break
			}
		}
	*/

	// ================================

	// Prevent race condition with Mutex
	/*
		x := 0
		var mutex sync.Mutex
		for i := 1; i <= 1000; i++ {
			go func() {
				for j := 1; j <= 100; j++ {
					mutex.Lock()
					x = x + 1
					mutex.Unlock()
				}
			}()
		}
		time.Sleep(5 * time.Second)
		fmt.Println("Counter = ", x)
	*/

	// ================================

	/*
		account := BankAccount{}

		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					account.AddBalance(1)
					fmt.Println(account.Balance)
				}
			}()
		}
		time.Sleep(5 * time.Second)
	*/

	// ================================

	// Deadlock
	/*
		user1 := UserBalance{
			Name: "Eko",
			Balance: 1000000,
		}

		user2 := UserBalance{
			Name: "Budi",
			Balance: 1000000,
		}

		go Transfer(&user1, &user2, 100000)
		go Transfer(&user2, &user1, 200000)

		time.Sleep(10 * time.Second)

		fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
		fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
	*/

	// ================================

	/*
		group := &sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			go RunAsynchronus(group)
		}

		group.Wait()
		fmt.Println("Selesai")
	*/

	// ================================

	/*
		once := sync.Once{}
		group := sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			group.Add(1)
			go func() {
				once.Do(OnlyOnce)
				group.Done()
			}()
		}

		group.Wait()
		fmt.Println("Counter", counter)
	*/

	// ================================

	/*
		pool := sync.Pool{
			New: func() any {
				return "New"
			},
		}

		pool.Put("Eko")
		pool.Put("Kurniawan")
		pool.Put("Khannedy")

		for i := 0; i < 10; i++ {
			go func() {
				data := pool.Get()
				fmt.Println(data)
				time.Sleep(1 * time.Second)
				pool.Put(data)
			}()
		}
		time.Sleep(11 * time.Second)
		fmt.Println("Selesai")
	*/

	// ================================

	/*
		data := &sync.Map{}
		group := &sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			go AddToMap(data, i, group)
		}

		group.Wait()

		data.Range(func(key, value any) bool {
			fmt.Println(key, ":", value)
			return true
		})
	*/

	// ================================

	/*
		for i := 0; i < 10; i++ {
			go WaitCondition(i)
		}

		go func()  {
			for i := 0; i < 10; i++ {
				time.Sleep(1 * time.Second)
				cond.Signal()
			}
		}()

		// go func() {
		// 	time.Sleep(1 * time.Second)
		// 	cond.Broadcast()
		// }()

		group.Wait()
	*/

	// ================================

	/*
		var x int64 = 0
		group := sync.WaitGroup{}

		for i := 1; i <= 1000; i++ {
			group.Add(1)
			go func()  {
				for j := 1; j <= 100; j++ {
					atomic.AddInt64(&x, 1)
				}
				group.Done()
			}()
		}
		group.Wait()
		fmt.Println("Counter = ", x)
	*/

	// ================================

	/*
		timer := time.NewTimer(5 * time.Second)
		fmt.Println(time.Now())

		time := <- timer.C
		fmt.Println(time)
	*/

	// ================================

	/*
		timer := time.After(5 * time.Second)
		fmt.Println(time.Now())

		time := <- timer
		fmt.Println(time)
	*/

	// ================================

	/*
		group := sync.WaitGroup{}
		group.Add(1)

		time.AfterFunc(5 * time.Second, func() {
			fmt.Println(time.Now())
			group.Done()
		})
		fmt.Println(time.Now())
		group.Wait()
	*/

	// ================================

	/*
		ticker := time.NewTicker(1 * time.Second)

		go func() {
			time.Sleep(5 * time.Second)
			ticker.Stop()
		}()

		for time := range ticker.C {
			fmt.Println(time)
		}
	*/

	// ================================

	/*
		channel := time.Tick(1 * time.Second)

		for time := range channel {
			fmt.Println(time)
		}
	*/

	// ================================

	/*
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
	*/
}

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

var counter = 0

func OnlyOnce() {
	counter++
}

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.Lock()
	balance := account.Balance
	account.RWMutex.Unlock()
	return balance
}

// Send channel
func FromIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

// Receive channel
func FromOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// ===========================================================
// ===========================================================
// ===========================================================

/* func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)
	producer := config.NewKafkaProducer(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Producer: producer,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} */
