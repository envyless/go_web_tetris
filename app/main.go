package main

import (
	"bufio"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var uiString string

func startGame() {
	// engine init
	engine := GameEngine{}

	// add modules, for game engine
	blockCtrl := New(10, 20, &engine)

	for {
		engine.Update()
		uiString = blockCtrl.Draw()
		//tick
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var loggerConfig = zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zap.DebugLevel)

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return
	}

	defer l.Close()

	// Using sync.Map to not deal with concurrency slice/map issues
	var connMap = &sync.Map{}

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Error("error accepting connection", zap.Error(err))
			return
		}

		id := uuid.New().String()
		connMap.Store(id, conn)

		go handleUserConnection(id, conn, connMap, logger)
	}
}

func handleUserConnection(id string, c net.Conn, connMap *sync.Map, logger *zap.Logger) {
	defer func() {
		c.Close()
		connMap.Delete(id)
	}()

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			logger.Error("error reading from client", zap.Error(err))
			return
		}

		connMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if _, err := conn.Write([]byte(userInput)); err != nil {
					logger.Error("error on writing to connection", zap.Error(err))
				}
			}

			return true
		})
	}
}

// var addr = flag.String("addr", ":8080", "http service address")

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL)
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	http.ServeFile(w, r, "home.html")
// }

// func writeTest(w http.ResponseWriter, r *http.Request) {
// 	var i int
// 	i = 0

// 	for {
// 		fmt.Fprint(w, "\033[H\033[2J")
// 		fmt.Fprint(w, "Hello World"+strconv.Itoa(i))
// 		i++
// 	}

// }

// func main() {
// 	flag.Parse()
// 	hub := newHub()
// 	go startGame()
// 	go hub.run()
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		writeTest(w, r)
// 	})
// 	//http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(hub, w, r)
// 	})
// 	err := http.ListenAndServe(*addr, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
