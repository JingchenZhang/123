package main

import (
    "log"
    //"/root/nutsdb"
    "github.com/nutsdb/nutsdb"
)

func main() {
    db, err := nutsdb.Open(
        nutsdb.DefaultOptions,
        nutsdb.WithDir("/tmp/nutsdb"), // 数据库会自动创建这个目录文件
    )
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

}
