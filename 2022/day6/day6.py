def find_marker(signal: str, marker: str = "message") -> tuple:
    index = 13 if marker == "message" else 3
    sequence = signal[: index + 1]
    while True:
        index += 1
        if len(set(sequence)) == len(sequence):
            return index, sequence
        sequence = sequence[1:] + signal[index]


with open("input.txt") as f:
    signal = f.readlines()[0]
    print(find_marker(signal, "message"))
