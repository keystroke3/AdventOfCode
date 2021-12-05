from collections import Counter


def get_columns(rows):
    columns = []
    for binary in rows:
        for index, bit in enumerate(binary):
            if len(columns) < index + 1:
                columns.append([bit])
            else:
                columns[index].append(bit)
    return columns


def eliminate_rows(rows, prefered: int, pos: int = 0):
    focus_col = get_columns(rows)[pos]
    count = Counter(focus_col)
    zeros = count["0"]
    ones = count["1"]
    if zeros == ones:
        selector = prefered
    elif zeros > ones:
        if prefered == 1:
            selector = 0
        else:
            selector = 1
    else:
        if prefered == 1:
            selector = 1
        else:
            selector = 0
    new_rows = [rows[i] for i, j in enumerate(focus_col) if j == str(selector)]
    if len(new_rows) == 1:
        return new_rows[0]
    return eliminate_rows(new_rows, prefered, pos + 1)


def part1():
    with open("input.txt") as logs:
        columns = get_columns(logs)
    gamma_rate = "".join([max(col, key=col.count) for col in columns])
    epsilon_rate = "".join([min(col, key=col.count) for col in columns])
    print((int(gamma_rate, 2) * int(epsilon_rate, 2)))


def part2():
    with open("input.txt") as logs:
        logs = [i.strip() for i in logs]
        oxygen = int(eliminate_rows(logs, 1), 2)
        co2 = int(eliminate_rows(logs, 0), 2)
        print(oxygen * co2)
