module Queue

mutable struct queue
    data::Array{Number}
    front
    len
end

function New(x::Array{Number})
    q = queue(x, x[1], len(x))
    return q
end

function New()
    q = queue(Number[], nothing, 0)
    return q
end

function New(x::Number)
    q = queue([x], x, 1)
    return q
end

function enqueue!(q::queue, x::Number)
    push!(q.data, x)
    q.len += 1
    println(q.data)
end

function dequeue!(q::queue)
    if Base.length(q.data) <= 1
        q.data = Number[]
        q.len = 0
        q.front = nothing
        return
    end
    q.data = q.data[2:end]
    q.len -= 1
    q.front = q.data[1]
    println(q.data)
end

function front(q::queue)
    return q.front
end

function length(q::queue)
    return q.len
end

end