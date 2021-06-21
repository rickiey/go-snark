## 编译
    1. 拉取源码，拉取依赖仓库：git submodule update --init --recursive
    2. 修改代码
    3. 进入extern/filecoin-ffi, 执行make all
    4. 在go-snark目录, 执行go build -o go-snark cmd/snark-server/server.go

## 修改
    1. extern/bellperson/src/gpu/locks.rs
        pub fn lock() -> GPULock {} 注释这一行 //f.lock_exclusive().unwrap();
    2. extern/filecoin-ffi/rust/Cargo.toml
        在最底下添加 
        [patch.crates-io]
        bellperson = {path = "../../bellperson"}
## 运行
    1. 分别添加配置文件config0.toml, config1.toml
    2. 执行go-snark --conf=config0.toml 使用0号GPU
    3. 执行go-snark --conf=config1.toml 使用1号GPU