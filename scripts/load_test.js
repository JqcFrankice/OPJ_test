import http from 'k6/http';
import { check, sleep } from 'k6';
// 导入自定义的 TCP 扩展模块 (需构建后生效)
import tcp from 'k6/x/tcp'; 

export const options = {
  vus: 10,
  duration: '30s',
};

export default function () {
  // 1. HTTP 测试示例
  const res = http.get('https://test.k6.io');
  check(res, { 'is status 200': (r) => r.status === 200 });

  // 2. TCP 测试示例 (使用自定义插件)
  // 假设 TCP 插件提供了 connect, send, close 方法
  const client = tcp.connect('127.0.0.1:8080');
  client.send('login_request_data');
  const response = client.receive();
  check(response, { 'tcp response received': (r) => r !== '' });
  client.close();

  sleep(1);
}
