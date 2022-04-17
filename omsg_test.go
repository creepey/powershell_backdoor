// package main

// import (
// 	"bytes"
// 	"encoding/hex"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net"
// 	"testing"
// 	"time"
// )

// type testServer struct {
// 	s *Server
// }

// func (t *testServer) OnAccept(conn net.Conn) bool { return true }
// func (t *testServer) OnData(conn net.Conn, cmd, ext uint16, data []byte) error {
// 	if _, err := t.s.Send(conn, cmd, ext, data); err != nil {
// 		log.Fatalln(err)
// 	}
// 	return nil
// }
// func (t *testServer) OnRecvError(conn net.Conn, err error) { fmt.Println("youcuowu") }
// func (t *testServer) OnClientClose(conn net.Conn)          {}

// type testClient struct{}

// func (t *testClient) OnSuccess(conn net.Conn) {}
// func (t *testClient) OnData(cmd, ext uint16, data []byte) error {
// 	fmt.Println(hex.Dump(data[:40]))
// 	copy(recvBuff, data)
// 	return nil
// }
// func (t *testClient) OnRecvError(err error) {}
// func (t *testClient) OnClose()              {}

// var (
// 	sendBuff = make([]byte, 1024)
// 	recvBuff = make([]byte, 1024)
// )

// func Test_l(t *testing.T) {
// 	if _, err := rand.Read(sendBuff); err != nil {
// 		t.Fatal(err)
// 	}

// 	go func() {
// 		ts := &testServer{}
// 		s, err := Listen("tcp", ":1234")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		s.Run(ts)
// 	}()

// 	time.Sleep(time.Second)

// 	tc := &testClient{}
// 	c, err := Dial("tcp", ":1234", tc)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for i := 0; i < 3; i++ {
// 		if _, err := c.Send(1, uint16(i), sendBuff); err != nil {
// 			t.Fatal(err)
// 		}
// 	}

// 	time.Sleep(time.Second * 3)
// 	if !bytes.Equal(sendBuff, recvBuff) {
// 		t.Fail()
// 	}
// }
package main
