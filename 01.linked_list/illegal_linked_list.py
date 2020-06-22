MAX = 100000
data = [-1] * MAX
pre = [-1] * MAX
nxt = [-1] * MAX
unused = 1


def insert(addr: int, value: int) -> None:
    global unused

    data[unused] = value
    pre[unused] = addr
    nxt[unused] = nxt[addr]

    if nxt[addr] != -1:
        pre[nxt[addr]] = unused

    nxt[addr] = unused

    unused += 1


def erase(addr: int) -> None:
    nxt[pre[addr]] = nxt[addr]

    if nxt[addr] != -1:
        pre[nxt[addr]] = pre[addr]


def travel() -> None:
    cursor = nxt[0]

    while cursor != -1:
        print(f"{data[cursor]}  ", end="  ")
        cursor = nxt[cursor]


if __name__ == '__main__':
    insert(0, 10)
    travel()
