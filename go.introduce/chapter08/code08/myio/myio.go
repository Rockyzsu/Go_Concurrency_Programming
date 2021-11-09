package myio

import "os"

type MyIO struct {
	data  string
	bytes []byte
	file  *os.File
}

//io.Reader接口
func (myio *MyIO) Read(p []byte) (n int, err error) {

	if len(p) == 0 {
		return 0, nil
	}
	myio.data = string(p)
	return myio.file.Read(p)
}

//io.Writer接口
func (myio *MyIO) Write(p []byte) (n int, err error) {

	if len(p) == 0 {
		return 0, nil
	}
	myio.bytes = p
	return myio.file.Write(p)
}
func (myio *MyIO) Close() error {
	if myio == nil {
		return nil
	}
	return myio.file.Close()
}

func Create(name string) (*MyIO, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	io1 := &MyIO{file: f}
	return io1, nil
}

func Open(name string) (*MyIO, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	io1 := &MyIO{file: f}
	return io1, nil
}
