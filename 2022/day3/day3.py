from typing import TextIO
from string import ascii_letters


def get_group_bags(f: TextIO, group_size: int = 3) -> list:
    while True:
        yield [f.readline().strip() for _ in range(group_size)]


def part2_priority_sum(f: TextIO) -> int:
    group_bags = get_group_bags(f)
    priority_sum = 0
    while True:
        bags = next(group_bags)
        if not any(bags):
            return priority_sum
        common_item = list(set(bags[0]) & set(bags[1]) & set(bags[2]))[0]
        priority_sum += ascii_letters.index(common_item) + 1


def part1_priority_sum(f: TextIO) -> int:
    priority_sum = 0
    while True:
        items = f.readline().strip()
        if items == "":
            return priority_sum
        compartment_1 = set(items[: len(items) // 2])
        compartment_2 = set(items[len(items) // 2 :])
        common_item = list(compartment_1.intersection(compartment_2))[0]
        priority_sum += ascii_letters.index(common_item) + 1


with open("test_input.txt") as f:
    print(part2_priority_sum(f))
