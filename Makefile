.PHONY: build clean

# 默认编译目标
build:
	go build -ldflags="-s -w" -trimpath -o py4fzf_release main.go

# 开发版本（带调试信息）
dev:
	go build -o py4fzf1 main.go

# 清理编译产物
clean:
	rm -f py4fzf1
	rm -f py4fzf_release

# # 安装到系统
# install: build
# 	cp py4fzf1 /usr/local/bin/

# # 卸载
# uninstall:
# 	rm -f /usr/local/bin/py4fzf1 