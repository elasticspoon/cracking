Node = Struct.new(:value, :next)

def palindrome?(head)
  even_vals = {}
  num_odds = 0

  until head.nil?
    if even_vals[head.value] == false
      even_vals[head.value] = true
      num_odds -= 1
    else
      even_vals[head.value] = false
      num_odds += 1
    end
    head = head.next
  end

  num_odds <= 1
end
