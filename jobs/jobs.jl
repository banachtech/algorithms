p = [50, 20, 30, 25, 15]
d = [2, 1, 2, 1, 3]
loc = zeros(5)
profit = 0
idx = 1:length(p)
while length(p) >= 1
    println(p)
    i = argmax(p)
    if loc[d[i]] == 0
        loc[d[i]] = p[i]
    end
    p = remove(p, i)
    d = remove(d, i)
end

println(loc)

function remove(x, i)
    if i == 1
        return x[2:end]
    elseif i == length(x)
        return x[1:end-1]
    end
    return append!(x[1:i-1], x[i+1:end])
end