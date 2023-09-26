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
