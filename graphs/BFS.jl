# Breadth First Search algorithm
# Input graph adjacency list G and vertex v
# Output list of vertices of G reachable from v
function BFS(V, E, v)
    # Initialise queue with v
    q = [v]
    # mark all vertices except v as unexplored
    status = falses(length(V))
    status[v] = true
    st = first.(E)
    ed = last.(E)
    while !isempty(q)
        s = q[1]
        if length(q) == 1
            q = []
        else
            q = q[2:end]
        end
        incidents = ed[st .== s]
        for w ∈ incidents
            if !status[w]
                status[w] = true
                push!(q, w)
            end
        end
    end
    return V[status]
end


# Adjacency matrix version
function BFS(A::BitMatrix, v)
    # Initialise queue with v
    q = [v]
    # mark all vertices except v as unexplored
    V = collect(1:size(A,1))
    status = falses(length(V))
    status[v] = true
    while !isempty(q)
        s = q[1]
        if length(q) == 1
            q = []
        else
            q = q[2:end]
        end
        incidents = V[A[s,:]]
        for w ∈ incidents
            if !status[w]
                status[w] = true
                push!(q, w)
            end
        end
    end
    return V[status]
end


# Shortest Path
# Small modification to BFS
function shortestPath(A::BitMatrix, v)
    # Initialise queue with v
    q = [v]
    # mark all vertices except v as unexplored
    V = collect(1:size(A,1))
    status = falses(length(V))
    status[v] = true
    # initialise shortest path lengths to infty
    d = fill(Inf, length(V))
    d[v] = 0
    while !isempty(q)
        s = q[1]
        if length(q) == 1
            q = []
        else
            q = q[2:end]
        end
        incidents = V[A[s,:]]
        for w ∈ incidents
            if !status[w]
                status[w] = true
                push!(q, w)
                d[w] = d[s] + 1
            end
        end
    end
    return d
end

function shortestPath(A, v, w)
    d = shortestPath(A, v)
    return d[w]
end

# Testing
using Random, LinearAlgebra

function createAdjMat(n; directed=true)
    M = BitMatrix(rand([false, true], (n,n)))
    for i in 1:n
        M[i,i] = 0
    end
    if !directed
        M = Symmetric(M)
    end
    return M
end

function createAdjList(n, m; directed=true)
    E = Vector[]
    V = collect(1:n)
    while length(E) < m
        u, v = rand(V, 2)
        if u != v
            push!(E, [u,v])
        end
    end
    if !directed
        for e in E
            e1 = [e[2], e[1]]
            if !(e1 in E)
                push!(E, e1)
            end
        end
    end
    return V, E
end

function test(n)
    M = BitMatrix(rand([false, true], (n,n)))
    for i in 1:n
        M[i,i] = 0
    end
    E = Vector[]
    for i in 1:n, j in 1:n
        if M[i,j] == 1
            push!(E, [i,j])
        end
    end
    V = collect(1:n)
    chk = falses(n)
    for v in V
        chk[v] = (BFS(V, E, v) == BFS(M, v))
    end
    return all(chk)
end

# Adjacency list format
E = [[1,2], [1,3],[2,4],[3,4],[4,5],[4,6],[5,6]]
V = collect(1:6)

# Adjacency matrix
A = BitArray(undef, (6,6))
for e in E
    A[e...] = true
end

for i in V
    println(BFS(V, E, i) == BFS(A, i))
end

shortestPath(A, 1, 4)

M = BitMatrix(rand([false, true], (n,n)))
for i in 1:n
    M[i,i] = 0
end