// solution 1
// sort ranges
// normalize ranges to not have overlaps
// sort nums
// for each range:
// // binary search left boundary
// // binary search right boundary
// // calculate count based on index difference
// // add count to result
// return result

// solution 2
// sort ranges
// sort nums
// initialize range index
// initialize num index
// while num < range: increment num index
// while num in range: increment num index, increment result
// while num > range: increment range index
// return result

use std::io;
use std::io::BufRead;

fn main() {
    let res = 0;

    let ranges = vec![];
    let nums = vec![];

    let lines = io::stdin().lock().lines();

    for line in lines {
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

        let a = match a.parse::<i32>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse left boundary of range: {}", e),
        };

        let b = match b.parse::<i32>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse right boundary of range: {}", e),
        };

        ranges.append((a, b));
    }

    for line in lines {
        let line = match line {
            Ok(v) => v,
            Err(e) => panic!("failed to read line: {}", e),
        };

        let num = match line.parse::<i32>() {
            Ok(v) => v,
            Err(e) => panic!("failed to parse num: {}", e),
        };

        nums.append(num);
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
