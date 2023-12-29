meetings = [["John", 7], ["Sarah", 3], ["John", 1], ["Sarah", 3], ["Sarah", 5]]

def interary(meetings, interval)
  meetings.sort_by! { |x| x[1] }

  result = []

  meetings.each do |meeting|
    _, time = meeting

    break if interval < time

    result.push(meeting)
    interval -= time
  end

  result
end

puts interary(meetings, 2) == [["John", 1]]
puts interary(meetings, 4) == [["John", 1], ["Sarah", 3]]
