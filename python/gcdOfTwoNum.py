def gcd(m, n):
    if m==0:
        return n
    elif n==0:
        return m
    else:
        return gcd(n, m%n)

def nonrecGCD(m ,n):
    if m == 0:
        return m
    elif n ==0:
        return n
    else:
        while(n%m !=0):
            temp = m
            m = n
            n = temp%n
        return m


m = int(input("Enter number1:"))
n = int(input("Enter number2:"))
print(f"Greatest Common Devisor of {m}, {n} is: {gcd(m, n)}")
print(f"Greatest Common Devisor of {m}, {n} is: {nonrecGCD(m, n)}")