def matrixOperate(arr[][], m ,n):
    for i in range(0, m-1):
        for j in range(0,n-1):
            temp[i][j] = arr[i][j]

    for i in range(0,m-1):
        for j in range(0,n-1):
            if(arr[i][j] == '0'):
                for i in range(0,m-1):
                    temp[i][j] = 0
                for j in range(0,n-1):
                    temp[i][j] = 0

    for i in range(0,m-1):
        for j in range(0,n-1):
            arr[i][j] = temp[i][j]