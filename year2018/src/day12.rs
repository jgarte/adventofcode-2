use std::collections::VecDeque;

use base::{Part, Solver};

pub fn get_solver() -> Box<dyn Solver> {
    Box::new(Day12)
}

struct Day12;

impl Solver for Day12 {
    fn solve(&self, part: Part, input: &str) -> Result<String, String> {
        let (pots, map) = parse_input(input);
        match part {
            Part::One => Err("day 12 part 1 not yet implemented".to_string()),
            Part::Two => Err("day 12 part 2 not yet implemented".to_string()),
        }
    }
}

fn parse_input(input: &str) -> (VecDeque<usize>, Vec<usize>) {
    let mut lines = input.lines();
    let initial_state_line = lines.next().unwrap();
    let parts = initial_state_line.split(": ").collect::<Vec<&str>>();
    let initial_bits = bitstring_to_bits(parts[1]);
    lines.next();
    let rest = lines.collect::<Vec<&str>>();
    let mut map = vec![0; 1 << 5];
    for &line in &rest {
        let (idx, bit) = parse_pattern(line);
        map[idx] = bit;
    }
    (VecDeque::from(initial_bits), map)
}

fn parse_pattern(line: &str) -> (usize, usize) {
    let parts = line.split(" => ").collect::<Vec<&str>>();
    let pattern = bitstring_to_bits(parts[0]);
    let output = bitstring_to_bits(parts[1]);
    (bits_to_usize(&pattern), bits_to_usize(&output))
}

fn bits_to_usize(bits: &[usize]) -> usize {
    bits.iter().fold(0, |x, &bit| (x << 1) | bit)
}

fn bitstring_to_bits(bitstring: &str) -> Vec<usize> {
    bitstring
        .chars()
        .map(|c| match c {
            '.' => 0,
            '#' => 1,
            _ => panic!("invalid char: {}", c),
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    mod part1 {
        use super::*;

        #[test]
        fn with_input() {
            let solver = get_solver();
            let input = include_str!("../../inputs/2018/12").trim();
            let expected = "expected output";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }

        #[test]
        fn example() {
            let solver = get_solver();
            let input = "\
initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #\
            ";
            let expected = "325";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }
    }

    mod part2 {
        use super::*;

        #[test]
        fn with_input() {
            let solver = get_solver();
            let input = include_str!("../../inputs/2018/12").trim();
            let expected = "expected output";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }

        #[test]
        fn example() {
            let solver = get_solver();
            let input = "put some input here";
            let expected = "expected output";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }
    }
}