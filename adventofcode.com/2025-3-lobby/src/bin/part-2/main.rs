use std::io;
use std::io::BufRead;

const DIGITS: usize = 12;

fn main() {
    let lines = io::stdin().lock().lines().map(|x| x.unwrap());

    let mut res = 0;

    for line in lines {
        let digits: Vec<char> = line.chars().collect();

        let mut num_digits = ['0'; DIGITS];
        let mut n = 0;

        for i in 0..DIGITS {
            let d = &mut num_digits[i];
            *d = digits[n];
            for j in n + 1..digits.len() - DIGITS + 1 + i {
                if digits[j] > *d {
                    *d = digits[j];
                    n = j;
                }
            }
            n += 1;
        }

        res += String::from_iter(num_digits).parse::<i64>().unwrap();
    }

    println!("{}", res);
}
