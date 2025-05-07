## Rust
### Installation 
```cmd
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
rustc --version # rustc 1.86.0 (05f9846f8 2025-03-31)
```
### Hello World Program save as main.rs

```rust
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

### Variables
```cmd
let apples = 5;  // immutable variable by default
let mut name = "raj"  // mutable variable should add mut before the variable name.
```

### Guessing number
```rust
use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
    
    let secret_number = rand::thread_rng().gen_range(1..=100);

    loop {
        println!("Please input your guess.");
        let mut guess = String::new();
        
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");
        println!("You guessed: {}", guess.trim());
        
        let guess: u32 = guess.trim().parse().expect("Please enter a number");

        // println!("The random number is: {}", secret_number);
    
            println!("Please input your guess.");
            match guess.cmp(&secret_number){
                Ordering::Less => println!("Too small!"),
                Ordering::Greater => println!("Too big!"),
                Ordering::Equal => { println!("You win!"); break; }
            }
    }

}

```

