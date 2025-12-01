use std::io;
use std::io::Read;

fn main() {
    // Read input.
    let mut input = String::new();
    io::stdin()
        .read_to_string(&mut input)
        .expect("stdin should be readable");
    let input = input;

    // Parse instructions.
    let instructions = input.split("\n");
    for instruction in instructions {
        println!("{}", instruction)
    }
}
