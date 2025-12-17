use std::io;
use std::io::BufRead;

fn main() {
    let mut res = 0;

    let mut ranges = vec![];
    let mut nums = vec![];

    let mut lines = io::stdin().lock().lines();

    for line in lines.by_ref() {
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

        ranges.push((a, b));
    }

    for line in lines {
        let line = match line {
            Ok(v) => v,
            Err(e) => panic!("failed to read line: {}", e),
        };

        let num = match line.parse::<i64>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse num: {}", e),
        };

        nums.push(num);
    }

    ranges.sort();
    nums.sort();

    let mut range_i = 0;
    let mut num_i = 0;

    while num_i < nums.len() && range_i < ranges.len() {
        while num_i < nums.len() && range_i < ranges.len() && nums[num_i] < ranges[range_i].0 {
            num_i += 1;
        }

        while num_i < nums.len() && range_i < ranges.len() && nums[num_i] <= ranges[range_i].1 {
            num_i += 1;
            res += 1;
        }

        while num_i < nums.len() && range_i < ranges.len() && nums[num_i] > ranges[range_i].1 {
            range_i += 1;
        }
    }

    println!("{}", res);
}
