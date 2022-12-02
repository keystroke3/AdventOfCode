"""
Advent of Code 2022 Day1: Calories
"""
from typing import TextIO

# Part 1
def highest_calory(f: TextIO) -> int:
    calory_count = 0
    top = 0
    while True:
        calories = f.readline()
        if calories == "\n":
            if calory_count > top:
                top = calory_count
            calory_count = 0
            continue
        elif calories == "":
            return top
        calory_count += int(calories)


# Part 2
def top3_calories(f: TextIO) -> int:
    calory_count = 0
    top1 = 0
    top2 = 0
    top3 = 0
    while True:
        calories = f.readline()
        try:
            calory_count += int(calories)
        except ValueError:
            if calories in ("\n", ""):
                if calory_count > top1:
                    top1, top2, top3 = calory_count, top1, top2
                elif calory_count > top2:
                    top2, top3 = calory_count, top2
                elif calory_count > top3:
                    top3 = calory_count
                calory_count = 0
            if calories == "":
                return sum(top1, top2, top3)


# with open("input.txt") as f:
#     print(highest_calory(f))

# with open("input.txt") as f:
#     print(top3_calories(f))
