MAX = 1000005
_data = [-1] * MAX
pre = [-1] * MAX
nxt = [-1] * MAX
unused = 1


def insert(addr, value):
    global unused
    _data[unused] = value
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
    _cursor = nxt[0]

    while _cursor != -1:
        print(f"{_data[_cursor]}", end=end)
        _cursor = nxt[_cursor]

    print("")


data = input()
cursor = 0

for i in data:
    insert(cursor, i)
    cursor += 1

n = int(input())
list_len = len(data)

while n > 0:
    cmd = input()
    n -= 1

    if cmd[0] == "P":
        insert(cursor, cmd[2])
        cursor = nxt[cursor]
    elif cmd[0] == "L":
        if pre[cursor] != -1:
            cursor = pre[cursor]
    elif cmd[0] == "D":
        if nxt[cursor] != -1:
            cursor = nxt[cursor]
    elif cmd[0] == "B":
        if cursor == 0:
            continue
        erase(cursor)
        cursor = pre[cursor]

travel(end="")
