using LinearAlgebra


function coin(d)
    d = reverse(digits(d))
    n = length(d)
    P = Matrix{Int64}(undef, n, n)
    for i = 1:n, j = 1:n
        P[i,j] = 0
    end
    for i in 1:n
        P[i,i] = d[i]
    end
    for i in 1:n-1
        P[i,i+1] = max(d[i], d[i+1])
    end
    S = Int64[]
    for i = 1:n, j = 1:n
        if i == j || j == i+1
            continue
        else
            #tmp = [d[i]+P[min(n,i+2),j], d[i] + P[min(n,i+1),max(1,j-1)], d[j]+P[i,max(1,j-2)], d[j]+P[min(n,i+1), max(1,j-1)]]
            tmp = [d[i]+P[i+2,j], d[i] + P[i+1,j-1], d[j]+P[i,max(1,j-2)], d[j]+P[min(n,i+1), max(1,j-1)]]
            P[i,j] = maximum(tmp)
        end
    end
    return P
end

