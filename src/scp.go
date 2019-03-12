package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"bufio"
	"sync"
)

type Command interface {
	GetCommand() string
}

type Host struct {
	host string
	port string
	user string
}

type SSH struct {
	host Host
	cmd  string
}

type SCP struct {
	host Host
	localfile string
	remotefile string
}

func (s SSH) GetCommand() string {
	return fmt.Sprintf("ssh -p%s %s@%s '%s'", s.host.port, s.host.user, s.host.host, s.cmd)
}

func (s SCP) GetCommand() string {
	return fmt.Sprintf("scp -r -P%s %s %s@%s:%s", s.host.port, s.localfile, s.host.user, s.host.host, s.remotefile)
}

func ParseServers(filepath string) ([]Host, error) {
	var hosts []Host
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		info := strings.Fields(string(scanner.Text()))
		if len(info) != 3 {
			continue
		}
		hosts = append(hosts, Host{info[0], info[1], info[2]})
	}
	return hosts, nil
}

func DoAction(c string, h string, wg *sync.WaitGroup) {
	defer wg.Done()
	output, err := exec.Command("bash", "-c", c).Output()
	if err != nil {
		log.Fatal("Failed to excute command: " + c)
	}
	fmt.Printf("%s output:\n %s\n", h, output)
}

var (
	t  = flag.String("t", "", "操作类型: ssh|scp")
	f  = flag.String("f", "", "列表文件: 文件格式 ip ssh端口 用户名 (空格分隔, 一行一个主机)")
	c  = flag.String("c", "", "运行的命令: 必须与-t ssh选项一起使用")
	lf = flag.String("lf", "", "要拷贝的本地文件: 必须与-t scp选项一起使用")
	rf = flag.String("rf", "", "拷贝为的远程文件: 必须与-t scp选项一起使用")
)

func main() {
	flag.Parse()
	if (*t != "ssh" && *t != "scp") || (*t == "ssh" && *c == "") || (*t == "scp" && (*lf == "" || *rf == "")) {
		return
	}
	hosts, err := ParseServers(*f)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(hosts))
	for _, host := range hosts {
		switch *t {
		case "ssh":
			s := SSH{host, *c}
			go DoAction(s.GetCommand(), s.host.host, wg)
		case "scp":
			s := SCP{host, *lf, *rf}
			go DoAction(s.GetCommand(), s.host.host, wg)
		default:
			return
		}
	}
	wg.Wait()
}

/**
用法：先编辑好服务器队列，假设名称为server.list，格式如下：(ip port username)

192.168.11.1 22 root
192.168.11.2 22 root
192.168.11.3 22 root
当然，这里执行的主机上要有其他机器的SSH权限。

scp拷贝，假设拷贝/root/do.sh：

./sshworker -f server.list -t scp -lf /root/do.sh -rf /root/do.sh
ssh执行命令：

./sshworker -f server.list -t ssh -c "/bin/bash /root/do.sh"
备注下知识点：

flag包用于解析命令行输入；
bufio.NewScanner这个函数可以逐行读取文件内容；
exec.Command用bash -c "command"这种方式的话，可以在command里面用管道等复杂指令；
goroutine用于并发，waitgroup用于控制并发(不过我这里并没有控制)；
如果要执行长时间的命令，用at now <<< "command"扔后台里执行比较好，缺点就是不知道执行情况，不过可以输出到日志。
 */