class Phrase

  def initialize(word)
    @word = word
    @result = Hash.new {|hash,k| hash[k] = 0}
  end

  def word_count
    words.each do |word|
      @result[word.downcase] += 1
    end
    @word = ''
    @result
  end

  private

  def words
    @word.split(/[^A-Za-z0-9]/).reject(&:empty?)
  end


end
