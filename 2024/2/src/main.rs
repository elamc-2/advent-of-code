use anyhow::Ok;
use itertools::Itertools;
use std::{env, fs::read_to_string, slice::Iter};

fn is_safe(it: Iter<i32>) -> bool {
    let (mut inc, mut dec) = (true, true);
    it.into_iter().tuple_windows().all(|(a, b)| {
        inc = inc && a <= b;
        dec = dec && a >= b;
        (1..=3).contains(&(a - b).abs()) && (inc || dec)
    })
}
pub fn with_dampener<T: Clone>(vec: &Vec<i32>) -> bool {
    for (i, _) in vec.iter().enumerate() {
        let mut v = vec.clone();
        v.remove(i);
        if is_safe(v.iter()) {
            return true;
        }
    }
    false
}

fn main() -> anyhow::Result<()> {
    let path = env::args().nth(1).ok_or(anyhow::anyhow!("no path"))?;
    let read = read_to_string(path)?;
    let input = read.lines();

    let mut parsed = input.map(|l| {
        l.split(" ")
            .map(|x| -> i32 { x.parse().unwrap() })
            .collect_vec()
    });

    let p1 = parsed.clone().filter(|it| is_safe(it.iter())).count();
    let p2 = parsed
        .by_ref()
        .filter(|it| with_dampener::<i32>(it))
        .count();

    println!("{}", p1);
    println!("{}", p2);
    Ok(())
}
