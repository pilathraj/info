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
### Run
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
### Variables 
```rust
    println!("Hello, world!");
    const PI :f64 = 3.14; // Constants
    println!("PI is {}", PI); // 3.14
    let x = 5;
    //x = 24; // cannot assign twice to immutable variable
    println!("Value of x is {x}");
    let x = x+1; // Shadowing with let
    println!("Value of x is {x}"); // 6
    
    {
        let x = x*2; // within block scope
        println!("Value of x is {x}"); // 12
        
    }
    println!("Value of x is {x}"); // 6
    let spaces = "   ";
    println!("Value of spaces is {spaces}"); // 
    let spaces = spaces.len(); 
    println!("Value of spaces is {spaces}"); // 3
    
    let mut s = "  ";
    //s =s.len(); // expected `&str`, found `usize`
    println!("Value of s is {s}"); // 2
```
### Data Types
#### Scalar Types:
   - **Integer Types**
   
   |Length|Signed|Unsigned|
   |-----|-------|---------|
   |8-bit |	i8	|u8|
   |16-bit	| i16	|u16|
   |32-bit	|i32	|u32|
   |64-bit|	i64	|u64|
   |128-bit|	i128|u128|
   |arch	|isize	|usize|


   - **Floating-Point Types**
   
   |Length|Signed| |
   |-----|-------|-|
   |32-bit	|f32	||
   |64-bit|	f64	|default|

   - **Boolean Type**
```rs
let t = true;
let f :bool = false;
```
   - **Character Type**
```rs
    let c = 'z';
    let z: char = 'ℤ'; // with explicit type annotation
    let heart_eyed_cat = '😻';
```
#### Compound Types
   
3. ****
4. 


