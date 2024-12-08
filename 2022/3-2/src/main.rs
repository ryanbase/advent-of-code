fn main() {
    let mut total = 0;
    let mut lines = include_str!("../input.txt").split("\n");
    let count = lines.clone().count();
    if count < 3 {
        println!("Invalid input");
        return;
    }
    let mut i = 0;
    while i < count {
        let arr1 = get_array(lines.next().unwrap());
        let arr2 = get_array(lines.next().unwrap());
        let arr3 = get_array(lines.next().unwrap());
        for j in 0..52 {
            if arr1[j] > 0 && arr2[j] > 0 && arr3[j] > 0 {
                total += j + 1;
            }
        }
        i += 3;
    }
    println!("{}", total);
}

fn get_array(line: &str) -> [i32; 52] {
    let mut arr: [i32; 52] = [0; 52];
    for c in line.chars() {
        if c.is_lowercase() {
            arr[(c as usize) - ('a' as usize)] += 1;
        } else {
            arr[((c as usize) - ('A' as usize)) + 26] += 1;
        }
    }
    return arr;
}
