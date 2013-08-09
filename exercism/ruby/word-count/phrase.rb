class Phrase

  def initialize(word)
    @word = word
    @result = Hash.new {|hash,k| hash[k] = 0}
  end

  def word_count
    @word.split(/[^A-Za-z0-9]/).each do |word|
      @result[word.downcase] += 1 unless word.empty?
    end
    @word = ''
    @result
  end

end
