fn main() {
    let mut score: i32 = 0;
    for line in include_str!("../input/advent2-input.txt").split("\n") {
        let opp_opt: Option<char> = line.chars().nth(0);
        let you_opt: Option<char> = line.chars().nth(2);
        if opp_opt == None || you_opt == None {
            continue;
        }
        let opp: char = opp_opt.unwrap();
        let result: char = you_opt.unwrap();

        if result == 'X' {
            // X = lose
            score += get_lose_score(opp);
        } else if result == 'Y' {
            // Y = tie
            score += get_tie_score(opp) + 3;
        } else if result == 'Z' {
            // Z = win
            score += get_win_score(opp) + 6;
        }
    }
    println!("{}", score);
}

fn get_lose_score(opp: char) -> i32 {
    if opp == 'A' {
        return 3;
    }
    if opp == 'B' {
        return 1;
    }
    return 2;
}

fn get_tie_score(opp: char) -> i32 {
    if opp == 'A' {
        return 1;
    }
    if opp == 'B' {
        return 2;
    }
    return 3;
}

fn get_win_score(opp: char) -> i32 {
    if opp == 'A' {
        return 2;
    }
    if opp == 'B' {
        return 3;
    }
    return 1;
}
