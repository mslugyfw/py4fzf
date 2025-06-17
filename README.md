# py4fzf1
use pinyin in fzf 在fzf中使用拼音拼音首字母查询

# 自用备份

# 文件：
- main.go : py4fzf代码。
- py4fzf(linux amd64 exe) ：将stdin输入中的汉字转换为拼音首字母和全拼。（输入: 我爱中国； 输出: 我爱中国|wazg|woaizhongguo） "py4fzf -o"将 “我爱中国|wazg|woaizhongguo”变成“我爱中国”。
- fff : 替代fzf的脚本，参数发送给脚本中的fzf。
- fff.yazi : yazi插件，放置到~/.config/yazi/plugins/, 修改keymap.toml中{ on = "z",         run = "plugin fff",                  desc = "Jump to a file/directory via fff" }，yazi中按“z”调用。


