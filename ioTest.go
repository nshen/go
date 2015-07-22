package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/**

====================================================================
io 为IO原语（I/O primitives）提供基本的接口
====================================================================

最重要的接口 Reader 和 Writer

读取实现接口的数据到byte slice里

type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}

os.File 同时实现了io.Reader和io.Writer
strings.Reader 实现了io.Reader
bufio.Reader/Writer 分别实现了io.Reader和io.Writer
bytes.Buffer 同时实现了io.Reader和io.Writer
bytes.Reader 实现了io.Reader
compress/gzip.Reader/Writer 分别实现了io.Reader和io.Writer
crypto/cipher.StreamReader/StreamWriter 分别实现了io.Reader和io.Writer
crypto/tls.Conn 同时实现了io.Reader和io.Writer
encoding/csv.Reader/Writer 分别实现了io.Reader和io.Writer
mime/multipart.Part 实现了io.Reader

--------------

ReaderAt WriterAt接口指定偏移读取

type ReaderAt interface {
        ReadAt(p []byte, off int64) (n int, err error)
}

type WriterAt interface {
        WriteAt(p []byte, off int64) (n int, err error)
}

---------------
//ReaderFrom和WriterTo 一次性从某个地方读或写到某个地方去

type ReaderFrom interface {
		//ReadFrom reads data from r until EOF or error.
        ReadFrom(r Reader) (n int64, err error)
}
type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
}

-----------

type Seeker interface {
    Seek(offset int64, whence int) (ret int64, err error)
}

Seek 设置下一次 Read 或 Write 的偏移量为 offset，它的解释取决于 whence：
0 表示相对于文件的起始处，
1 表示相对于当前的偏移，
2 表示相对于其结尾处。

os包里定义了相应常量
const (
    SEEK_SET int = 0 // seek relative to the origin of the file
    SEEK_CUR int = 1 // seek relative to the current offset
    SEEK_END int = 2 // seek relative to the end
)

reader := strings.NewReader("Go语言学习园地")
reader.Seek(-6, os.SEEK_END)
r, _, _ := reader.ReadRune()
fmt.Printf("%c\n", r)

------------

type Closer interface {
    Close() error
}
该接口比较简单，只有一个Close()方法，用于关闭数据流。
文件(os.File)、归档（压缩包）、数据库连接、Socket等需要手动关闭的资源都实现了Closer接口。


---------------
ByteReader和ByteWriter

type ByteReader interface {
    ReadByte() (c byte, err error)
}
type ByteWriter interface {
    WriteByte(c byte) error
}

bufio.Reader/Writer 分别实现了io.ByteReader和io.ByteWriter
bytes.Buffer 同时实现了io.ByteReader和io.ByteWriter
bytes.Reader 实现了io.ByteReader
strings.Reader 实现了io.ByteReader
-------------

====================================================================
io/ioutil 封装一些实用的I/O函数
====================================================================

func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm os.FileMode) error

-----------------
f, err := ioutil.TempFile("", "gofmt") //临时文件需要自己手动关闭
defer func() {
    f.Close()
    os.Remove(f.Name())
}()



====================================================================
bufio 实现带缓冲I/O
====================================================================

// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	// r:从buf中读走的字节（偏移）；w:buf中填充内容的偏移；
    // w - r 是buf中可被读的长度（缓存数据的大小），也是Buffered()方法的返回值
	err          error
	lastByte     int // 最后一次读到的字节（ReadByte/UnreadByte)
	lastRuneSize int // 最后一次读到的Rune的大小(ReadRune/UnreadRune)
}

bufio 包提供了两个实例化 bufio.Reader 对象的函数：NewReader 和 NewReaderSize。其中，NewReader 函数是调用 NewReaderSize 函数实现的：



====================================================================
fmt 实现格式化I/O，类似C语言中的printf和scanf
====================================================================

*/

func ioTest() {
	newDivider("ioTest.go")

	ioReadTest()
	ioWriteAtTest()

}

func ioWriteAtTest() {
	file, err := os.Create("./tmp/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	n, err := file.WriteAt([]byte("Go语言学习园地"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func ioReadTest() {
	someReader := strings.NewReader("from stringgg")

	data, err := ReadFrom(someReader, 12)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data, len(data))
	for _, str := range data {
		fmt.Print(string(str))
	}
	fmt.Print("\n")
}

//ReadFrom可以把实现io.Reader接口任何东东(标准输入、文件、字符串等)读到byte数组里
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	if n > 0 { //如果n>0,不管有没有err都要处理,在 n > 0 且数据被读完了的情况下，返回的error有可能是EOF也有可能是nil。
		return p[:n], nil
	}
	return p, err
}
