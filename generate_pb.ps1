# 自动生成 Protobuf 的 PowerShell 脚本
$pbDir = "proto/pb"
New-Item -ItemType Directory -Force -Path $pbDir

# 确保安装了插件
Write-Host "安装 Protobuf Go 插件..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 获取插件路径
$goBin = Join-Path $env:USERPROFILE "go\bin\protoc-gen-go.exe"

# 生成代码
Write-Host "正在生成..."
protoc --plugin=protoc-gen-go=$goBin --go_out=$pbDir --go_opt=paths=source_relative proto/msg.proto

Write-Host "Success!"
