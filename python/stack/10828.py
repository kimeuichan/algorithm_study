from collections import deque

n = int(input())

stack = deque()

while n > 0:
    string = input()
    cmd = string.split()
    if cmd[0] == 'push':
        stack.append(cmd[1])
    elif cmd[0] == "top":
        if stack:
            print(stack[-1])
        else:
            print(-1)
    elif cmd[0] == "empty":
        if stack:
            print(0)
        else:
            print(1)
    elif cmd[0] == "pop":
        if stack:
            print(stack.pop())
        else:
            print(-1)
    elif cmd[0] == "size":
        print(len(stack))

    n -= 1
