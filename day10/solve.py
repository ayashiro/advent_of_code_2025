import sys
def convert(v, l):
    vv = set(v)
    ans = 0

    for i in range(l):
        ans <<= 1
        if i in vv:
            ans |= 1
    return ans


    
with open(sys.argv[1]) as f :
    ans = 0
    for line in f.readlines():
        elements = line.strip().split()
        target = [1 if v == '#' else 0 for v in elements[0][1:-1]]
        mp = [tuple(map(int, v[1:-1].split(","))) for v in elements[1:-1]]

        goal = 0
        for v in target:
            goal <<= 1
            goal |= v
        directions = [convert(v, len(target)) for v in mp]
        dist = [-1] * (1<<len(target))
        from collections import deque
        qu = deque()
        qu.append((0, 0))
        while qu:
            now, step = qu.popleft()

            if dist[now] != -1 :
                continue
            dist[now] = step
            for v in directions :

                nxt = (now ^ v)
                if dist[nxt] == -1:
                    qu.append((nxt, step + 1))
        counter = [0] * len(target)
        for v in mp:
            for k in v:
                counter[k] += 1
        ans += dist[goal]                
        goal = tuple(map(int, elements[-1][1:-1].split(",")))
        print(goal, mp)
        t = 1
        for v in goal: t *= v
        print(t)

    print(ans)
