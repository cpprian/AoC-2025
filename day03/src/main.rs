use std::fs;

use anyhow::Result;
use regex::Regex;

const REGEX_EXPR: &str = r"\d";

fn sum_digits(left: i32, right: i32) -> i32 {
    left * 10 + right
}

fn find_highest_joltage(data: Vec<i32>) -> i32 {
    let n = data.len();

    if n == 0 {
        return 0;
    } else if n == 1 {
        return data[0];
    } else if n == 2 {
        return data[0] * 10 + data[1];
    }

    let mut max = 0;
    let mut best_left = vec![0; n];
    best_left[0] = data[0];
    for idx in 1..n {
        best_left[idx] = best_left[idx - 1].max(data[idx]);
    }
    
    for idx in 1..n {
        let temp = sum_digits(best_left[idx - 1], data[idx]);
        if temp > max {
            max = temp;
        }
    }
    max
}

fn solve_part1(input: Vec<Vec<i32>>) -> i32 {
    let mut sum: i32 = 0;
    for idm in input {
        sum += find_highest_joltage(idm);
    }
    sum
}

fn parse(input: &str) -> Result<Vec<Vec<i32>>, anyhow::Error> {
    let re = Regex::new(REGEX_EXPR).unwrap();
    let mut output: Vec<Vec<i32>> = vec![];

    for line in input.lines() {
        let data = re
            .find_iter(line)
            .map(|i| i.as_str().parse::<i32>())
            .collect::<Result<Vec<_>, _>>()?;
        output.push(data);
    }
    Ok(output)
}

fn main() -> Result<()> {
    let filename = format!("../inputs/day03");
    let input = fs::read_to_string(&filename)
        .map_err(|e| anyhow::anyhow!("Failed to read {}: {}", filename, e))?;

    match parse(input.as_str()) {
        Ok(data) => {
            println!("Solution to part 1: {}", solve_part1(data));
        }
        Err(_) => {
            panic!("Should not throw any errors!");
        }
    };
    return Ok(());
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "987654321111111
811111111111119
234234234234278
818181911112111";

    #[test]
    fn test_part1() {
        match parse(INPUT) {
            Ok(data) => {
                assert_eq!(solve_part1(data), 357);
            }
            Err(_) => {
                panic!("Should not throw any errors!");
            }
        };
    }
}
