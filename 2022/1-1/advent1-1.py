def main():
    file = open("advent1-1-input.txt", "r")
    lines = file.readlines()

    top1 = 0
    top2 = 0
    top3 = 0
    curr = 0

    for line in lines:
        if line == "\n":
            if curr > top1:
                top3 = top2
                top2 = top1
                top1 = curr
            elif curr > top2:
                top3 = top2
                top2 = curr
            elif curr > top3:
                top3 = curr
            curr = 0
        else:
            curr += int(line)

    if curr > top1:
        top3 = top2
        top2 = top1
        top1 = curr
    elif curr > top2:
        top3 = top2
        top2 = curr
    elif curr > top3:
        top3 = curr

    total = top1 + top2 + top3
    print(total)


main()