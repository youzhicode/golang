package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
	golang 处理错误的五种策略
	1、传播错误，即在函数内部调用子函数时发生了错误，导致该函数发生错误
		处理方案：一般直接将错误返回给调用方，或者将整理过的错误信息返回给调用方

	2、如果一个错误时偶然性的，或者不可预知的问题导致的，
		处理方案: 明智的选择是重新尝试失败的操作

	3、如果发生错误，程序无法继续运行，我们应该采用第三种策略
		处理方案：输出错误程序并结束程序，需要注意的是，这种策略只应在main中使用
				，如果是库函数而言，应该向上传播错误

	4、第四种策略， 有时候，我们只需要输出错误信息就足够了
	log.Printf("Ping failed: %v;networking disabled", err)
	或者标准错误流输出信息
	fmt.Fprintf(os.Stderr, "Ping failed: %v;networking disabled", err)

	log包中所有的函数都会为没有换行符的字符串添加上换行符

	5、第五种策略, 我们可以直接忽略掉错误
		我们应该养成一个好的习惯，当你决定忽略某一个错误的时候，应该清晰的写下你的意图

	Go中，错误处理有一套独特的编程风格，检查某一个子函数是否失败后，我们通常将处理失败的逻辑代码
	放在处理成功的代码之前。如果某一个错误会导致函数返回，那么成功时的逻辑代码不应该放在else语句块中，而应直接
	放在函数体中。Go中大部分函数的代码结构几乎都相同，首先一系列的初始化检查，防止错误的发生，之后是函数的实际逻辑
*/

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	// 文件结尾错误EOF
	// 如何判断是读取过程中出现错误，还是读取到最后结束的错误
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reader
		}
		if err != nil {
			os.Exit(1)
		}
	}

	// 第三种处理错误的策略
	if err := WaitForServer(""); err != nil {
		//fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		//os.Exit(1)
		// 使用log.Fatal可以更加简洁的实现与上文同样的结果,log中的所有函数，都默认会在错误信息之前输出时间信息
		// 长时间运行的服务通常采用默认的数据格式，而交互工具很少采用包含如此多信息的格式
		// 我们可以设置log的前缀信息屏蔽时间信息，一般而言，前缀信息会被设置成命令名
		log.SetPrefix("Wait:")
		log.SetFlags(0)
		log.Fatalf("Site is donw:%v\n", err)

	}

}
