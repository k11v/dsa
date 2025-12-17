use std::cmp::max;
use std::io;
use std::io::BufRead;

fn main() {
    let mut res = 0;

    let mut ranges = vec![];

    for line in io::stdin().lock().lines() {
        let line = match line {
            Ok(v) => v,
            Err(e) => panic!("failed to read line: {}", e),
        };

        if line == "" {
            break;
        }

        let (a, b) = match line.split_once('-') {
            Some(v) => v,
            None => panic!("failed to parse range: '-' not found"),
        };

        let a = match a.parse::<i64>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse left boundary of range: {}", e),
        };

        let b = match b.parse::<i64>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse right boundary of range: {}", e),
        };

        ranges.push((a, b + 1));
    }

    ranges.sort();

    let mut last_r = 0;

    for range in ranges {
        let l = max(range.0, last_r);
        let r = max(range.1, last_r);
        res += r - l;
        last_r = r;
    }

    println!("{}", res);
}
