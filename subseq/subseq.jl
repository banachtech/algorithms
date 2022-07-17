
function subsq(a, b)
    Y = foo(a, b)
    I = bar(Y)
    if isempty(I)
        println("no common elements")
        return nothing
    end
    println("longest sub-sequence: ", reverse(digits(b))[I])
    return length(I)
end

function foo(a, b)
    a = reverse(digits(a))
    b = reverse(digits(b))

    n = length(a)
    m = length(b)

    X = falses(n, m)

    for i in 1:n
        for j in 1:m
            X[i, j] = a[i] == b[j]
        end
    end
    return X
end

function update!(S, j)
    if isempty(S)
        append!(S, j)
    else
        if S[end] < j
            append!(S, j)
        else
            S[end] = j
        end
    end
end

function bar(Y)
    S = Int64[]
    for i in 1:size(Y, 1)
        for j in 1:size(Y, 2)
            if Y[i, j]
                update!(S, j)
                S = unique(S)
                if length(S) == size(Y, 2)
                    return S
                end
            end
        end
    end
    return S
end
