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

fn main() {
    let res = 0;

    let lines = io::stdio().lock().lines();

    for line in lines {
        if line == "" {
            break;
        }
    }

    for line in lines {}

    println!("{}", res);
}
