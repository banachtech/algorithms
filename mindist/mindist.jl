using LinearAlgebra, Random, BenchmarkTools

# generate random points
function fakedata(n)
    p = Vector[]
    for i in 1:n
        push!(p, rand(-1000_000_000:1000_000_000, 2))
    end
    sort!(p, lt=(x, y) -> isless(x[1], y[1]))
    return p
end

function foo(x)
    n = length(x)
    if n == 1
        return Inf
    end
    if n == 2
        return norm(x[1] .- x[2], 2)
    end
    if n == 3
        return minimum([norm(x[1] .- x[2], 2), norm(x[1] .- x[3], 2), norm(x[3] .- x[2], 2)])
    end
    d = Inf
    for i in 1:min(7, n)
        for j in i+1:min(7, n)
            d = min(d, norm(x[i] - x[j]), 2)
        end
    end
    return d
end

function mindist(x)
    n = length(x)
    if n == 1
        return Inf
    end
    if n == 2
        return norm(x[1] .- x[2], 2)
    end
    if n == 3
        return minimum([norm(x[1] .- x[2], 2), norm(x[1] .- x[3], 2), norm(x[3] .- x[2], 2)])
    end
    m = floor(Int, n / 2)
    mx = x[m][1]
    dleft = mindist(x[1:m])
    dright = mindist(x[m+1:end])
    d = min(dleft, dright)
    z = [r for r in x if abs(r[1] - mx) <= d]
    d = min(d, foo(z))
    return d
end


# test
p = fakedata(100000)
@elapsed mindist(p)

# brute force check
@elapsed begin
    tmp = Inf
    for i in 1:length(p)
        for j in i+1:length(p)
            tmp = min(tmp, norm(p[i] - p[j]), 2)
        end
    end
end
tmp