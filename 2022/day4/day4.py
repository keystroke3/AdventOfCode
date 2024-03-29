from typing import TextIO


def has_subset(set1: set, set2: set) -> bool:
    """If a range of assignments A is fully contained in range B
    then A is a subset of B and vice versa"""
    return set1.issubset(set2) or set2.issubset(set1)


def has_overlap(set1: set, set2: set) -> bool:
    return bool(set1.intersection(set2))


def get_assignment_sets(line: str) -> list:
    elves = line.split(",")
    assignments = [e.split("-") for e in elves]

    # this works for arbitrary number of elves, not just a pair
    assignment_sets = [set(range(int(a[0]), int(a[-1]) + 1)) for a in assignments]

    return assignment_sets


def part1_and_2(f: TextIO) -> int:
    overlap_count = 0
    subset_count = 0
    while True:
        line = f.readline().strip()
        if line == "":
            return f"{overlap_count=}, {subset_count=}"
        _sets = get_assignment_sets(line)
        if has_overlap(*_sets):
            overlap_count += 1
        if has_subset(*_sets):
            subset_count += 1


with open("input.txt") as f:
    print(part1_and_2(f))
