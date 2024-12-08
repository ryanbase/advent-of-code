fn main() {
    let lines = include_str!("../input.txt").split("\n");
    let mut total = 0;
    for line in lines {
        let mut pair = line.split(",");
        let mut range1 = pair.next().unwrap().split("-");
        let mut range2 = pair.next().unwrap().split("-");
        let min1 = range1.next().unwrap().parse::<i32>().unwrap();
        let max1 = range1.next().unwrap().parse::<i32>().unwrap();
        let min2 = range2.next().unwrap().parse::<i32>().unwrap();
        let max2 = range2.next().unwrap().parse::<i32>().unwrap();
        if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
            total += 1;
        }
    }
    println!("{}", total);
}
