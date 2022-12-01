"""Get product of submarine horizontal position  and depth after a series of movements
"""


class Position:
    horizontal = 0
    depth = 0
    aim = 0


pos = Position()


def change_postition_part1(direction: str, steps: int):
    if direction == "forward":
        pos.horizontal += steps
    elif direction == "up":
        pos.depth -= steps
    else:
        pos.depth += steps


def change_postition_part2(direction: str, steps: int):
    if direction == "forward":
        pos.horizontal += steps
        pos.depth += pos.aim * steps
    elif direction == "up":
        pos.aim -= steps
    elif direction == "down":
        pos.aim += steps


def navigate(part: int = 1):
    with open("input.txt") as movements:
        for movement in movements:
            direction, steps = movement.split(" ")
            if part == 2:
                change_postition_part2(direction, int(steps))
            else:
                change_postition_part1(direction, int(steps))


def part1():
    navigate()
    print(f"hor:{pos.horizontal}, depth:{pos.depth}")
    print(pos.horizontal * pos.depth)


def part2():
    navigate(2)
    print(f"hor:{pos.horizontal}, depth:{pos.depth}, aim:{pos.aim}")
    print(pos.horizontal * pos.depth)
