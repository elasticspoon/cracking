class Card
  attr_reader :rank, :suit

  def initialize(rank, suit)
    @rank, @suit = rank, suit
  end

  def to_s
    "#{@rank} of #{@suit}"
  end
end

class Deck
  SUITS = %w[clubs diamonds hearts spades]
  RANKS = %w[1 2 3 4 5 6 7 8 9 10 jack queen king]

  def initialize(num_decks = 1)
    deck = []
    num_decks.times do
      SUITS.each do |suit|
        RANKS.each do |rank|
          deck << Card.new(rank, suit)
        end
      end
    end
    @cards = deck
  end

  def shuffle
    @cards.shuffle!
  end

  def draw
    @cards.pop
  end

  def length
    @cards.length
  end
end

class BlackJack
  def initialize(num_decks = 1)
    @deck = Deck.new(num_decks)
    @deck.shuffle
  end

  def play
    dealer_score = dealer_setup
    return if dealer_score >= 21 && check_winner(dealer_score, 0)
    player_score = player_turn(dealer_score, 0)
    return if player_score > 21 && check_winner(dealer_score, player_score)
    dealer_score = dealer_turn(dealer_score, player_score)

    check_winner(dealer_score, player_score)
  end

  def check_winner(dealer_score, player_score)
    if dealer_score <= 21 && dealer_score >= player_score || player_score > 21
      puts "Dealer has #{dealer_score}. You have #{player_score}."
      puts "Dealer wins!"
      return
    end

    if player_score <= 21 && player_score > dealer_score || dealer_score > 21
      puts "Dealer has #{dealer_score}. You have #{player_score}."
      puts "Player wins!"
    end
  end

  def draw(party)
    v = @deck.draw

    puts "#{party} drew #{v}"
    return 10 if v.rank == "jack" || v.rank == "queen" || v.rank == "king"
    v.rank.to_i
  end

  def dealer_setup
    res = 0
    2.times do
      v = draw("Dealer")
      v = dealer_ace(res) if v == 1
      res += v
    end
    res
  end

  def dealer_turn(dealer_score, player_score)
    while dealer_score <= player_score && dealer_score < 21
      v = draw("Dealer")
      v = dealer_ace(dealer_score) if v == 1
      dealer_score += v
    end

    dealer_score
  end

  def player_turn(dealer_score, player_score)
    while player_score < 21
      puts "Dealer has #{dealer_score}. You have #{player_score}. Hit or stay?"
      while (v = gets.chomp) != "hit" && v != "stay"
        puts "Please enter hit or stay"
      end

      if v == "hit"
        v = draw("Player")
        v = player_ace(player_score) if v == 1
        player_score += v
      else
        return player_score
      end
    end

    player_score
  end

  def dealer_ace(score)
    (score + 11 > 21) ? 1 : 11
  end

  def player_ace(score)
    puts "You drew an ace! 1 or 11?"
    while (v = gets.chomp.to_i) != 1 && v != 11
      puts "Please enter 1 or 11"
    end

    v
  end
end

BlackJack.new.play
