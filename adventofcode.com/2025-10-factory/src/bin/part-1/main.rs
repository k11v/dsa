use std::error;
use std::io::{self, BufRead};

fn main() -> Result<(), Box<dyn error::Error>> {
    let lines = io::stdin().lock().lines();

    for line in lines {
        let line = line?;
        let words: Vec<_> = line.split(" ").collect();
        if words.len() < 3 {
            return Err("invalid machine".into());
        }

        let want_lights = parse_lights(words[0])?;
        let got_lights = vec![0; want_lights.len()];
        let buttons = &words[1..words.len() - 1]
            .iter()
            .map(|x| parse_button(x))
            .collect::<Result<Vec<Vec<i32>>, _>>()?;
        let joltages = parse_joltages(words[words.len() - 1])?;

        println!("{got_lights:?}");
        println!("{want_lights:?}");
        println!("{buttons:?}");
        println!("{joltages:?}");
    }

    Ok(())
}

fn parse_lights(s: &str) -> Result<Vec<i32>, Box<dyn error::Error>> {
    let l = s.get(..1).unwrap_or_default();
    let m = s.get(1..(s.len() - 1)).unwrap_or_default();
    let r = s.get(s.len() - 1..).unwrap_or_default();

    if l != "[" {
        return Err("want [".into());
    }

    let m = m
        .bytes()
        .map(|x| match x {
            b'.' => Ok(0),
            b'#' => Ok(1),
            _ => Err("want . or #".into()),
        })
        .collect::<Result<Vec<i32>, Box<dyn error::Error>>>()?;

    if r != "]" {
        return Err("want ]".into());
    }

    Ok(m)
}

fn parse_button(s: &str) -> Result<Vec<i32>, Box<dyn error::Error>> {
    let l = s.get(..1).unwrap_or_default();
    let m = s.get(1..(s.len() - 1)).unwrap_or_default();
    let r = s.get(s.len() - 1..).unwrap_or_default();

    if l != "(" {
        return Err("want (".into());
    }

    let m = m
        .split(",")
        .map(|x| x.parse::<i32>())
        .collect::<Result<Vec<_>, _>>()?;

    if r != ")" {
        return Err("want )".into());
    }

    Ok(m)
}

fn parse_joltages(s: &str) -> Result<Vec<i32>, Box<dyn error::Error>> {
    let l = s.get(..1).unwrap_or_default();
    let m = s.get(1..(s.len() - 1)).unwrap_or_default();
    let r = s.get(s.len() - 1..).unwrap_or_default();

    if l != "{" {
        return Err("want {".into());
    }

    let m = m
        .split(",")
        .map(|x| x.parse::<i32>())
        .collect::<Result<Vec<_>, _>>()?;

    if r != "}" {
        return Err("want }".into());
    }

    Ok(m)
}
