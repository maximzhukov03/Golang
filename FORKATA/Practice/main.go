package main

import (
	"fmt"
	"net"
)

type Users struct{
	pool []*User
}

type User struct{
	addr string
	new bool 
}

type Message struct{
	addr string
	payload []byte
}

type Server struct{
	listenAddr string
	ln net.Listener
	quitChan chan struct{}
	msgChan chan Message
	users Users
}

func NewServer(listenAddr string) *Server{
	return  &Server{
		listenAddr: listenAddr,
		quitChan: make(chan struct{}), 
		msgChan: make(chan Message, 10),
	}
}

func (s *Server) Start() error{
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil{
		return err
	}
	defer ln.Close()
	s.ln = ln
	fmt.Println("SERVER STARTED")

	go s.acceptLoop()

	<-s.quitChan

	return nil
}

func (s *Server) acceptLoop() {
	for{
		conn, err := s.ln.Accept()
		if err != nil{
			fmt.Println("ERROR ACCEPT", err) 
			continue
		}

		fmt.Println("NEW CONNECTION:", conn.RemoteAddr())
		
		s.users.pool = append(s.users.pool, &User{addr: conn.LocalAddr().String(), new: true})

		go s.readLoop(conn)

		


	}

	
}

func (s *Server) readLoop(conn net.Conn){
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("ERROR READ CONN", err)
			continue
		}

		s.msgChan <- Message{
			addr: conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		for _, user := range s.users.pool{
			if user.new{
				conn.Write([]byte(fmt.Sprintf("Welcome! [%s]\n", conn.LocalAddr().String())))
				user.new = false
			}
		}


	} 

}


func main(){
	server := NewServer("0.0.0.0:3000") 

	go func(){
		defer close(server.msgChan)
		for msg := range server.msgChan{
			fmt.Printf("Сообщение [от %s]: %s\n", msg.addr, string(msg.payload))
		}
	}()

	server.Start()


}