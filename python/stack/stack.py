
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
    string = input()
    cmd = string.split()
    if cmd[0] == 'push':
        push(cmd[1])
    elif cmd[0] == "top":
        print(top())
    elif cmd[0] == "empty":
        print(empty())
    elif cmd[0] == "pop":
        print(pop())
    elif cmd[0] == "size":
        print(size())

    n -= 1
