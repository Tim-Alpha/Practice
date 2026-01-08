balance = int(input())
count = 0

for note in [100, 20, 10, 5]:
    count += balance // note
    balance %= note

print(count)