tbl = {}

def fibonacci(n):
	if n not in tbl:
		if n <= 1:
			tbl[n]=n
		else:
			tbl[n] = fibonacci(n-1)+fibonacci(n-2)
	return tbl[n]

if __name__ == "__main__":
	n = int(input())
	print(fibonacci(n))
