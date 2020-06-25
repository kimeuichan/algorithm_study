MAX = 10000
_data = [0] * MAX
pos = 0


def push(val):
    global pos
    if pos >= MAX:
        return

    _data[pos] = val
    pos += 1


def pop():
    global pos
    if pos <= 0:
        return

    pos -= 1
    val = _data[pos]

    return val


if __name__ == '__main__':
    push(1)  # 1
    print(pop())  # 1

    push(2)  # 2
    push(3)  # 2 3
    print(pop())  # 3

    push(4)  # 2 4
    print(pop())  # 4
    print(pop())  # 2
    print(pop())  # None
