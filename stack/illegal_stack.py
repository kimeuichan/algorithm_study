MAX = 10006
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
        return -1

    pos -= 1
    val = _data[pos]

    return val


def size():
    return pos


def top():
    if pos <= 0:
        return -1
    return _data[pos-1]


def empty():
    if pos == 0:
        return 1

    return 0


if __name__ == '__main__':
    push(1)  # 1
    print(top())  # 1`
    print(pop())  # 1

    push(2)  # 2
    push(3)  # 2 3
    print(size())  # 2
    print(pop())  # 3

    push(4)  # 2 4
    print(top())  # 4
    print(pop())  # 4
    print(pop())  # 2
    print(pop())  # -1
