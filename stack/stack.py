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


n = int(input())

while n > 0:
    cmd = input()

    if cmd.startswith("push"):
        _, number = cmd.split()
        push(number)
    elif cmd == "top":
        print(top())
    elif cmd == "empty":
        print(empty())
    elif cmd == "pop":
        print(pop())
    elif cmd == "size":
        print(size())

    n -= 1
