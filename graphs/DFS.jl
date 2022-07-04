# Depth First Search algorithm
# Input graph adjacency list G and vertex v
# Output list of vertices of G reachable from v

# Implementation with stack
function DFS(V, E, v)
    # mark all vertices as unexplored
    status = falses(length(V))
    # Initialise stack with v
    s = [v]
    st = first.(E)
    ed = last.(E)
    while !isempty(s)
        u = s[end]
        if length(s) == 1
            s = []
        else
            s = s[1:end-1]
        end
        if !status[u]
            status[u] = true
            for w ∈ ed[st .== u]
                push!(s, w)
            end
        end
    end
    return V[status]
end

# Recursive implementation
function recursiveDFS(V, E, v)
    status = falses(length(V))
    function DFS1(V, E, v)
        status[v] = true
        st = first.(E)
        ed = last.(E)
        for w ∈ ed[st .== v]
            if !status[w]
                DFS1(V, E, w)
            end
        end
    end
    DFS1(V,E,v)
    return V[status]
end