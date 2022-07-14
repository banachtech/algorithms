# KnapSack Problem
# Given n objects with values V1, V2, ..., Vn and sizes S1, S2, ..., Sn and a knapsack with capacity C, find a packing that has maximum value.
# 1 based indexing sucks!
function knapsack(V, S, C)
    n = length(V)
    A = Matrix{Int64}(undef, n + 1, C + 1)
    A[1, :] .= 0
    for i in 2:n+1
        for c in 0:C
            if S[i-1] > c
                A[i, c+1] = A[i-1, c+1] #solution is also optimal for subproblem without object i
            else
                A[i, c+1] = max(A[i-1, c+1], A[i-1, c-S[i-1]+1] + V[i-1])
            end
        end
    end
    # Retrieve packed objects
    J = Int64[]
    c = C + 1
    for i in n+1:-1:2
        if S[i-1] <= c && A[i-1, c-S[i-1]+1] + V[i-1] > A[i-1, c]
            push!(J, i - 1)
            c = c - S[i-1]
        end
    end
    return A[n+1, C+1], J
end
