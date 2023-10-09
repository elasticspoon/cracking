class StackOfStacks
  private attr_reader :stacks, :max_size
  def initialize(max_size)
    @stacks = []
    @max_size = max_size
  end

  def push(value)
    stacks.push([]) if stacks.empty? || stacks.last.size == max_size
    stacks.last.push(value)
  end

  def pop
    return nil if stacks.empty?
    target = stacks.last
    value = target.pop
    stacks.pop if target.empty?

    value
  end

  def pop_at(index)
    return nil if stacks.empty? || index >= stacks.size || index < 0
    value = stacks[index].pop
    stacks.delete_at(index) if stacks[index].empty?
    value
  end
end

# Ghetto Tests
puts "StackOfStacks"
all_good = true
ss = StackOfStacks.new(3)
(1..10).each do |i|
  ss.push(i)
end

6.downto(4).each do |i|
  v = ss.pop_at(1)
  if v != i
    puts "expect #{i} got #{v}"
    all_good = false
  end
end

10.downto(7).each do |i|
  v = ss.pop
  if v != i
    puts "expect #{i} got #{v}"
    all_good = false
  end
end

3.downto(1).each do |i|
  v = ss.pop
  if v != i
    puts "expect #{i} got #{v}"
    all_good = false
  end
end

puts "all passed" if all_good
