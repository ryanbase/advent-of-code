fn main() {
    let mut total = 0;
    for line in include_str!("../input/input.txt").split("\n") {
        let mut comp1_lower: [i32; 26] = [0; 26];
        let mut comp1_upper: [i32; 26] = [0; 26];
        let mut comp2_lower: [i32; 26] = [0; 26];
        let mut comp2_upper: [i32; 26] = [0; 26];
        let half = line.len() / 2;
        for (i, c) in line.chars().enumerate() {
            if i < half {
                if c.is_lowercase() {
                    comp1_lower[(c as usize) - ('a' as usize)] += 1;
                } else {
                    comp1_upper[(c as usize) - ('A' as usize)] += 1;
                }
            } else {
                if c.is_lowercase() {
                    comp2_lower[(c as usize) - ('a' as usize)] += 1;
                } else {
                    comp2_upper[(c as usize) - ('A' as usize)] += 1;
                }
            }
        }
        for i in 0..26 {
            if comp1_lower[i] > 0 && comp2_lower[i] > 0 {
                total += i + 1;
            }
            if comp1_upper[i] > 0 && comp2_upper[i] > 0 {
                total += i + 27;
            }
        }
    }
    println!("{}", total);
}
