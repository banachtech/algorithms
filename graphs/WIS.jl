# Maximum Weight Independent Set
# Easier read with 0-based indexing
# i+1 index refers to vertex i
function WIS(W)
    n = length(W)
    A = Vector{Int64}(undef, n + 1)
    A[1] = 0
    A[2] = W[1]
    for i in 3:length(A)
        #two cases: optimal solution includes current vertex or does not
        A[i] = max(A[i-1], A[i-2] + W[i-1])
    end
    S = Int64[]
    i = n + 1
    while i > 2
        if A[i-1] < A[i-2] + W[i-1] #case 1, current vertex was selected
            push!(S, i - 1)
            i -= 2
        else
            i -= 1  #case 2, current vextex was not selected
        end
    end
    if i == 2
        push!(S, W[1])
    end
    return A[n+1], S
end
