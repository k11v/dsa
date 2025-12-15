use std::io;
use std::io::BufRead;

fn main() {
    let lines = io::stdin().lock().lines().map(|x| x.unwrap());

    let mut res = 0;

    for line in lines {
        let digits: Vec<char> = line.chars().collect();

        let mut a = digits[0];
        let mut l = 0;

        for i in 1..(digits.len() - 1) {
            if digits[i] > a {
                a = digits[i];
                l = i;
            }
        }

        let mut b = digits[l + 1];

        for i in l + 2..digits.len() {
            if digits[i] > b {
                b = digits[i];
            }
        }

        res += String::from_iter([a, b]).parse::<i32>().unwrap();
    }

    println!("{}", res);
}
