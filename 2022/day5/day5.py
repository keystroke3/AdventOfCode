from typing import TextIO


def get_crates(f: TextIO) -> dict:
    layers = []
    while True:
        layer = f.readline().strip("\n")
        print(layer)
        if layer == "":
            layers.pop()  # remove the with column labels
            break
        crates_at_layer = (
            layer.replace("]     [", "],_,[")
            .replace("] [", "],[")
            .replace("] ", "], ")
            .replace("    ", "_,")
            .replace("[", "")
            .replace("]", "")
            .strip(",")
            .split(",")
        )  # replace all blank spots with _ and separator spaces with , and split
        layers.append(crates_at_layer)

    crates = []
    for crate in range(len(layers[0])):
        column = [l[crate] for l in layers]
        column.reverse()
        crates.append(column)
    rotated_crates = [list("".join(c).replace("_", "")) for c in crates]
    return {i: crate for i, crate in enumerate(rotated_crates, start=1)}


def move_crates(f: TextIO, crates: dict, crane: int = 9001) -> dict:
    motion = f.readline()
    if motion == "":
        return crates
    motion = motion.split(" ")
    num = int(motion[1])
    source = int(motion[3])
    target = int(motion[-1])

    to_move = crates[source][-num:]
    if crane == 9000:
        to_move.reverse()
    crates[source] = crates[source][:-num]
    crates[target].extend(to_move)

    return move_crates(f, crates, crane)


with open("input.txt") as f:
    part1 = 9000
    part2 = 9001
    crates = get_crates(f)
    final_crates = move_crates(f, crates, part2)
    for c in final_crates.values():
        print(c)
    print("".join([layer[-1] for layer in final_crates.values()]))
