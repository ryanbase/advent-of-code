fn main() {
    let mut score: i32 = 0;
    for line in include_str!("../input/advent2-input.txt").split("\n") {
        let opp_opt: Option<char> = line.chars().nth(0);
        let you_opt: Option<char> = line.chars().nth(2);
        if opp_opt == None || you_opt == None {
            continue;
        }
        let opp: char = opp_opt.unwrap();
        let you: char = you_opt.unwrap();

        if you == 'X' {
            score += 1;
        } else if you == 'Y' {
            score += 2;
        } else if you == 'Z' {
            score += 3;
        }

        if (you == 'X' && opp == 'A') || (you == 'Y' && opp == 'B') || (you == 'Z' && opp == 'C') {
            score += 3;
        } else if (you == 'X' && opp == 'C')
            || (you == 'Y' && opp == 'A')
            || (you == 'Z' && opp == 'B')
        {
            score += 6
        }
    }
    println!("{}", score);
}
