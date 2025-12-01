use anyhow::Result;
use regex::Regex;
use std::{fs, i32};

const REGEX_EXPR: &str = r"(?<direction>^[R|L])(?<number>[1-9]\d*)";
const STARTING_POINT: i32 = 50;
const START_DIAL: i32 = 0;
const FULL_DIAL: i32 = 100;
const BEFORE_START: i32 = START_DIAL - 1;

fn solve_part1(input: &str) -> i32 {
    let Some(rotations) = parse(input) else {
        return 0;
    };

    match rotations.len() {
        1 => return 0,
        _ => {
            let mut counter = 0;
            let mut curr = STARTING_POINT;
            for el in rotations {
                curr = (curr + el) % FULL_DIAL;
                match curr {
                    START_DIAL => counter += 1,
                    i32::MIN..=BEFORE_START => {
                        curr += FULL_DIAL;
                    }
                    _ => {
                        // skip
                    }
                }
            }
            return counter;
        }
    }
}

fn _check_is_too_behind(el: i32, rest: i32, curr: i32) -> bool {
    el > 0 && rest > 0 && curr + rest >= FULL_DIAL
}

fn _check_is_too_ahead(el: i32, rest: i32, curr: i32) -> bool {
    el < 0 && rest > 0 && curr > START_DIAL && rest >= curr
}

fn solve_part2(input: &str) -> i32 {
    let Some(rotations) = parse(input) else {
        return 0;
    };

    match rotations.len() {
        1 => return 0,
        _ => {
            let mut counter = 0;
            let mut curr = STARTING_POINT;
            for el in rotations {
                let step = el.abs();
                let rest = step % FULL_DIAL;
                counter += step / FULL_DIAL;
                
                if _check_is_too_behind(el, rest, curr) || _check_is_too_ahead(el, rest, curr) {
                    counter += 1;
                }

                curr = (curr + el) % FULL_DIAL;
                if curr < 0 {
                    curr += FULL_DIAL;
                }
            }
            return counter;
        }
    }
}

fn transform_rotation(rotation: &str) -> Result<i32, anyhow::Error> {
    let re = Regex::new(REGEX_EXPR).unwrap();
    let Some(r) = re.captures(rotation) else {
        return Ok(0);
    };

    let num = &r["number"].parse::<i32>().unwrap();
    if r["direction"].eq("L") {
        return Ok(num * -1);
    }
    Ok(*num)
}

fn parse(input: &str) -> Option<Vec<i32>> {
    let Some(rotations): Option<Vec<i32>> = input
        .lines()
        .map(|line| transform_rotation(line))
        .collect::<Result<_, _>>()
        .ok()
    else {
        return None;
    };

    Some(rotations)
}

fn main() -> Result<()> {
    let filename = format!("../inputs/day01");
    let input = fs::read_to_string(&filename)
        .map_err(|e| anyhow::anyhow!("Failed to read {}: {}", filename, e))?;

    // println!("Result part 1: {}", solve_part1(input.as_str()));
    println!("Result part 2: {}", solve_part2(input.as_str()));
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82";

    #[test]
    fn test_part1() {
        assert_eq!(solve_part1(INPUT), 3);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve_part2(INPUT), 6);
    }
}
