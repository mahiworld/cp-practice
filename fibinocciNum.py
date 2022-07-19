def fibinocci(num):
    if num ==0 or num ==1:
        return num
    else:
        return fibinocci(num-1) + fibinocci(num-2)

def nonrecFibinocci(num):
    if num == 0 or num == 1:
        return num
    else:
        prev1 = 0
        prev2 = 1
        curr = None
        while(num>1):
            curr = prev1 + prev2
            prev1 = prev2
            prev2 = curr
            num = num -1
        return curr



print(fibinocci(int(input("Enter Fib number to get: "))))
print(nonrecFibinocci(int(input("Enter Fib number to get: "))))