## Rust
### Installation 
```cmd
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
rustc --version # rustc 1.86.0 (05f9846f8 2025-03-31)
```
### hello world

```rust
# save as main.rs
fn main(){
println!("Hello Pilathraj, welcome to Rust world");
}
```
### run
```rust
rustc main.rs
./main # Hello Pilathraj, welcome to Rust world
```
### Cargo - Dependency Management
```rust
cargo --version  # cargo 1.86.0 (adf9b6ad1 2025-02-28)
cargo new hello_cargo  # new project created hello_cargo -> src and  Cargo.toml
cargo build   # target folder has been created
cargo run     # Hello World
cargo check   # Compile only, there is no binary created.
```

