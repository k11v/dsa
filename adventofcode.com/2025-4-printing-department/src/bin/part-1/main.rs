use std::io;
use std::io::BufRead;

fn main() {
    let rows: Vec<Vec<char>> = io::stdin()
        .lock()
        .lines()
        .map(|x| x.unwrap().chars().collect())
        .collect();

    if rows.len() == 0 {
        return;
    }

    let n = rows.len();
    let m = rows[0].len();

    let mut count = 0;

    for i in 0..n {
        for j in 0..m {
            if rows[i][j] == '@' && count_neighbors(&rows, n, m, i, j) < 4 {
                count += 1;
            }
        }
    }

    println!("{}", count);
}

fn count_neighbors(rows: &Vec<Vec<char>>, n: usize, m: usize, i: usize, j: usize) -> i32 {
    let mut count = 0;

    let n = n as i32;
    let m = m as i32;
    let i = i as i32;
    let j = j as i32;

    for (ni, nj) in [
        (i - 1, j - 1),
        (i - 1, j),
        (i - 1, j + 1),
        (i, j - 1),
        (i, j + 1),
        (i + 1, j - 1),
        (i + 1, j),
        (i + 1, j + 1),
    ] {
        if ni >= 0 && ni < n && nj >= 0 && nj < m && rows[ni as usize][nj as usize] == '@' {
            count += 1;
        }
    }

    return count;
}
