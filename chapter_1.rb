require "stringio"

def compress(string)
  compressed = StringIO.new

  consecutive_count = 1
  last_char = string[0]

  (1...string.length).each do |i|
    if string[i] == last_char
      consecutive_count += 1
    else
      compressed.write("#{last_char}#{consecutive_count}")
      consecutive_count = 1
      last_char = string[i]
    end
  end
  compressed.write("#{last_char}#{consecutive_count}")
  (compressed.length < string.length && compressed.length > 0) ? compressed.string : string
end

def set_0_matrix(arr)
  cols_to_null = []

  arr.each do |row|
    null_row = false
    row.each_with_index do |v, i|
      if v == 0
        cols_to_null.push(i)
        null_row = true
      end
    end
    row.each_with_index { |v, i| row[i] = 0 } if null_row
  end

  cols_to_null.each do |col|
    (0...arr.length).each do |i|
      arr[i][col] = 0
    end
  end

  arr
end

def is_rotation?(a, b)
  return false if a.length != b.length
  (a + a).include?(b)
end
