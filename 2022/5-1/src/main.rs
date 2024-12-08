fn main() {
    let mut lines = include_str!("../input.txt").split("\n");
    let mut stacks;
    for (i, line) in lines.enumerate() {
        if i == 0 {
            let size = (line.clone().chars().count() + 1) / 4;
            stacks = [Vec::new(); size];
        }
    }

    let mut total = 0;
    let mut setup = true;
    // for line in lines {
    //     if line.is_empty() {
    //         setup = false;
    //         continue;
    //     }
    //     if setup {
    //         line.
    //     } else {

    //     }
    // }
    println!("{}", total);
}
