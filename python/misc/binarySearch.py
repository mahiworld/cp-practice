from numpy import append


def recBinarySearch(array, ele, low, high):
    if low == high:
        if ele == array[low]:
            return low
        else:
            return None
    else:
        mid = int((low+high)/2)
        if ele == array[mid]:
            return mid
        else:
            if ele < array[mid]:
                return recBinarySearch(array,ele,low, mid-1)
            else:
                return recBinarySearch(array, ele, mid+1, high)
        
        return None

def nonrecBinarySearch(array, ele, low, high):
    if low == high:
        if ele == array[low]:
            return low
        else: 
            return None
    
    else:
        while(low<high):
            mid = int(low + (high-low)/2)
            if ele == array[mid]:
                return mid
            else:
                if ele < array[mid]:
                    high = mid-1
                else:
                    low = mid+1 

    return None

sortedArray = input("Enter sorrted array elements: ").split(',')
print(sortedArray)
sortedIntArray = []
for ele in sortedArray:
    sortedIntArray.append(int(ele))
print(sortedIntArray)
eleToSearch = int(input("Enter element to be searched: "))
print(eleToSearch)

i = 0
j = len(sortedIntArray)-1

print("\n")
result = recBinarySearch(sortedIntArray, eleToSearch, i, j)
if result != None:
    print("Element is present at index "+ str(result))
else:
    print("Not found")

result = nonrecBinarySearch(sortedIntArray, eleToSearch, i, j)
if result != None:
    print("Element is present at index "+ str(result))
else:
    print("Not found")