def bubble_sort(array):
    for i in range(len(array)):
        print("i:", i)
        for j in range(len(array) - 1):
            print("j:", j)
            if array[j] < array[j + 1]:
                array[j], array[j + 1] = array[j + 1], array[j]
    print(a)


a = [1, 2, 3, 4, 5]
print(a)
bubble_sort(a)
