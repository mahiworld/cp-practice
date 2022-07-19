def mult(m, n):
    if m==0 or n==0:
        return 0
    else:
        return m + mult(m, n-1)

def nonrecMult(m, n):
    if m ==0 or n==0:
        return 0
    else:
        r = 0
        while(n>=1):
            r += m
            n = n-1
        return r 

m = int(input("Enter first number: "))
n = int(input("Enter second number: "))
result = mult(m, n)
print(f"Multiplication of {m}, {n} is: {result}")
result = nonrecMult(m, n)
print(f"Multiplication of {m}, {n} is: {result}")