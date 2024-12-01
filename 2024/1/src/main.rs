use anyhow::Ok;
use itertools::Itertools;
use std::{env, fs::read_to_string};

fn main() -> anyhow::Result<()> {
    let path = env::args().nth(1).ok_or(anyhow::anyhow!("no path"))?;
    let read = read_to_string(path)?;
    let input = read.lines();

    let mut r = vec![];
    let mut l = vec![];

    input.for_each(|x| {
        let (a, b) = x.split_whitespace().collect_tuple().unwrap();
        r.push(a.parse::<i32>().unwrap());
        l.push(b.parse::<i32>().unwrap());
    });

    fn part1(r: &mut Vec<i32>, l: &mut Vec<i32>) -> i32 {
        r.sort();
        l.sort();
        r.iter().zip(l.iter()).map(|(a, b)| (a - b).abs()).sum()
    }

    fn part2(r: &mut Vec<i32>, l: &mut Vec<i32>) -> i32 {
        l.iter()
            .map(|a| a * r.iter().filter(|&b| b == a).count() as i32)
            .sum()
    }

    println!("{}", part1(&mut r, &mut l));
    println!("{}", part2(&mut r, &mut l));
    Ok(())
}
