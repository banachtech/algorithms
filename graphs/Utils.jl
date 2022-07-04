# Graph utility functions.
module Utils
using LinearAlgebra, Random

export createAdjList, createAdjMat

"""
    createAdjMat(n; directed=true)

Create a random adjacency matrix of size n. If directed is false then a symmetric matrix is returned.

"""
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

"""
    createAdjList(n, m; directed=true)

Create a random adjacency list for a graph of n nodes and m edges. If directed is false, then the list of edges has 2m elements; for ever edge (u,v) the edge (v,u) is also included.

"""
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

end