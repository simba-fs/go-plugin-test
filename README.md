# go-plugin-test
golang 動態載入外掛  
使用 1.8 新增的 [plugin](https://pkg.go.dev/plugin) 模組，動態載入 .so (share object) 檔  
這次實作了一個簡單的命令解譯器，透過外掛，可以不用重新編譯主程式就動態載入命令  

# 命令
這次的命令是經由 `go build -buildmode=plugin` 編譯過的 go 檔，必須暴露一個 `Exec` 函式，型態是 `func ([]string) error`。編譯過的 `.so` 檔必須放在 `./dist/` 目錄下，檔名就是命令，例如執行 `ls -a .` 的時候，解譯器就會去 `./dist/` 找 `ls.so` 這個命令，並且把 `["ls", "-a", "."]` 傳進函式 `Exec` 當作參數  
實際例子可以參考 `plguins/ls.so` 這個檔案

# 編譯外掛
執行 `./build.sh`，目前 go 只支援 Linux, FreeBSD, and macOS，其他的平台不能編譯

# 測試
執行 `go run .`，會開啟一個陽春的命令界面，在這裡打指令，程式會去 `dist` 目錄找同名的外掛並執行搭他的 `Exec` 函式。用 `exec` 或 `CTRL-c` 可以離開
