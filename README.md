# ProxyVerifier

ProxyVerifier 是一个用于验证 HTTP 代理功能和匿名性的 Go 工具。它通过多个 IP 信息服务检查代理，并比较结果以确定代理是否正常工作并保持匿名。

ProxyVerifier is a Go tool designed to verify the functionality and anonymity of HTTP proxies. It checks proxies against multiple IP information services and compares the results to determine if the proxy is working correctly and maintaining anonymity.

## 特性 (Features)

- 通过多个 IP 信息服务验证 HTTP 代理
- 通过比较本地 IP 和代理 IP 检查代理匿名性
- 支持并发代理检查
- 可自定义代理请求超时时间
- 以可读格式输出结果

- Verifies HTTP proxies against multiple IP information services
- Checks proxy anonymity by comparing local IP with proxy IP
- Supports concurrent proxy checking
- Customizable timeout for proxy requests
- Outputs results in a readable format

## 安装 (Installation)

要使用 ProxyVerifier，您需要在系统上安装 Go。然后，您可以克隆存储库并构建项目：

To use ProxyVerifier, you need to have Go installed on your system. Then, you can clone the repository and build the project:

```bash
git clone https://github.com/timwhitez/ProxyVerifier.git
cd ProxyVerifier
go build
```

## 使用方法 (Usage)

要使用 ProxyVerifier，请使用以下命令行参数运行编译后的二进制文件：

To use ProxyVerifier, run the compiled binary with the following command-line arguments:

```bash
./ProxyVerifier -f <proxy_file> -t <timeout> -c <concurrency>
```

参数 (Arguments):
- `-f`: 包含代理地址的文件路径（每行一个）
- `-t`: 代理请求超时时间（秒），默认为 10
- `-c`: 用于代理检查的并发 goroutine 数量，默认为 10

- `-f`: Path to the file containing proxy addresses (one per line)
- `-t`: Timeout for proxy requests in seconds (default: 10)
- `-c`: Number of concurrent goroutines for proxy checking (default: 10)

示例 (Example):
```bash
./ProxyVerifier -f proxies.txt -t 15 -c 20
```

## 代理文件格式 (Proxy File Format)

代理文件应该每行包含一个代理地址，格式如下：

The proxy file should contain one proxy address per line in the following format:

```
http://ip:port
```

示例 (Example):
```
http://192.168.1.1:8080
http://10.0.0.1:3128
```

## 注意事项 (Note)

本工具仅用于教育和测试目的。请确保您有权使用和测试您正在验证的代理。

This tool is intended for educational and testing purposes only. Ensure you have permission to use and test the proxies you're verifying.


## 贡献 (Contributing)

欢迎贡献！请随时提交 Pull Request。

Contributions are welcome! Please feel free to submit a Pull Request.
