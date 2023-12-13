def find_path(maze)
  stack = []
  visited = {}

  stack.push([0, 0])
  until stack.empty?
    curr = stack.last
    x, y = curr

    return stack if curr == [maze.rows - 1, maze.cols - 1]

    if x + 1 < maze.cols && !maze.invalid?(x + 1, y) && !visited[[x + 1, y]]
      stack.push([x + 1, y])
      next
    end
    if y + 1 < maze.rows && !maze.invalid?(x, y + 1) && !visited[[x, y + 1]]
      stack.push([x, y + 1])
      next
    end

    visited[curr] = true
    stack.pop
  end

  nil
end

def subsets(array)
  visited = {}

  subsets_helper(array, visited)

  visited.keys
end

def subsets_helper(array, visited)
  return if visited[array] || array.empty?
  visited[array] = true

  array.each_with_index do |_, i|
    new_array = [*array]
    new_array.delete_at(i)
    subsets_helper(new_array, visited) unless visited[new_array]
  end
end

def other_subsets(array)
  result = []
  array.length.times do |i|
    result += array.combination(i + 1).to_a
  end

  result
end
