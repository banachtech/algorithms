using LinearAlgebra


function coin(d)
    d = reverse(digits(d))
    n = length(d)
    P = zeros(Int, n, n)
    for i in 1:n
        P[i, i] = d[i]
        if i < n
            P[i, i+1] = max(d[i], d[i+1])
        end
    end
    for i = 1:n, j = i+2:n
        tmp = [d[i] + P[min(n, i + 2), j], d[i] + P[min(n, i + 1), max(1, j - 1)], d[j] + P[i, max(1, j - 2)], d[j] + P[min(n, i + 1), max(1, j - 1)]]
        P[i, j] = maximum(tmp)
    end
    return P
end


function coin2(d)
    d = reverse(digits(d))
    n = length(d)
    P = zeros(Int, n, n)
    for i = 1:n, j = i:n
        if i == j
            P[i, j] = d[i]
        elseif j == i + 1
            P[i, j] = max(d[i], d[j])
        else
            tmp = [d[i] + P[min(n, i + 2), j], d[i] + P[min(n, i + 1), max(1, j - 1)], d[j] + P[i, max(1, j - 2)], d[j] + P[min(n, i + 1), max(1, j - 1)]]
            P[i, j] = maximum(tmp)
        end
    end
    return P
end