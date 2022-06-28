# Parse input
str = readlines()
n = parse(Int, str[1])
st = zeros(n)
ed = zeros(n)
for i = 1:n
    st[i], ed[i] = parse.(Int, split(str[i+1]))
end

pts = Any[]
while length(ed) > 0
    m = minimum(ed)
    push!(pts, m)
    global ed = ed[st.>m]
    global st = st[st.>m]
end
println(pts)