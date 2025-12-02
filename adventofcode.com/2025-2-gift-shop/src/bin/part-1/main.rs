use std::io::{self, Read};

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    let mut result = 0;

    let ranges = input.split(",");

    for range in ranges {
        // Parse start and end.
        let (s, e) = range.split_once("-").unwrap();
        let mut s = s.trim().to_owned();
        let mut e = e.trim().to_owned();

        // Round start and end to an even number of digits.
        if s.len() % 2 != 0 {
            s = "1".to_owned() + &"0".repeat(s.len());
        }
        if e.len() % 2 != 0 {
            e = "9".repeat(e.len() - 1);
        }

        // Make start and end halfs.
        let sh = &s[0..s.len() / 2];
        let eh = &e[0..e.len() / 2];

        // Parse integers from start and end halfs.
        let mut shi = sh.parse::<i64>().unwrap();
        let mut ehi = eh.parse::<i64>().unwrap();

        // Adjust start and end halfs.
        if sh.repeat(2) < s {
            // If invalid ID is not within bounds, exclude it by moving the start.
            shi += 1;
        }
        if eh.repeat(2) <= e {
            // If invalid ID is within bounds, include it by moving the end.
            ehi += 1;
        }

        // Compute the sum of invalid IDs.
        while shi < ehi {
            let invalid_id = shi * 10i64.pow(shi.ilog10() + 1) + shi;
            result += invalid_id;
            shi += 1;
        }
    }

    println!("{}", result);
}
