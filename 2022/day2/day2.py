"""
Advent Of Code day 2: Rock Paper Scissors
"""
from typing import TextIO

# Shapes
# Rock = A
# Paper = B
# Scissors = C

# Strengths
# Rock > Scissors (A > C)
# Paper > Rock (B > A)
# Scissors > Paper (C > B)

WIN_POINTS = 6
DRAW_POINTS = 3
LOSS_POINTS = 0


def part1_round_points(shapes: str) -> int:
    SHAPE_POINTS = dict(X=1, Y=2, Z=3)
    WINS = ("A Y", "B Z", "C X")
    LOSSES = ("A Z", "B X", "C Y")

    if shapes in WINS:
        result_point = WIN_POINTS
    elif shapes in LOSSES:
        result_point = LOSS_POINTS
    else:
        result_point = DRAW_POINTS
    shape_point = SHAPE_POINTS[shapes[-1]]

    return shape_point + result_point


def part2_round_points(shapes: str) -> int:
    LOOSE = "X"
    DRAW = "Y"
    WIN = "Z"

    LOSSES = dict(A="C", B="A", C="B")
    WINS = dict(A="B", B="C", C="A")
    SHAPE_POINTS = dict(A=1, B=2, C=3)
    goal = shapes[-1]
    opponent_shape = shapes[0]
    if goal == LOOSE:
        shape = LOSSES[opponent_shape]
        result_point = LOSS_POINTS
        shape_point = SHAPE_POINTS[shape]
    elif goal == WIN:
        shape = WINS[opponent_shape]
        result_point = WIN_POINTS
        shape_point = SHAPE_POINTS[shape]
    else:
        shape = opponent_shape
        result_point = DRAW_POINTS
        shape_point = SHAPE_POINTS[shape]
    return result_point + shape_point


def get_score(f: TextIO, part: int = 1) -> int:
    score = 0
    while True:
        shapes = f.readline().strip()
        if shapes == "":
            return score
        if part == 1:
            score += part1_round_points(shapes)
        else:
            score += part2_round_points(shapes)


with open("input.txt") as f:
    print(get_score(f, 2))
