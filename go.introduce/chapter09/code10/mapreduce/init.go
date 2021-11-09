package mapreduce

//包内全局变量
var (
	size      = 10
	chtxt1    = make(chan string, size)
	chtxt2    = make(chan string, size)
	chtxt3    = make(chan string, size)
	chtxt4    = make(chan string, size)
	chmap1    = make(chan map[string]int, size)
	chmap2    = make(chan map[string]int, size)
	chmap3    = make(chan map[string]int, size)
	chmap4    = make(chan map[string]int, size)
	chreduce1 = make(chan map[string]int, size)
	chreduce2 = make(chan map[string]int, size)
	chtemp1   = make(chan map[string]int, size)
	chtemp2   = make(chan map[string]int, size)
)
