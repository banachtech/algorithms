# Parse input
str = readlines()
d, m = parse(Int, str[1]), parse(Int, str[2])
n = parse(Int, str[3])
s = parse.(Int, split(str[4]))

# Compute number of stops
dist = diff([0; s; d])
n = length(dist) + 1
i = 1
traveled = dist[1]
stops = Any[]
stop = nothing
println(dist)
while i + 1 != n
    global traveled = dist[i]
    while ((traveled <= m) && (i + 2 <= n))
        global stop = i
        traveled += dist[i+1]
        global i = i + 1
    end
    push!(stops, stop)
    println(stops)
end
println(length(stops))
