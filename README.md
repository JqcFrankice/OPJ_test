# 游戏压力测试框架 (基于 k6 + Go)

这是一个使用 [k6](https://k6.io/) 和 Go 语言扩展机制 `xk6` 构建的高性能游戏压力测试框架原型，支持 HTTP 和自定义 TCP 协议。

## 项目结构

- `scripts/`: 存放 k6 测试脚本 (JavaScript)。
- `extensions/`: 存放自定义 Go 扩展代码。

## 环境搭建

1. **安装 k6**:
   请访问 [k6 官网](https://k6.io/docs/get-started/installation/) 安装 k6。

2. **安装 xk6 (用于构建带自定义插件的 k6)**:
   ```bash
   go install go.k6.io/xk6/cmd/xk6@latest
   ```

## 如何使用

### 1. 构建自定义 k6 (带 TCP 插件)

```bash
xk6 build --with github.com/your-username/k6-game-test/extensions/tcp=.
```
这会生成一个名为 `k6` 的可执行文件，其中包含了您的自定义 TCP 逻辑。

### 2. 运行测试脚本

```bash
./k6 run scripts/load_test.js
```
