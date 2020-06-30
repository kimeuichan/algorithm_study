MAX = 100000
data = [-1] * MAX
pre = [-1] * MAX
nxt = [-1] * MAX
unused = 1


def insert(addr, value):
    global unused
    data[unused] = value
    pre[unused] = addr
    nxt[unused] = nxt[addr]

    if nxt[addr] != -1:
        pre[nxt[addr]] = unused

    nxt[addr] = unused

    unused += 1


def erase(addr):
    nxt[pre[addr]] = nxt[addr]

    if nxt[addr] != -1:
        pre[nxt[addr]] = pre[addr]


def travel(end=","):
    cursor = nxt[0]

    while cursor != -1:
        print(f"{data[cursor]}", end=end)
        cursor = nxt[cursor]

    print("")


if __name__ == '__main__':
    insert(0, 10)
    travel()  # 10
    insert(1, 20)
    travel()  # 10 20
    insert(2, 30)
    travel()  # 10 20 30

    erase(1)
    travel()  # 20 30

    insert(2, 40)
    travel()  # 20 40 30

