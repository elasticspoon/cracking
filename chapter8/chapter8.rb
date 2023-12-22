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

class TowerOfHanoi
  attr_reader :size, :towers

  def initialize(size)
    @size = size
    @towers = [[], [], []]
    @towers[0] = (1..size).to_a.reverse
  end

  def height(tower)
    @towers[tower].length
  end

  def move(from, to)
    raise "invalid move: from #{from} (#{@towers[from]}) to #{to} (#{@towers[to]})" if from < 0 || from > 2 || to < 0 || to > 2
    raise "invalid move: from #{from} (#{@towers[from]}) to #{to} (#{@towers[to]})" if @towers[from].empty? || (@towers[to].any? && @towers[from].last > @towers[to].last)
    @towers[to].push(@towers[from].pop)
  end

  def other_tower(a, b)
    case [a, b]
    when [0, 1], [1, 0]
      return 2
    when [0, 2], [2, 0]
      return 1
    when [1, 2], [2, 1]
      return 0
    end
    raise "invalid towers"
  end
end

def move_tower(tower)
  move_tower_helper(0, 2, tower.size, tower)
end

def move_tower_helper(from, to, size, tower)
  if size == 1
    tower.move(from, to)
  else
    other = tower.other_tower(from, to)
    move_tower_helper(from, other, size - 1, tower)
    tower.move(from, to)
    move_tower_helper(other, to, size - 1, tower)
  end
end

def permutations(string)
  result = {"" => true}

  until string.empty?
    new_result = {}
    c = string[0]
    result.each_key do |s|
      (s.length + 1).times do |i|
        new_result[(s[...i] + c + s[i...])] = true unless s[i] == c
      end
    end
    string = string[1..]
    result = new_result
  end

  result.keys
end

def perms(string)
  perms_helper("", string.chars.tally)
end

def perms_helper(prefix, map)
  return [prefix] if map.empty?
  res = []
  map.each do |k, v|
    new_map = map.dup
    new_map[k] -= 1
    new_map.delete(k) if new_map[k] == 0
    res.concat(perms_helper(prefix + k, new_map))
  end

  res
end

def parens(n)
  parens_helper("", n, n)
end

def parens_helper(prefix, open, closed)
  return [prefix] if open == 0 && closed == 0
  return parens_helper(prefix + ")", open, closed - 1) if open == 0
  return parens_helper(prefix + "(", open - 1, closed) if open == closed

  parens_helper(prefix + "(", open - 1, closed) + parens_helper(prefix + ")", open, closed - 1)
end

def tallest_stack(heights, widths, depths)
  bases = get_bases(heights, widths, depths)

  memo = []
  max_height = 0
  (0...heights.length).each do |i|
    height = get_h(i, heights, bases, memo)
    max_height = height if height > max_height
  end

  max_height
end

def get_bases(heights, widths, depths)
  bases = Array.new(heights.length) { nil }
  (0...heights.length).each do |i|
    (0...heights.length).each do |j|
      next unless widths[i] < widths[j] && depths[i] < depths[j] && heights[i] < heights[j]
      bases[i] = j if bases[i].nil? || heights[bases[i]] > heights[j]
    end
  end

  bases
end

def get_h(index, heights, bases, memo)
  return memo[index] if memo[index]
  return heights[index] unless bases[index]

  res = heights[index] + get_h(bases[index], heights, bases, memo)
  memo[index] = res
  res
end

BOARD_SIZE = 8
def get_index(x, y)
  x * BOARD_SIZE + y
end

def get_pos(index)
  [index / BOARD_SIZE, index % BOARD_SIZE]
end

def valid?(pot_index, queen_pots)
  x, y = get_pos(pot_index)
  queen_pots.none? do |index|
    new_x, new_y = get_pos(index)
    x == new_x || y == new_y || (x - new_x).abs == (y - new_y).abs
  end
end

def queen_locations
  res = []
  (0...BOARD_SIZE * BOARD_SIZE).each do |i|
    queen_pots = {}
    queen_pots[i] = true
    res.concat(q_helper(7, queen_pots))
  end
end

def q_helper(num_queens, queen_pots)
  return queen_pots.keys if num_queens == 0
  res = []

  (0...BOARD_SIZE * BOARD_SIZE).each do |i|
    next unless valid?(i, queen_pots.keys)

    new_queen_pots = queen_pots.dup
    new_queen_pots[i] = true

    res.concat(q_helper(num_queens - 1, new_queen_pots))
  end

  res
end

def max_stack(boxes)
  boxes = boxes.sort.reverse

  max_height = 0
  memo = []
  (0...boxes.length).each do |i|
    height = rec_max_stack(boxes, i, memo)
    max_height = height if height > max_height
  end

  max_height
end

def rec_max_stack(boxes, index, memo)
  return memo[index] if memo[index]

  max_height = boxes[index][0]

  (index...boxes.length).each do |i|
    next unless stackable?(boxes[index], boxes[i])

    pot_height = boxes[index][0] + rec_max_stack(boxes, i, memo)
    max_height = pot_height if pot_height > max_height
  end

  memo[index] = max_height
  max_height
end

def stackable?(bottom, top)
  top[0] < bottom[0] && top[1] < bottom[1] && top[2] < bottom[2]
end
