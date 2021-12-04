"""Day 1 is finding the number of depth increases registered by the submarine sonar
"""


def count_increased_depths(depths: list):
    increased_count = 0
    last_depth = 0
    for depth in depths:
        if int(depth) > last_depth:
            increased_count += 1
        last_depth = int(depth)

    return increased_count - 1


def part1():
    with open("input.txt") as depths:
        print(count_increased_depths(depths))


def part2():
    with open("input.txt") as depths:
        d = list(depths)
        sliced_depths = []
        for count, depth in enumerate(d):
            sliced_depths.append(sum([int(d) for d in d[count : count + 3]]))  # noqa
        print(count_increased_depths(sliced_depths))
